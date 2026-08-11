package Data_Function

import (
	pkg_Data_Ord "gopurs/output/Data.Ord"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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

var cache_on__gopurs_runtime_Value_3980724833 gopurs_runtime.Value
var once_on__gopurs_runtime_Value_3980724833 sync.Once
func Get_on__gopurs_runtime_Value_3980724833() gopurs_runtime.Value {
	once_on__gopurs_runtime_Value_3980724833.Do(func() {
		cache_on__gopurs_runtime_Value_3980724833 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_on__gopurs_runtime_Value_3980724833(f_0_box, g_1_box, x_2_box, y_3_box)
})
	})
	return cache_on__gopurs_runtime_Value_3980724833
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

var cache_flip__gopurs_runtime_Value_3525081280 gopurs_runtime.Value
var once_flip__gopurs_runtime_Value_3525081280 sync.Once
func Get_flip__gopurs_runtime_Value_3525081280() gopurs_runtime.Value {
	once_flip__gopurs_runtime_Value_3525081280.Do(func() {
		cache_flip__gopurs_runtime_Value_3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__gopurs_runtime_Value_3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__gopurs_runtime_Value_3525081280
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

var cache_const__gopurs_runtime_Value_3524684546 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_3524684546 sync.Once
func Get_const__gopurs_runtime_Value_3524684546() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_3524684546.Do(func() {
		cache_const__gopurs_runtime_Value_3524684546 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_3524684546(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_3524684546
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

var cache_const__gopurs_runtime_Value_1426827922 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_1426827922 sync.Once
func Get_const__gopurs_runtime_Value_1426827922() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_1426827922.Do(func() {
		cache_const__gopurs_runtime_Value_1426827922 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_1426827922(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_1426827922
}

var cache_const__gopurs_runtime_Value_2569866098 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_2569866098 sync.Once
func Get_const__gopurs_runtime_Value_2569866098() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_2569866098.Do(func() {
		cache_const__gopurs_runtime_Value_2569866098 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_2569866098(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_2569866098
}

var cache_const__gopurs_runtime_Value_4181451586 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_4181451586 sync.Once
func Get_const__gopurs_runtime_Value_4181451586() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_4181451586.Do(func() {
		cache_const__gopurs_runtime_Value_4181451586 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_4181451586(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_4181451586
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

var cache_const__gopurs_runtime_Value_641934996 gopurs_runtime.Value
var once_const__gopurs_runtime_Value_641934996 sync.Once
func Get_const__gopurs_runtime_Value_641934996() gopurs_runtime.Value {
	once_const__gopurs_runtime_Value_641934996.Do(func() {
		cache_const__gopurs_runtime_Value_641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__gopurs_runtime_Value_641934996(a_0_box, v_1_box)
})
	})
	return cache_const__gopurs_runtime_Value_641934996
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

func Call_on__gopurs_runtime_Value_3980724833(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_flip__gopurs_runtime_Value_2673533882(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__gopurs_runtime_Value_3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_const__gopurs_runtime_Value_3524684546(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_const__gopurs_runtime_Value_1832354163(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_1426827922(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_2569866098(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__gopurs_runtime_Value_4181451586(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
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

func Call_const__gopurs_runtime_Value_641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
n_2_loop = gopurs_runtime.Int((n_2.IntVal) - (1))
acc_3_loop = gopurs_runtime.Apply(f_0, acc_3)
continue go__go_1_0_0
__t1 = gopurs_runtime.Value{}
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


