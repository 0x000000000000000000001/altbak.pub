package Data_Tuple_Nested

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	unsafe "unsafe"
)

var cache_uncurry9 gopurs_runtime.Value
var once_uncurry9 sync.Once
func Get_uncurry9() gopurs_runtime.Value {
	once_uncurry9.Do(func() {
		cache_uncurry9 = gopurs_runtime.Func2(func(f_prime_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_uncurry9(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}, inner_arg6 interface{}, inner_arg7 interface{}, inner_arg8 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply9(f_prime_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5), gopurs_runtime.Any(inner_arg6), gopurs_runtime.Any(inner_arg7), gopurs_runtime.Any(inner_arg8)))
}, v_1_box))
})
	})
	return cache_uncurry9
}

var cache_uncurry8 gopurs_runtime.Value
var once_uncurry8 sync.Once
func Get_uncurry8() gopurs_runtime.Value {
	once_uncurry8.Do(func() {
		cache_uncurry8 = gopurs_runtime.Func2(func(f_prime_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_uncurry8(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}, inner_arg6 interface{}, inner_arg7 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply8(f_prime_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5), gopurs_runtime.Any(inner_arg6), gopurs_runtime.Any(inner_arg7)))
}, v_1_box))
})
	})
	return cache_uncurry8
}

var cache_uncurry7 gopurs_runtime.Value
var once_uncurry7 sync.Once
func Get_uncurry7() gopurs_runtime.Value {
	once_uncurry7.Do(func() {
		cache_uncurry7 = gopurs_runtime.Func2(func(f_prime_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_uncurry7(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}, inner_arg6 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply7(f_prime_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5), gopurs_runtime.Any(inner_arg6)))
}, v_1_box))
})
	})
	return cache_uncurry7
}

var cache_uncurry6 gopurs_runtime.Value
var once_uncurry6 sync.Once
func Get_uncurry6() gopurs_runtime.Value {
	once_uncurry6.Do(func() {
		cache_uncurry6 = gopurs_runtime.Func2(func(f_prime_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_uncurry6(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply6(f_prime_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5)))
}, v_1_box))
})
	})
	return cache_uncurry6
}

var cache_uncurry5 gopurs_runtime.Value
var once_uncurry5 sync.Once
func Get_uncurry5() gopurs_runtime.Value {
	once_uncurry5.Do(func() {
		cache_uncurry5 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_uncurry5(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply5(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4)))
}, v_1_box))
})
	})
	return cache_uncurry5
}

var cache_uncurry4 gopurs_runtime.Value
var once_uncurry4 sync.Once
func Get_uncurry4() gopurs_runtime.Value {
	once_uncurry4.Do(func() {
		cache_uncurry4 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_uncurry4(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply4(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3)))
}, v_1_box))
})
	})
	return cache_uncurry4
}

var cache_uncurry3 gopurs_runtime.Value
var once_uncurry3 sync.Once
func Get_uncurry3() gopurs_runtime.Value {
	once_uncurry3.Do(func() {
		cache_uncurry3 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_uncurry3(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply3(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2)))
}, v_1_box))
})
	})
	return cache_uncurry3
}

var cache_uncurry2 gopurs_runtime.Value
var once_uncurry2 sync.Once
func Get_uncurry2() gopurs_runtime.Value {
	once_uncurry2.Do(func() {
		cache_uncurry2 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_uncurry2(func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, v_1_box))
})
	})
	return cache_uncurry2
}

var cache_uncurry10 gopurs_runtime.Value
var once_uncurry10 sync.Once
func Get_uncurry10() gopurs_runtime.Value {
	once_uncurry10.Do(func() {
		cache_uncurry10 = gopurs_runtime.Func2(func(f_prime_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_uncurry10(func(inner_arg0 interface{}, inner_arg1 interface{}, inner_arg2 interface{}, inner_arg3 interface{}, inner_arg4 interface{}, inner_arg5 interface{}, inner_arg6 interface{}, inner_arg7 interface{}, inner_arg8 interface{}, inner_arg9 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply10(f_prime_0_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1), gopurs_runtime.Any(inner_arg2), gopurs_runtime.Any(inner_arg3), gopurs_runtime.Any(inner_arg4), gopurs_runtime.Any(inner_arg5), gopurs_runtime.Any(inner_arg6), gopurs_runtime.Any(inner_arg7), gopurs_runtime.Any(inner_arg8), gopurs_runtime.Any(inner_arg9)))
}, v_1_box))
})
	})
	return cache_uncurry10
}

var cache_uncurry1 gopurs_runtime.Value
var once_uncurry1 sync.Once
func Get_uncurry1() gopurs_runtime.Value {
	once_uncurry1.Do(func() {
		cache_uncurry1 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_uncurry1(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
}, v_1_box))
})
	})
	return cache_uncurry1
}

var cache_tuple9 gopurs_runtime.Value
var once_tuple9 sync.Once
func Get_tuple9() gopurs_runtime.Value {
	once_tuple9.Do(func() {
		cache_tuple9 = gopurs_runtime.Func9(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value, i_8_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tuple9(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(c_2_box), gopurs_runtime.UnboxAny(d_3_box), gopurs_runtime.UnboxAny(e_4_box), gopurs_runtime.UnboxAny(f_5_box), gopurs_runtime.UnboxAny(g_6_box), gopurs_runtime.UnboxAny(h_7_box), gopurs_runtime.UnboxAny(i_8_box))
})
	})
	return cache_tuple9
}

