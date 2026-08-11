package Data_Functor_Product_Nested

import (
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_product9 gopurs_runtime.Value
var once_product9 sync.Once
func Get_product9() gopurs_runtime.Value {
	once_product9.Do(func() {
		cache_product9 = gopurs_runtime.Func9(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value, i_8_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product9(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box, h_7_box, i_8_box)
})
	})
	return cache_product9
}

var cache_product8 gopurs_runtime.Value
var once_product8 sync.Once
func Get_product8() gopurs_runtime.Value {
	once_product8.Do(func() {
		cache_product8 = gopurs_runtime.Func8(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product8(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box, h_7_box)
})
	})
	return cache_product8
}

var cache_product7 gopurs_runtime.Value
var once_product7 sync.Once
func Get_product7() gopurs_runtime.Value {
	once_product7.Do(func() {
		cache_product7 = gopurs_runtime.Func7(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product7(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box)
})
	})
	return cache_product7
}

var cache_product6 gopurs_runtime.Value
var once_product6 sync.Once
func Get_product6() gopurs_runtime.Value {
	once_product6.Do(func() {
		cache_product6 = gopurs_runtime.Func6(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product6(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box)
})
	})
	return cache_product6
}

var cache_product5 gopurs_runtime.Value
var once_product5 sync.Once
func Get_product5() gopurs_runtime.Value {
	once_product5.Do(func() {
		cache_product5 = gopurs_runtime.Func5(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product5(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box)
})
	})
	return cache_product5
}

var cache_product4 gopurs_runtime.Value
var once_product4 sync.Once
func Get_product4() gopurs_runtime.Value {
	once_product4.Do(func() {
		cache_product4 = gopurs_runtime.Func4(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product4(a_0_box, b_1_box, c_2_box, d_3_box)
})
	})
	return cache_product4
}

var cache_product3 gopurs_runtime.Value
var once_product3 sync.Once
func Get_product3() gopurs_runtime.Value {
	once_product3.Do(func() {
		cache_product3 = gopurs_runtime.Func3(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product3(a_0_box, b_1_box, c_2_box)
})
	})
	return cache_product3
}

var cache_product2 gopurs_runtime.Value
var once_product2 sync.Once
func Get_product2() gopurs_runtime.Value {
	once_product2.Do(func() {
		cache_product2 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product2(a_0_box, b_1_box)
})
	})
	return cache_product2
}

var cache_product10 gopurs_runtime.Value
var once_product10 sync.Once
func Get_product10() gopurs_runtime.Value {
	once_product10.Do(func() {
		cache_product10 = gopurs_runtime.Func10(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value, i_8_box gopurs_runtime.Value, j_9_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product10(a_0_box, b_1_box, c_2_box, d_3_box, e_4_box, f_5_box, g_6_box, h_7_box, i_8_box, j_9_box)
})
	})
	return cache_product10
}

var cache_product1 gopurs_runtime.Value
var once_product1 sync.Once
func Get_product1() gopurs_runtime.Value {
	once_product1.Do(func() {
		cache_product1 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product1(a_0_box)
})
	})
	return cache_product1
}

var cache_get9 gopurs_runtime.Value
var once_get9 sync.Once
func Get_get9() gopurs_runtime.Value {
	once_get9.Do(func() {
		cache_get9 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get9(v_0_box)
})
	})
	return cache_get9
}

var cache_get8 gopurs_runtime.Value
var once_get8 sync.Once
func Get_get8() gopurs_runtime.Value {
	once_get8.Do(func() {
		cache_get8 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get8(v_0_box)
})
	})
	return cache_get8
}

var cache_get7 gopurs_runtime.Value
var once_get7 sync.Once
func Get_get7() gopurs_runtime.Value {
	once_get7.Do(func() {
		cache_get7 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get7(v_0_box)
})
	})
	return cache_get7
}

var cache_get6 gopurs_runtime.Value
var once_get6 sync.Once
func Get_get6() gopurs_runtime.Value {
	once_get6.Do(func() {
		cache_get6 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get6(v_0_box)
})
	})
	return cache_get6
}

var cache_get5 gopurs_runtime.Value
var once_get5 sync.Once
func Get_get5() gopurs_runtime.Value {
	once_get5.Do(func() {
		cache_get5 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get5(v_0_box)
})
	})
	return cache_get5
}

var cache_get4 gopurs_runtime.Value
var once_get4 sync.Once
func Get_get4() gopurs_runtime.Value {
	once_get4.Do(func() {
		cache_get4 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get4(v_0_box)
})
	})
	return cache_get4
}

var cache_get3 gopurs_runtime.Value
var once_get3 sync.Once
func Get_get3() gopurs_runtime.Value {
	once_get3.Do(func() {
		cache_get3 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get3(v_0_box)
})
	})
	return cache_get3
}

var cache_get2 gopurs_runtime.Value
var once_get2 sync.Once
func Get_get2() gopurs_runtime.Value {
	once_get2.Do(func() {
		cache_get2 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get2(v_0_box)
})
	})
	return cache_get2
}

var cache_get10 gopurs_runtime.Value
var once_get10 sync.Once
func Get_get10() gopurs_runtime.Value {
	once_get10.Do(func() {
		cache_get10 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get10(v_0_box)
})
	})
	return cache_get10
}

var cache_get1 gopurs_runtime.Value
var once_get1 sync.Once
func Get_get1() gopurs_runtime.Value {
	once_get1.Do(func() {
		cache_get1 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get1(v_0_box)
})
	})
	return cache_get1
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, d_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, e_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, g_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_7, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, i_8, pkg_Data_Unit.Get_unit()})}})}})}})}})}})}})}})}})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, d_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, e_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, g_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_7, pkg_Data_Unit.Get_unit()})}})}})}})}})}})}})}})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, d_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, e_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, g_6, pkg_Data_Unit.Get_unit()})}})}})}})}})}})}})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, d_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, e_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, f_5, pkg_Data_Unit.Get_unit()})}})}})}})}})}})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, d_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, e_4, pkg_Data_Unit.Get_unit()})}})}})}})}})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, d_3, pkg_Data_Unit.Get_unit()})}})}})}})}
}

func Call_product3(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, c_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var c_2 gopurs_runtime.Value = c_2_loop
_ = c_2
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, c_2, pkg_Data_Unit.Get_unit()})}})}})}
}

func Call_product2(a_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, b_1, pkg_Data_Unit.Get_unit()})}})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, d_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, e_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, g_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, h_7, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, i_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, j_9, pkg_Data_Unit.Get_unit()})}})}})}})}})}})}})}})}})}})}
}

func Call_product1(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, pkg_Data_Unit.Get_unit()})}
}

func Call_get9(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}

func Call_get8(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}

func Call_get7(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}

func Call_get6(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}

func Call_get5(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}

func Call_get4(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}

func Call_get3(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}

func Call_get2(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V0
}

func Call_get10(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V1.UnsafePtr).V0
}

func Call_get1(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
}


