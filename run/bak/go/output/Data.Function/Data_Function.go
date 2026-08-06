package Data_Function

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
)

var cache_lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		cache_lessThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 380165415) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_lessThanOrEq
}

var cache_on gopurs_runtime.Value
var once_on sync.Once
func Get_on() gopurs_runtime.Value {
	once_on.Do(func() {
		cache_on = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_on(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_on
}

var cache_on__gopurs_runtime_Value_2880499451 gopurs_runtime.Value
var once_on__gopurs_runtime_Value_2880499451 sync.Once
func Get_on__gopurs_runtime_Value_2880499451() gopurs_runtime.Value {
	once_on__gopurs_runtime_Value_2880499451.Do(func() {
		cache_on__gopurs_runtime_Value_2880499451 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_on__gopurs_runtime_Value_2880499451(f_0_box, g_1_box, x_2_box, y_3_box))
})
	})
	return cache_on__gopurs_runtime_Value_2880499451
}

var cache_on__gopurs_runtime_Value_1387348731 gopurs_runtime.Value
var once_on__gopurs_runtime_Value_1387348731 sync.Once
func Get_on__gopurs_runtime_Value_1387348731() gopurs_runtime.Value {
	once_on__gopurs_runtime_Value_1387348731.Do(func() {
		cache_on__gopurs_runtime_Value_1387348731 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_on__gopurs_runtime_Value_1387348731(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_on__gopurs_runtime_Value_1387348731
}

var cache_on__gopurs_runtime_Value_2620097339 gopurs_runtime.Value
var once_on__gopurs_runtime_Value_2620097339 sync.Once
func Get_on__gopurs_runtime_Value_2620097339() gopurs_runtime.Value {
	once_on__gopurs_runtime_Value_2620097339.Do(func() {
		cache_on__gopurs_runtime_Value_2620097339 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_on__gopurs_runtime_Value_2620097339(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_on__gopurs_runtime_Value_2620097339
}

var cache_flip gopurs_runtime.Value
var once_flip sync.Once
func Get_flip() gopurs_runtime.Value {
	once_flip.Do(func() {
		cache_flip = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip
}

var cache_flip__gopurs_runtime_Value_779556730 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_779556730 sync.Once
func Get_flip__gopurs_runtime_Value_779556730() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_779556730.Do(func() {
		cache_flip__gopurs_runtime_Value_779556730 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_flip__gopurs_runtime_Value_779556730(f_0_box, b_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(a_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_flip__gopurs_runtime_Value_779556730
}

var cache_flip__gopurs_runtime_Value_1667250810 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_1667250810 sync.Once
func Get_flip__gopurs_runtime_Value_1667250810() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_1667250810.Do(func() {
		cache_flip__gopurs_runtime_Value_1667250810 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_flip__gopurs_runtime_Value_1667250810(f_0_box, b_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(a_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_flip__gopurs_runtime_Value_1667250810
}

var cache_flip__gopurs_runtime_Value_1767337850 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_1767337850 sync.Once
func Get_flip__gopurs_runtime_Value_1767337850() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_1767337850.Do(func() {
		cache_flip__gopurs_runtime_Value_1767337850 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_1767337850(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_1767337850
}

var cache_flip__gopurs_runtime_Value_3766688826 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_3766688826 sync.Once
func Get_flip__gopurs_runtime_Value_3766688826() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_3766688826.Do(func() {
		cache_flip__gopurs_runtime_Value_3766688826 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_3766688826(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_3766688826
}

var cache_flip__gopurs_runtime_Value_968411546 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_968411546 sync.Once
func Get_flip__gopurs_runtime_Value_968411546() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_968411546.Do(func() {
		cache_flip__gopurs_runtime_Value_968411546 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_968411546(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_968411546
}

var cache_flip__gopurs_runtime_Value_2309783898 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_2309783898 sync.Once
func Get_flip__gopurs_runtime_Value_2309783898() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_2309783898.Do(func() {
		cache_flip__gopurs_runtime_Value_2309783898 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_2309783898(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_2309783898
}

var cache_flip__gopurs_runtime_Value_3843262778 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_3843262778 sync.Once
func Get_flip__gopurs_runtime_Value_3843262778() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_3843262778.Do(func() {
		cache_flip__gopurs_runtime_Value_3843262778 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_3843262778(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_3843262778
}

var cache_flip__gopurs_runtime_Value_1669991162 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_1669991162 sync.Once
func Get_flip__gopurs_runtime_Value_1669991162() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_1669991162.Do(func() {
		cache_flip__gopurs_runtime_Value_1669991162 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_flip__gopurs_runtime_Value_1669991162(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(b_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), a_2_box))
})
	})
	return cache_flip__gopurs_runtime_Value_1669991162
}

var cache_flip__gopurs_runtime_Value_3582434554 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_3582434554 sync.Once
func Get_flip__gopurs_runtime_Value_3582434554() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_3582434554.Do(func() {
		cache_flip__gopurs_runtime_Value_3582434554 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_3582434554(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_3582434554
}

var cache_flip__gopurs_runtime_Value_1282462010 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_1282462010 sync.Once
func Get_flip__gopurs_runtime_Value_1282462010() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_1282462010.Do(func() {
		cache_flip__gopurs_runtime_Value_1282462010 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_1282462010(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_1282462010
}

var cache_flip__gopurs_runtime_Value_325067610 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_325067610 sync.Once
func Get_flip__gopurs_runtime_Value_325067610() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_325067610.Do(func() {
		cache_flip__gopurs_runtime_Value_325067610 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_325067610(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_325067610
}

var cache_flip__gopurs_runtime_Value_3775495482 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_3775495482 sync.Once
func Get_flip__gopurs_runtime_Value_3775495482() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_3775495482.Do(func() {
		cache_flip__gopurs_runtime_Value_3775495482 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_3775495482(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_3775495482
}

var cache_flip__gopurs_runtime_Value_3541531802 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_3541531802 sync.Once
func Get_flip__gopurs_runtime_Value_3541531802() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_3541531802.Do(func() {
		cache_flip__gopurs_runtime_Value_3541531802 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_3541531802(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_3541531802
}

var cache_flip__gopurs_runtime_Value_1566949786 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_1566949786 sync.Once
func Get_flip__gopurs_runtime_Value_1566949786() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_1566949786.Do(func() {
		cache_flip__gopurs_runtime_Value_1566949786 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_1566949786(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_1566949786
}

var cache_flip__gopurs_runtime_Value_2673533882 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_2673533882 sync.Once
func Get_flip__gopurs_runtime_Value_2673533882() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_2673533882.Do(func() {
		cache_flip__gopurs_runtime_Value_2673533882 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_2673533882(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_2673533882
}

var cache_go__const gopurs_runtime.Value
var once_go__const sync.Once
func Get_go__const() gopurs_runtime.Value {
	once_go__const.Do(func() {
		cache_go__const = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_go__const(a_0_box, v_1_box)
})
	})
	return cache_go__const
}

var cache_const__gopurs_runtime_Value_467072791 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_467072791 sync.Once
func Get_const__gopurs_runtime_Value_467072791() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_467072791.Do(func() {
		cache_const__gopurs_runtime_Value_467072791 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_const__gopurs_runtime_Value_467072791(a_0_box.IntVal, v_1_box.IntVal))
})
	})
	return cache_const__gopurs_runtime_Value_467072791
}

var cache_const__gopurs_runtime_Value_1846763962 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_1846763962 sync.Once
func Get_const__gopurs_runtime_Value_1846763962() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_1846763962.Do(func() {
		cache_const__gopurs_runtime_Value_1846763962 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_const__gopurs_runtime_Value_1846763962(a_0_box.IntVal, v_1_box))
})
	})
	return cache_const__gopurs_runtime_Value_1846763962
}

var cache_const__gopurs_runtime_Value_2731150322 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_2731150322 sync.Once
func Get_const__gopurs_runtime_Value_2731150322() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_2731150322.Do(func() {
		cache_const__gopurs_runtime_Value_2731150322 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_const__gopurs_runtime_Value_2731150322((a_0_box.IntVal) != (0), v_1_box))
})
	})
	return cache_const__gopurs_runtime_Value_2731150322
}

var cache_const__gopurs_runtime_Value_2390202835 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_2390202835 sync.Once
func Get_const__gopurs_runtime_Value_2390202835() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_2390202835.Do(func() {
		cache_const__gopurs_runtime_Value_2390202835 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_2390202835(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_2390202835
}

var cache_const__gopurs_runtime_Value_239908602 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_239908602 sync.Once
func Get_const__gopurs_runtime_Value_239908602() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_239908602.Do(func() {
		cache_const__gopurs_runtime_Value_239908602 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_239908602(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_239908602
}

var cache_const__gopurs_runtime_Value_1205362034 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_1205362034 sync.Once
func Get_const__gopurs_runtime_Value_1205362034() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_1205362034.Do(func() {
		cache_const__gopurs_runtime_Value_1205362034 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_1205362034(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_1205362034
}

var cache_const__gopurs_runtime_Value_3707916826 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_3707916826 sync.Once
func Get_const__gopurs_runtime_Value_3707916826() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_3707916826.Do(func() {
		cache_const__gopurs_runtime_Value_3707916826 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_3707916826(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_3707916826
}

var cache_const__gopurs_runtime_Value_1548733586 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_1548733586 sync.Once
func Get_const__gopurs_runtime_Value_1548733586() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_1548733586.Do(func() {
		cache_const__gopurs_runtime_Value_1548733586 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_1548733586(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_1548733586
}

var cache_const__gopurs_runtime_Value_2304792434 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_2304792434 sync.Once
func Get_const__gopurs_runtime_Value_2304792434() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_2304792434.Do(func() {
		cache_const__gopurs_runtime_Value_2304792434 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_2304792434(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_2304792434
}

var cache_const__gopurs_runtime_Value_3128343706 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_3128343706 sync.Once
func Get_const__gopurs_runtime_Value_3128343706() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_3128343706.Do(func() {
		cache_const__gopurs_runtime_Value_3128343706 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_3128343706(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_3128343706
}

var cache_const__gopurs_runtime_Value_752911026 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_752911026 sync.Once
func Get_const__gopurs_runtime_Value_752911026() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_752911026.Do(func() {
		cache_const__gopurs_runtime_Value_752911026 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_752911026(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_752911026
}

var cache_const__gopurs_runtime_Value_131698394 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_131698394 sync.Once
func Get_const__gopurs_runtime_Value_131698394() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_131698394.Do(func() {
		cache_const__gopurs_runtime_Value_131698394 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_131698394(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_131698394
}

var cache_const__gopurs_runtime_Value_2608703442 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_2608703442 sync.Once
func Get_const__gopurs_runtime_Value_2608703442() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_2608703442.Do(func() {
		cache_const__gopurs_runtime_Value_2608703442 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_2608703442(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_2608703442
}

var cache_const__gopurs_runtime_Value_1832354163 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_1832354163 sync.Once
func Get_const__gopurs_runtime_Value_1832354163() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_1832354163.Do(func() {
		cache_const__gopurs_runtime_Value_1832354163 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_1832354163(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_1832354163
}

var cache_const__gopurs_runtime_Value_3978148602 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_3978148602 sync.Once
func Get_const__gopurs_runtime_Value_3978148602() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_3978148602.Do(func() {
		cache_const__gopurs_runtime_Value_3978148602 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_3978148602(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_3978148602
}

var cache_const__gopurs_runtime_Value_2694363354 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_2694363354 sync.Once
func Get_const__gopurs_runtime_Value_2694363354() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_2694363354.Do(func() {
		cache_const__gopurs_runtime_Value_2694363354 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_2694363354(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_2694363354
}

var cache_const__gopurs_runtime_Value_1496134642 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_1496134642 sync.Once
func Get_const__gopurs_runtime_Value_1496134642() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_1496134642.Do(func() {
		cache_const__gopurs_runtime_Value_1496134642 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_1496134642(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_1496134642
}

var cache_applyN gopurs_runtime.Value
var once_applyN sync.Once
func Get_applyN() gopurs_runtime.Value {
	once_applyN.Do(func() {
		cache_applyN = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyN(f_0_box)
})
	})
	return cache_applyN
}

var cache_applyFlipped gopurs_runtime.Value
var once_applyFlipped sync.Once
func Get_applyFlipped() gopurs_runtime.Value {
	once_applyFlipped.Do(func() {
		cache_applyFlipped = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyFlipped(x_0_box, f_1_box)
})
	})
	return cache_applyFlipped
}

var cache_apply gopurs_runtime.Value
var once_apply sync.Once
func Get_apply() gopurs_runtime.Value {
	once_apply.Do(func() {
		cache_apply = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply(f_0_box, x_1_box)
})
	})
	return cache_apply
}

var cache_apply__gopurs_runtime_Value_458711162 gopurs_runtime.Value
var once_apply__gopurs_runtime_Value_458711162 sync.Once
func Get_apply__gopurs_runtime_Value_458711162() gopurs_runtime.Value {
	once_apply__gopurs_runtime_Value_458711162.Do(func() {
		cache_apply__gopurs_runtime_Value_458711162 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__gopurs_runtime_Value_458711162(f_0_box, x_1_box)
})
	})
	return cache_apply__gopurs_runtime_Value_458711162
}

func Call_on(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_on__gopurs_runtime_Value_2880499451(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) bool {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return (gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3)).IntVal) != (0)
}

func Call_on__gopurs_runtime_Value_1387348731(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_on__gopurs_runtime_Value_2620097339(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(g_1, x_2), gopurs_runtime.Apply(g_1, y_3))
}

func Call_flip(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_779556730(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 []gopurs_runtime.Value = a_2_loop
_ = a_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0, gopurs_runtime.Array(a_2), b_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_flip__gopurs_runtime_Value_1667250810(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 []gopurs_runtime.Value = a_2_loop
_ = a_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0, gopurs_runtime.Array(a_2), b_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_flip__gopurs_runtime_Value_1767337850(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_3766688826(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_968411546(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_2309783898(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_3843262778(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_1669991162(f_0_loop gopurs_runtime.Value, b_1_loop []gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 []gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(f_0, a_2, gopurs_runtime.Array(b_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_flip__gopurs_runtime_Value_3582434554(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_1282462010(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_325067610(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_3775495482(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_3541531802(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_1566949786(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_2673533882(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_go__const(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_467072791(a_0_loop int64, v_1_loop int64) int64 {
var a_0 int64 = a_0_loop
_ = a_0
var v_1 int64 = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_1846763962(a_0_loop int64, v_1_loop gopurs_runtime.Value) int64 {
var a_0 int64 = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_2731150322(a_0_loop bool, v_1_loop gopurs_runtime.Value) bool {
var a_0 bool = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_2390202835(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Bool((a_0.IntVal) != (0))
}

func Call_const__gopurs_runtime_Value_239908602(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_1205362034(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_3707916826(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_1548733586(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_2304792434(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_3128343706(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_752911026(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_131698394(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_2608703442(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_1832354163(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_3978148602(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_2694363354(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_1496134642(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_applyN(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_0 gopurs_runtime.Value
go__go_1_0_0 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var n_2_loop gopurs_runtime.Value = n_2_loop_val
var acc_3_loop gopurs_runtime.Value = acc_3_loop_val
go__go_1_0_0:
for {
if false { continue go__go_1_0_0 }
var n_2 gopurs_runtime.Value = n_2_loop
_ = n_2
var acc_3 gopurs_runtime.Value = acc_3_loop
_ = acc_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), n_2, gopurs_runtime.Int(0)).IntVal) != (0) {
__t1 = acc_3
goto end_branch_1
} else {

}
}
{
if true {
n_2_loop = gopurs_runtime.Int((n_2.IntVal) - (1))
acc_3_loop = gopurs_runtime.Apply(f_0, acc_3)
continue go__go_1_0_0
__t1 = gopurs_runtime.Value{}
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
}()
})
})
return go__go_1_0_0
}

func Call_applyFlipped(x_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(f_1, x_0)
}

func Call_apply(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, x_1)
}

func Call_apply__gopurs_runtime_Value_458711162(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, x_1)
}