var cache_tuple8 gopurs_runtime.Value
var once_tuple8 sync.Once
func Get_tuple8() gopurs_runtime.Value {
	once_tuple8.Do(func() {
		cache_tuple8 = gopurs_runtime.Func8(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tuple8(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(c_2_box), gopurs_runtime.UnboxAny(d_3_box), gopurs_runtime.UnboxAny(e_4_box), gopurs_runtime.UnboxAny(f_5_box), gopurs_runtime.UnboxAny(g_6_box), gopurs_runtime.UnboxAny(h_7_box))
})
	})
	return cache_tuple8
}

var cache_tuple7 gopurs_runtime.Value
var once_tuple7 sync.Once
func Get_tuple7() gopurs_runtime.Value {
	once_tuple7.Do(func() {
		cache_tuple7 = gopurs_runtime.Func7(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tuple7(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(c_2_box), gopurs_runtime.UnboxAny(d_3_box), gopurs_runtime.UnboxAny(e_4_box), gopurs_runtime.UnboxAny(f_5_box), gopurs_runtime.UnboxAny(g_6_box))
})
	})
	return cache_tuple7
}

var cache_tuple6 gopurs_runtime.Value
var once_tuple6 sync.Once
func Get_tuple6() gopurs_runtime.Value {
	once_tuple6.Do(func() {
		cache_tuple6 = gopurs_runtime.Func6(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tuple6(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(c_2_box), gopurs_runtime.UnboxAny(d_3_box), gopurs_runtime.UnboxAny(e_4_box), gopurs_runtime.UnboxAny(f_5_box))
})
	})
	return cache_tuple6
}

var cache_tuple5 gopurs_runtime.Value
var once_tuple5 sync.Once
func Get_tuple5() gopurs_runtime.Value {
	once_tuple5.Do(func() {
		cache_tuple5 = gopurs_runtime.Func5(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tuple5(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(c_2_box), gopurs_runtime.UnboxAny(d_3_box), gopurs_runtime.UnboxAny(e_4_box))
})
	})
	return cache_tuple5
}

var cache_tuple4 gopurs_runtime.Value
var once_tuple4 sync.Once
func Get_tuple4() gopurs_runtime.Value {
	once_tuple4.Do(func() {
		cache_tuple4 = gopurs_runtime.Func4(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tuple4(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(c_2_box), gopurs_runtime.UnboxAny(d_3_box))
})
	})
	return cache_tuple4
}

var cache_tuple3 gopurs_runtime.Value
var once_tuple3 sync.Once
func Get_tuple3() gopurs_runtime.Value {
	once_tuple3.Do(func() {
		cache_tuple3 = gopurs_runtime.Func3(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tuple3(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(c_2_box))
})
	})
	return cache_tuple3
}

var cache_tuple2 gopurs_runtime.Value
var once_tuple2 sync.Once
func Get_tuple2() gopurs_runtime.Value {
	once_tuple2.Do(func() {
		cache_tuple2 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tuple2(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(b_1_box))
})
	})
	return cache_tuple2
}

var cache_tuple10 gopurs_runtime.Value
var once_tuple10 sync.Once
func Get_tuple10() gopurs_runtime.Value {
	once_tuple10.Do(func() {
		cache_tuple10 = gopurs_runtime.Func10(func(a_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, c_2_box gopurs_runtime.Value, d_3_box gopurs_runtime.Value, e_4_box gopurs_runtime.Value, f_5_box gopurs_runtime.Value, g_6_box gopurs_runtime.Value, h_7_box gopurs_runtime.Value, i_8_box gopurs_runtime.Value, j_9_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tuple10(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(b_1_box), gopurs_runtime.UnboxAny(c_2_box), gopurs_runtime.UnboxAny(d_3_box), gopurs_runtime.UnboxAny(e_4_box), gopurs_runtime.UnboxAny(f_5_box), gopurs_runtime.UnboxAny(g_6_box), gopurs_runtime.UnboxAny(h_7_box), gopurs_runtime.UnboxAny(i_8_box), gopurs_runtime.UnboxAny(j_9_box))
})
	})
	return cache_tuple10
}

var cache_tuple1 gopurs_runtime.Value
var once_tuple1 sync.Once
func Get_tuple1() gopurs_runtime.Value {
	once_tuple1.Do(func() {
		cache_tuple1 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tuple1(gopurs_runtime.UnboxAny(a_0_box))
})
	})
	return cache_tuple1
}

var cache_over9 gopurs_runtime.Value
var once_over9 sync.Once
func Get_over9() gopurs_runtime.Value {
	once_over9.Do(func() {
		cache_over9 = gopurs_runtime.Func2(func(o_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over9(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(o_0_box, gopurs_runtime.Any(inner_arg0)))
}, v_1_box)
})
	})
	return cache_over9
}

var cache_over8 gopurs_runtime.Value
var once_over8 sync.Once
func Get_over8() gopurs_runtime.Value {
	once_over8.Do(func() {
		cache_over8 = gopurs_runtime.Func2(func(o_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over8(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(o_0_box, gopurs_runtime.Any(inner_arg0)))
}, v_1_box)
})
	})
	return cache_over8
}

var cache_over7 gopurs_runtime.Value
var once_over7 sync.Once
func Get_over7() gopurs_runtime.Value {
	once_over7.Do(func() {
		cache_over7 = gopurs_runtime.Func2(func(o_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over7(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(o_0_box, gopurs_runtime.Any(inner_arg0)))
}, v_1_box)
})
	})
	return cache_over7
}

