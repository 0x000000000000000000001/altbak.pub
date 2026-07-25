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
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool(((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415)) != (true))
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

var cache_const_ gopurs_runtime.Value
var once_const_ sync.Once
func Get_const_() gopurs_runtime.Value {
	once_const_.Do(func() {
		cache_const_ = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const_(a_0_box, v_1_box)
})
	})
	return cache_const_
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

func Call_flip(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_const_(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_applyN(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(n_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
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
continue go__1_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return go__1_0
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