var cache_over6 gopurs_runtime.Value
var once_over6 sync.Once
func Get_over6() gopurs_runtime.Value {
	once_over6.Do(func() {
		cache_over6 = gopurs_runtime.Func2(func(o_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over6(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(o_0_box, gopurs_runtime.Any(inner_arg0)))
}, v_1_box)
})
	})
	return cache_over6
}

var cache_over5 gopurs_runtime.Value
var once_over5 sync.Once
func Get_over5() gopurs_runtime.Value {
	once_over5.Do(func() {
		cache_over5 = gopurs_runtime.Func2(func(o_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over5(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(o_0_box, gopurs_runtime.Any(inner_arg0)))
}, v_1_box)
})
	})
	return cache_over5
}

var cache_over4 gopurs_runtime.Value
var once_over4 sync.Once
func Get_over4() gopurs_runtime.Value {
	once_over4.Do(func() {
		cache_over4 = gopurs_runtime.Func2(func(o_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over4(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(o_0_box, gopurs_runtime.Any(inner_arg0)))
}, v_1_box)
})
	})
	return cache_over4
}

var cache_over3 gopurs_runtime.Value
var once_over3 sync.Once
func Get_over3() gopurs_runtime.Value {
	once_over3.Do(func() {
		cache_over3 = gopurs_runtime.Func2(func(o_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over3(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(o_0_box, gopurs_runtime.Any(inner_arg0)))
}, v_1_box)
})
	})
	return cache_over3
}

var cache_over2 gopurs_runtime.Value
var once_over2 sync.Once
func Get_over2() gopurs_runtime.Value {
	once_over2.Do(func() {
		cache_over2 = gopurs_runtime.Func2(func(o_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over2(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(o_0_box, gopurs_runtime.Any(inner_arg0)))
}, v_1_box)
})
	})
	return cache_over2
}

var cache_over10 gopurs_runtime.Value
var once_over10 sync.Once
func Get_over10() gopurs_runtime.Value {
	once_over10.Do(func() {
		cache_over10 = gopurs_runtime.Func2(func(o_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over10(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(o_0_box, gopurs_runtime.Any(inner_arg0)))
}, v_1_box)
})
	})
	return cache_over10
}

var cache_over1 gopurs_runtime.Value
var once_over1 sync.Once
func Get_over1() gopurs_runtime.Value {
	once_over1.Do(func() {
		cache_over1 = gopurs_runtime.Func2(func(o_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over1(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(o_0_box, gopurs_runtime.Any(inner_arg0)))
}, v_1_box)
})
	})
	return cache_over1
}

var cache_get9 gopurs_runtime.Value
var once_get9 sync.Once
func Get_get9() gopurs_runtime.Value {
	once_get9.Do(func() {
		cache_get9 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_get9(v_0_box))
})
	})
	return cache_get9
}

var cache_get8 gopurs_runtime.Value
var once_get8 sync.Once
func Get_get8() gopurs_runtime.Value {
	once_get8.Do(func() {
		cache_get8 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_get8(v_0_box))
})
	})
	return cache_get8
}

var cache_get7 gopurs_runtime.Value
var once_get7 sync.Once
func Get_get7() gopurs_runtime.Value {
	once_get7.Do(func() {
		cache_get7 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_get7(v_0_box))
})
	})
	return cache_get7
}

var cache_get6 gopurs_runtime.Value
var once_get6 sync.Once
func Get_get6() gopurs_runtime.Value {
	once_get6.Do(func() {
		cache_get6 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_get6(v_0_box))
})
	})
	return cache_get6
}

var cache_get5 gopurs_runtime.Value
var once_get5 sync.Once
func Get_get5() gopurs_runtime.Value {
	once_get5.Do(func() {
		cache_get5 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_get5(v_0_box))
})
	})
	return cache_get5
}

var cache_get4 gopurs_runtime.Value
var once_get4 sync.Once
func Get_get4() gopurs_runtime.Value {
	once_get4.Do(func() {
		cache_get4 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_get4(v_0_box))
})
	})
	return cache_get4
}

var cache_get3 gopurs_runtime.Value
var once_get3 sync.Once
func Get_get3() gopurs_runtime.Value {
	once_get3.Do(func() {
		cache_get3 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_get3(v_0_box))
})
	})
	return cache_get3
}

var cache_get2 gopurs_runtime.Value
var once_get2 sync.Once
func Get_get2() gopurs_runtime.Value {
	once_get2.Do(func() {
		cache_get2 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_get2(v_0_box))
})
	})
	return cache_get2
}

var cache_get10 gopurs_runtime.Value
var once_get10 sync.Once
func Get_get10() gopurs_runtime.Value {
	once_get10.Do(func() {
		cache_get10 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_get10(v_0_box))
})
	})
	return cache_get10
}

var cache_get1 gopurs_runtime.Value
var once_get1 sync.Once
func Get_get1() gopurs_runtime.Value {
	once_get1.Do(func() {
		cache_get1 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_get1(v_0_box))
})
	})
	return cache_get1
}

var cache_curry9 gopurs_runtime.Value
var once_curry9 sync.Once
func Get_curry9() gopurs_runtime.Value {
	once_curry9.Do(func() {
		cache_curry9 = gopurs_runtime.Func(func(z_0_box gopurs_runtime.Value) gopurs_runtime.Value {
var z_0_loop interface{} = gopurs_runtime.UnboxAny(z_0_box)
return gopurs_runtime.Func(func(f_prime_1_box gopurs_runtime.Value) gopurs_runtime.Value {
var f_prime_1_loop func(gopurs_runtime.Value) interface{} = func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_prime_1_box, inner_arg0))
}
return gopurs_runtime.Func(func(a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
var a_2_loop interface{} = gopurs_runtime.UnboxAny(a_2_box)
return gopurs_runtime.Func(func(b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
var b_3_loop interface{} = gopurs_runtime.UnboxAny(b_3_box)
return gopurs_runtime.Func(func(c_4_box gopurs_runtime.Value) gopurs_runtime.Value {
var c_4_loop interface{} = gopurs_runtime.UnboxAny(c_4_box)
return gopurs_runtime.Func(func(d_5_box gopurs_runtime.Value) gopurs_runtime.Value {
var d_5_loop interface{} = gopurs_runtime.UnboxAny(d_5_box)
return gopurs_runtime.Func(func(e_6_box gopurs_runtime.Value) gopurs_runtime.Value {
var e_6_loop interface{} = gopurs_runtime.UnboxAny(e_6_box)
return gopurs_runtime.Func(func(f_7_box gopurs_runtime.Value) gopurs_runtime.Value {
var f_7_loop interface{} = gopurs_runtime.UnboxAny(f_7_box)
return gopurs_runtime.Func(func(g_8_box gopurs_runtime.Value) gopurs_runtime.Value {
var g_8_loop interface{} = gopurs_runtime.UnboxAny(g_8_box)
return gopurs_runtime.Func(func(h_9_box gopurs_runtime.Value) gopurs_runtime.Value {
var h_9_loop interface{} = gopurs_runtime.UnboxAny(h_9_box)
return gopurs_runtime.Func(func(i_10_box gopurs_runtime.Value) gopurs_runtime.Value {
var i_10_loop interface{} = gopurs_runtime.UnboxAny(i_10_box)
return func() gopurs_runtime.Value {
var z_0 interface{} = z_0_loop
_ = z_0
var f_prime_1 func(gopurs_runtime.Value) interface{} = f_prime_1_loop
_ = f_prime_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
var c_4 interface{} = c_4_loop
_ = c_4
var d_5 interface{} = d_5_loop
_ = d_5
var e_6 interface{} = e_6_loop
_ = e_6
var f_7 interface{} = f_7_loop
_ = f_7
var g_8 interface{} = g_8_loop
_ = g_8
var h_9 interface{} = h_9_loop
_ = h_9
var i_10 interface{} = i_10_loop
_ = i_10
return gopurs_runtime.Any(f_prime_1(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{e_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{f_7, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{g_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{h_9, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{i_10, z_0})}})}})}})}})}})}})}})}})})))
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
	return cache_curry9
}

var cache_curry8 gopurs_runtime.Value
var once_curry8 sync.Once
func Get_curry8() gopurs_runtime.Value {
	once_curry8.Do(func() {
		cache_curry8 = gopurs_runtime.Func10(func(z_0_box gopurs_runtime.Value, f_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value, d_5_box gopurs_runtime.Value, e_6_box gopurs_runtime.Value, f_7_box gopurs_runtime.Value, g_8_box gopurs_runtime.Value, h_9_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_curry8(gopurs_runtime.UnboxAny(z_0_box), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_prime_1_box, inner_arg0))
}, gopurs_runtime.UnboxAny(a_2_box), gopurs_runtime.UnboxAny(b_3_box), gopurs_runtime.UnboxAny(c_4_box), gopurs_runtime.UnboxAny(d_5_box), gopurs_runtime.UnboxAny(e_6_box), gopurs_runtime.UnboxAny(f_7_box), gopurs_runtime.UnboxAny(g_8_box), gopurs_runtime.UnboxAny(h_9_box)))
})
	})
	return cache_curry8
}

var cache_curry7 gopurs_runtime.Value
var once_curry7 sync.Once
func Get_curry7() gopurs_runtime.Value {
	once_curry7.Do(func() {
		cache_curry7 = gopurs_runtime.Func9(func(z_0_box gopurs_runtime.Value, f_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value, d_5_box gopurs_runtime.Value, e_6_box gopurs_runtime.Value, f_7_box gopurs_runtime.Value, g_8_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_curry7(gopurs_runtime.UnboxAny(z_0_box), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_prime_1_box, inner_arg0))
}, gopurs_runtime.UnboxAny(a_2_box), gopurs_runtime.UnboxAny(b_3_box), gopurs_runtime.UnboxAny(c_4_box), gopurs_runtime.UnboxAny(d_5_box), gopurs_runtime.UnboxAny(e_6_box), gopurs_runtime.UnboxAny(f_7_box), gopurs_runtime.UnboxAny(g_8_box)))
})
	})
	return cache_curry7
}

var cache_curry6 gopurs_runtime.Value
var once_curry6 sync.Once
func Get_curry6() gopurs_runtime.Value {
	once_curry6.Do(func() {
		cache_curry6 = gopurs_runtime.Func8(func(z_0_box gopurs_runtime.Value, f_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value, d_5_box gopurs_runtime.Value, e_6_box gopurs_runtime.Value, f_7_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_curry6(gopurs_runtime.UnboxAny(z_0_box), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_prime_1_box, inner_arg0))
}, gopurs_runtime.UnboxAny(a_2_box), gopurs_runtime.UnboxAny(b_3_box), gopurs_runtime.UnboxAny(c_4_box), gopurs_runtime.UnboxAny(d_5_box), gopurs_runtime.UnboxAny(e_6_box), gopurs_runtime.UnboxAny(f_7_box)))
})
	})
	return cache_curry6
}

var cache_curry5 gopurs_runtime.Value
var once_curry5 sync.Once
func Get_curry5() gopurs_runtime.Value {
	once_curry5.Do(func() {
		cache_curry5 = gopurs_runtime.Func7(func(z_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value, d_5_box gopurs_runtime.Value, e_6_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_curry5(gopurs_runtime.UnboxAny(z_0_box), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, inner_arg0))
}, gopurs_runtime.UnboxAny(a_2_box), gopurs_runtime.UnboxAny(b_3_box), gopurs_runtime.UnboxAny(c_4_box), gopurs_runtime.UnboxAny(d_5_box), gopurs_runtime.UnboxAny(e_6_box)))
})
	})
	return cache_curry5
}

var cache_curry4 gopurs_runtime.Value
var once_curry4 sync.Once
func Get_curry4() gopurs_runtime.Value {
	once_curry4.Do(func() {
		cache_curry4 = gopurs_runtime.Func6(func(z_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value, d_5_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_curry4(gopurs_runtime.UnboxAny(z_0_box), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, inner_arg0))
}, gopurs_runtime.UnboxAny(a_2_box), gopurs_runtime.UnboxAny(b_3_box), gopurs_runtime.UnboxAny(c_4_box), gopurs_runtime.UnboxAny(d_5_box)))
})
	})
	return cache_curry4
}

var cache_curry3 gopurs_runtime.Value
var once_curry3 sync.Once
func Get_curry3() gopurs_runtime.Value {
	once_curry3.Do(func() {
		cache_curry3 = gopurs_runtime.Func5(func(z_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value, c_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_curry3(gopurs_runtime.UnboxAny(z_0_box), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, inner_arg0))
}, gopurs_runtime.UnboxAny(a_2_box), gopurs_runtime.UnboxAny(b_3_box), gopurs_runtime.UnboxAny(c_4_box)))
})
	})
	return cache_curry3
}

var cache_curry2 gopurs_runtime.Value
var once_curry2 sync.Once
func Get_curry2() gopurs_runtime.Value {
	once_curry2.Do(func() {
		cache_curry2 = gopurs_runtime.Func4(func(z_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_curry2(gopurs_runtime.UnboxAny(z_0_box), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, inner_arg0))
}, gopurs_runtime.UnboxAny(a_2_box), gopurs_runtime.UnboxAny(b_3_box)))
})
	})
	return cache_curry2
}

var cache_curry10 gopurs_runtime.Value
var once_curry10 sync.Once
func Get_curry10() gopurs_runtime.Value {
	once_curry10.Do(func() {
		cache_curry10 = gopurs_runtime.Func(func(z_0_box gopurs_runtime.Value) gopurs_runtime.Value {
var z_0_loop interface{} = gopurs_runtime.UnboxAny(z_0_box)
return gopurs_runtime.Func(func(f_prime_1_box gopurs_runtime.Value) gopurs_runtime.Value {
var f_prime_1_loop func(gopurs_runtime.Value) interface{} = func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_prime_1_box, inner_arg0))
}
return gopurs_runtime.Func(func(a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
var a_2_loop interface{} = gopurs_runtime.UnboxAny(a_2_box)
return gopurs_runtime.Func(func(b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
var b_3_loop interface{} = gopurs_runtime.UnboxAny(b_3_box)
return gopurs_runtime.Func(func(c_4_box gopurs_runtime.Value) gopurs_runtime.Value {
var c_4_loop interface{} = gopurs_runtime.UnboxAny(c_4_box)
return gopurs_runtime.Func(func(d_5_box gopurs_runtime.Value) gopurs_runtime.Value {
var d_5_loop interface{} = gopurs_runtime.UnboxAny(d_5_box)
return gopurs_runtime.Func(func(e_6_box gopurs_runtime.Value) gopurs_runtime.Value {
var e_6_loop interface{} = gopurs_runtime.UnboxAny(e_6_box)
return gopurs_runtime.Func(func(f_7_box gopurs_runtime.Value) gopurs_runtime.Value {
var f_7_loop interface{} = gopurs_runtime.UnboxAny(f_7_box)
return gopurs_runtime.Func(func(g_8_box gopurs_runtime.Value) gopurs_runtime.Value {
var g_8_loop interface{} = gopurs_runtime.UnboxAny(g_8_box)
return gopurs_runtime.Func(func(h_9_box gopurs_runtime.Value) gopurs_runtime.Value {
var h_9_loop interface{} = gopurs_runtime.UnboxAny(h_9_box)
return gopurs_runtime.Func(func(i_10_box gopurs_runtime.Value) gopurs_runtime.Value {
var i_10_loop interface{} = gopurs_runtime.UnboxAny(i_10_box)
return gopurs_runtime.Func(func(j_11_box gopurs_runtime.Value) gopurs_runtime.Value {
var j_11_loop interface{} = gopurs_runtime.UnboxAny(j_11_box)
return func() gopurs_runtime.Value {
var z_0 interface{} = z_0_loop
_ = z_0
var f_prime_1 func(gopurs_runtime.Value) interface{} = f_prime_1_loop
_ = f_prime_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
var c_4 interface{} = c_4_loop
_ = c_4
var d_5 interface{} = d_5_loop
_ = d_5
var e_6 interface{} = e_6_loop
_ = e_6
var f_7 interface{} = f_7_loop
_ = f_7
var g_8 interface{} = g_8_loop
_ = g_8
var h_9 interface{} = h_9_loop
_ = h_9
var i_10 interface{} = i_10_loop
_ = i_10
var j_11 interface{} = j_11_loop
_ = j_11
return gopurs_runtime.Any(f_prime_1(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{e_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{f_7, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{g_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{h_9, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{i_10, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{j_11, z_0})}})}})}})}})}})}})}})}})}})})))
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
	})
	return cache_curry10
}

var cache_curry1 gopurs_runtime.Value
var once_curry1 sync.Once
func Get_curry1() gopurs_runtime.Value {
	once_curry1.Do(func() {
		cache_curry1 = gopurs_runtime.Func3(func(z_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_curry1(gopurs_runtime.UnboxAny(z_0_box), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, inner_arg0))
}, gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_curry1
}

func Call_uncurry9(f_prime_0_loop func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_prime_0 func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{} = f_prime_0_loop
_ = f_prime_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_prime_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)))))
}

func Call_uncurry8(f_prime_0_loop func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_prime_0 func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{} = f_prime_0_loop
_ = f_prime_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_prime_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)))))
}

func Call_uncurry7(f_prime_0_loop func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_prime_0 func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{} = f_prime_0_loop
_ = f_prime_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_prime_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)))))
}

func Call_uncurry6(f_prime_0_loop func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_prime_0 func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{} = f_prime_0_loop
_ = f_prime_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_prime_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)))))
}

func Call_uncurry5(f_0_loop func(interface{}, interface{}, interface{}, interface{}, interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_0 func(interface{}, interface{}, interface{}, interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)))))
}

func Call_uncurry4(f_0_loop func(interface{}, interface{}, interface{}, interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_0 func(interface{}, interface{}, interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)))))
}

func Call_uncurry3(f_0_loop func(interface{}, interface{}, interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_0 func(interface{}, interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)))))
}

func Call_uncurry2(f_0_loop func(interface{}, interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_0 func(interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)))))
}

func Call_uncurry10(f_prime_0_loop func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_prime_0 func(interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) interface{} = f_prime_0_loop
_ = f_prime_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_prime_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)))))
}

func Call_uncurry1(f_0_loop func(interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)))))
}

func Call_tuple9(a_0_loop interface{}, b_1_loop interface{}, c_2_loop interface{}, d_3_loop interface{}, e_4_loop interface{}, f_5_loop interface{}, g_6_loop interface{}, h_7_loop interface{}, i_8_loop interface{}) gopurs_runtime.Value {
var a_0 interface{} = a_0_loop
_ = a_0
var b_1 interface{} = b_1_loop
_ = b_1
var c_2 interface{} = c_2_loop
_ = c_2
var d_3 interface{} = d_3_loop
_ = d_3
var e_4 interface{} = e_4_loop
_ = e_4
var f_5 interface{} = f_5_loop
_ = f_5
var g_6 interface{} = g_6_loop
_ = g_6
var h_7 interface{} = h_7_loop
_ = h_7
var i_8 interface{} = i_8_loop
_ = i_8
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{e_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{g_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{h_7, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{i_8, gopurs_runtime.UnboxAny(pkg_Data_Unit.Get_unit())})}})}})}})}})}})}})}})}})})
}

func Call_tuple8(a_0_loop interface{}, b_1_loop interface{}, c_2_loop interface{}, d_3_loop interface{}, e_4_loop interface{}, f_5_loop interface{}, g_6_loop interface{}, h_7_loop interface{}) gopurs_runtime.Value {
var a_0 interface{} = a_0_loop
_ = a_0
var b_1 interface{} = b_1_loop
_ = b_1
var c_2 interface{} = c_2_loop
_ = c_2
var d_3 interface{} = d_3_loop
_ = d_3
var e_4 interface{} = e_4_loop
_ = e_4
var f_5 interface{} = f_5_loop
_ = f_5
var g_6 interface{} = g_6_loop
_ = g_6
var h_7 interface{} = h_7_loop
_ = h_7
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{e_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{g_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{h_7, gopurs_runtime.UnboxAny(pkg_Data_Unit.Get_unit())})}})}})}})}})}})}})}})})
}

func Call_tuple7(a_0_loop interface{}, b_1_loop interface{}, c_2_loop interface{}, d_3_loop interface{}, e_4_loop interface{}, f_5_loop interface{}, g_6_loop interface{}) gopurs_runtime.Value {
var a_0 interface{} = a_0_loop
_ = a_0
var b_1 interface{} = b_1_loop
_ = b_1
var c_2 interface{} = c_2_loop
_ = c_2
var d_3 interface{} = d_3_loop
_ = d_3
var e_4 interface{} = e_4_loop
_ = e_4
var f_5 interface{} = f_5_loop
_ = f_5
var g_6 interface{} = g_6_loop
_ = g_6
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{e_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{g_6, gopurs_runtime.UnboxAny(pkg_Data_Unit.Get_unit())})}})}})}})}})}})}})})
}

func Call_tuple6(a_0_loop interface{}, b_1_loop interface{}, c_2_loop interface{}, d_3_loop interface{}, e_4_loop interface{}, f_5_loop interface{}) gopurs_runtime.Value {
var a_0 interface{} = a_0_loop
_ = a_0
var b_1 interface{} = b_1_loop
_ = b_1
var c_2 interface{} = c_2_loop
_ = c_2
var d_3 interface{} = d_3_loop
_ = d_3
var e_4 interface{} = e_4_loop
_ = e_4
var f_5 interface{} = f_5_loop
_ = f_5
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{e_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{f_5, gopurs_runtime.UnboxAny(pkg_Data_Unit.Get_unit())})}})}})}})}})}})})
}

func Call_tuple5(a_0_loop interface{}, b_1_loop interface{}, c_2_loop interface{}, d_3_loop interface{}, e_4_loop interface{}) gopurs_runtime.Value {
var a_0 interface{} = a_0_loop
_ = a_0
var b_1 interface{} = b_1_loop
_ = b_1
var c_2 interface{} = c_2_loop
_ = c_2
var d_3 interface{} = d_3_loop
_ = d_3
var e_4 interface{} = e_4_loop
_ = e_4
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{e_4, gopurs_runtime.UnboxAny(pkg_Data_Unit.Get_unit())})}})}})}})}})})
}

func Call_tuple4(a_0_loop interface{}, b_1_loop interface{}, c_2_loop interface{}, d_3_loop interface{}) gopurs_runtime.Value {
var a_0 interface{} = a_0_loop
_ = a_0
var b_1 interface{} = b_1_loop
_ = b_1
var c_2 interface{} = c_2_loop
_ = c_2
var d_3 interface{} = d_3_loop
_ = d_3
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_3, gopurs_runtime.UnboxAny(pkg_Data_Unit.Get_unit())})}})}})}})})
}

func Call_tuple3(a_0_loop interface{}, b_1_loop interface{}, c_2_loop interface{}) gopurs_runtime.Value {
var a_0 interface{} = a_0_loop
_ = a_0
var b_1 interface{} = b_1_loop
_ = b_1
var c_2 interface{} = c_2_loop
_ = c_2
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_2, gopurs_runtime.UnboxAny(pkg_Data_Unit.Get_unit())})}})}})})
}

func Call_tuple2(a_0_loop interface{}, b_1_loop interface{}) gopurs_runtime.Value {
var a_0 interface{} = a_0_loop
_ = a_0
var b_1 interface{} = b_1_loop
_ = b_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_1, gopurs_runtime.UnboxAny(pkg_Data_Unit.Get_unit())})}})})
}

func Call_tuple10(a_0_loop interface{}, b_1_loop interface{}, c_2_loop interface{}, d_3_loop interface{}, e_4_loop interface{}, f_5_loop interface{}, g_6_loop interface{}, h_7_loop interface{}, i_8_loop interface{}, j_9_loop interface{}) gopurs_runtime.Value {
var a_0 interface{} = a_0_loop
_ = a_0
var b_1 interface{} = b_1_loop
_ = b_1
var c_2 interface{} = c_2_loop
_ = c_2
var d_3 interface{} = d_3_loop
_ = d_3
var e_4 interface{} = e_4_loop
_ = e_4
var f_5 interface{} = f_5_loop
_ = f_5
var g_6 interface{} = g_6_loop
_ = g_6
var h_7 interface{} = h_7_loop
_ = h_7
var i_8 interface{} = i_8_loop
_ = i_8
var j_9 interface{} = j_9_loop
_ = j_9
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{e_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{g_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{h_7, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{i_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{j_9, gopurs_runtime.UnboxAny(pkg_Data_Unit.Get_unit())})}})}})}})}})}})}})}})}})}})})
}

func Call_tuple1(a_0_loop interface{}) gopurs_runtime.Value {
var a_0 interface{} = a_0_loop
_ = a_0
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_0, gopurs_runtime.UnboxAny(pkg_Data_Unit.Get_unit())})})
}

func Call_over9(o_0_loop func(interface{}) interface{}, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var o_0 func(interface{}) interface{} = o_0_loop
_ = o_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(o_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))))), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1))})}})}})}})}})}})}})}})}})})
}

func Call_over8(o_0_loop func(interface{}) interface{}, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var o_0 func(interface{}) interface{} = o_0_loop
_ = o_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(o_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))))), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1))})}})}})}})}})}})}})}})})
}

func Call_over7(o_0_loop func(interface{}) interface{}, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var o_0 func(interface{}) interface{} = o_0_loop
_ = o_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(o_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))))), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1))})}})}})}})}})}})}})})
}

func Call_over6(o_0_loop func(interface{}) interface{}, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var o_0 func(interface{}) interface{} = o_0_loop
_ = o_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(o_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))))), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1))})}})}})}})}})}})})
}

func Call_over5(o_0_loop func(interface{}) interface{}, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var o_0 func(interface{}) interface{} = o_0_loop
_ = o_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(o_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))))), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1))})}})}})}})}})})
}

func Call_over4(o_0_loop func(interface{}) interface{}, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var o_0 func(interface{}) interface{} = o_0_loop
_ = o_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(o_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))))), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1))})}})}})}})})
}

func Call_over3(o_0_loop func(interface{}) interface{}, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var o_0 func(interface{}) interface{} = o_0_loop
_ = o_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(o_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))))), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1))})}})}})})
}

func Call_over2(o_0_loop func(interface{}) interface{}, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var o_0 func(interface{}) interface{} = o_0_loop
_ = o_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(o_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0))))), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1))})}})})
}

func Call_over10(o_0_loop func(interface{}) interface{}, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var o_0 func(interface{}) interface{} = o_0_loop
_ = o_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(o_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))))), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1))})}})}})}})}})}})}})}})}})}})})
}

func Call_over1(o_0_loop func(interface{}) interface{}, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var o_0 func(interface{}) interface{} = o_0_loop
_ = o_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(o_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0))))), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1))})})
}

func Call_get9(v_0_loop gopurs_runtime.Value) interface{} {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))
}

func Call_get8(v_0_loop gopurs_runtime.Value) interface{} {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))
}

func Call_get7(v_0_loop gopurs_runtime.Value) interface{} {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))
}

func Call_get6(v_0_loop gopurs_runtime.Value) interface{} {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))
}

func Call_get5(v_0_loop gopurs_runtime.Value) interface{} {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))
}

func Call_get4(v_0_loop gopurs_runtime.Value) interface{} {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))
}

func Call_get3(v_0_loop gopurs_runtime.Value) interface{} {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))
}

func Call_get2(v_0_loop gopurs_runtime.Value) interface{} {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1).UnsafePtr).V0))
}

func Call_get10(v_0_loop gopurs_runtime.Value) interface{} {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V1).UnsafePtr).V0))
}

func Call_get1(v_0_loop gopurs_runtime.Value) interface{} {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0))
}

func Call_curry8(z_0_loop interface{}, f_prime_1_loop func(gopurs_runtime.Value) interface{}, a_2_loop interface{}, b_3_loop interface{}, c_4_loop interface{}, d_5_loop interface{}, e_6_loop interface{}, f_7_loop interface{}, g_8_loop interface{}, h_9_loop interface{}) interface{} {
var z_0 interface{} = z_0_loop
_ = z_0
var f_prime_1 func(gopurs_runtime.Value) interface{} = f_prime_1_loop
_ = f_prime_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
var c_4 interface{} = c_4_loop
_ = c_4
var d_5 interface{} = d_5_loop
_ = d_5
var e_6 interface{} = e_6_loop
_ = e_6
var f_7 interface{} = f_7_loop
_ = f_7
var g_8 interface{} = g_8_loop
_ = g_8
var h_9 interface{} = h_9_loop
_ = h_9
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_prime_1(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{e_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{f_7, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{g_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{h_9, z_0})}})}})}})}})}})}})}})}))))
}

func Call_curry7(z_0_loop interface{}, f_prime_1_loop func(gopurs_runtime.Value) interface{}, a_2_loop interface{}, b_3_loop interface{}, c_4_loop interface{}, d_5_loop interface{}, e_6_loop interface{}, f_7_loop interface{}, g_8_loop interface{}) interface{} {
var z_0 interface{} = z_0_loop
_ = z_0
var f_prime_1 func(gopurs_runtime.Value) interface{} = f_prime_1_loop
_ = f_prime_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
var c_4 interface{} = c_4_loop
_ = c_4
var d_5 interface{} = d_5_loop
_ = d_5
var e_6 interface{} = e_6_loop
_ = e_6
var f_7 interface{} = f_7_loop
_ = f_7
var g_8 interface{} = g_8_loop
_ = g_8
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_prime_1(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{e_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{f_7, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{g_8, z_0})}})}})}})}})}})}})}))))
}

func Call_curry6(z_0_loop interface{}, f_prime_1_loop func(gopurs_runtime.Value) interface{}, a_2_loop interface{}, b_3_loop interface{}, c_4_loop interface{}, d_5_loop interface{}, e_6_loop interface{}, f_7_loop interface{}) interface{} {
var z_0 interface{} = z_0_loop
_ = z_0
var f_prime_1 func(gopurs_runtime.Value) interface{} = f_prime_1_loop
_ = f_prime_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
var c_4 interface{} = c_4_loop
_ = c_4
var d_5 interface{} = d_5_loop
_ = d_5
var e_6 interface{} = e_6_loop
_ = e_6
var f_7 interface{} = f_7_loop
_ = f_7
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_prime_1(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{e_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{f_7, z_0})}})}})}})}})}})}))))
}

func Call_curry5(z_0_loop interface{}, f_1_loop func(gopurs_runtime.Value) interface{}, a_2_loop interface{}, b_3_loop interface{}, c_4_loop interface{}, d_5_loop interface{}, e_6_loop interface{}) interface{} {
var z_0 interface{} = z_0_loop
_ = z_0
var f_1 func(gopurs_runtime.Value) interface{} = f_1_loop
_ = f_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
var c_4 interface{} = c_4_loop
_ = c_4
var d_5 interface{} = d_5_loop
_ = d_5
var e_6 interface{} = e_6_loop
_ = e_6
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{e_6, z_0})}})}})}})}})}))))
}

func Call_curry4(z_0_loop interface{}, f_1_loop func(gopurs_runtime.Value) interface{}, a_2_loop interface{}, b_3_loop interface{}, c_4_loop interface{}, d_5_loop interface{}) interface{} {
var z_0 interface{} = z_0_loop
_ = z_0
var f_1 func(gopurs_runtime.Value) interface{} = f_1_loop
_ = f_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
var c_4 interface{} = c_4_loop
_ = c_4
var d_5 interface{} = d_5_loop
_ = d_5
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{d_5, z_0})}})}})}})}))))
}

func Call_curry3(z_0_loop interface{}, f_1_loop func(gopurs_runtime.Value) interface{}, a_2_loop interface{}, b_3_loop interface{}, c_4_loop interface{}) interface{} {
var z_0 interface{} = z_0_loop
_ = z_0
var f_1 func(gopurs_runtime.Value) interface{} = f_1_loop
_ = f_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
var c_4 interface{} = c_4_loop
_ = c_4
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_3, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{c_4, z_0})}})}})}))))
}

func Call_curry2(z_0_loop interface{}, f_1_loop func(gopurs_runtime.Value) interface{}, a_2_loop interface{}, b_3_loop interface{}) interface{} {
var z_0 interface{} = z_0_loop
_ = z_0
var f_1 func(gopurs_runtime.Value) interface{} = f_1_loop
_ = f_1
var a_2 interface{} = a_2_loop
_ = a_2
var b_3 interface{} = b_3_loop
_ = b_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_3, z_0})}})}))))
}

func Call_curry1(z_0_loop interface{}, f_1_loop func(gopurs_runtime.Value) interface{}, a_2_loop interface{}) interface{} {
var z_0 interface{} = z_0_loop
_ = z_0
var f_1 func(gopurs_runtime.Value) interface{} = f_1_loop
_ = f_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, z_0})}))))
}
