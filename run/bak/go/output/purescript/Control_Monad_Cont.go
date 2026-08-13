package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Control_Monad_Cont_unwrap gopurs_runtime.Value
var once_Control_Monad_Cont_unwrap sync.Once
func Get_Control_Monad_Cont_unwrap() gopurs_runtime.Value {
	once_Control_Monad_Cont_unwrap.Do(func() {
		cache_Control_Monad_Cont_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Control_Monad_Cont_unwrap
}

var cache_Control_Monad_Cont_withCont gopurs_runtime.Value
var once_Control_Monad_Cont_withCont sync.Once
func Get_Control_Monad_Cont_withCont() gopurs_runtime.Value {
	once_Control_Monad_Cont_withCont.Do(func() {
		cache_Control_Monad_Cont_withCont = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_withCont(f_0_box, v_1_box, k_2_box)
})
	})
	return cache_Control_Monad_Cont_withCont
}

var cache_Control_Monad_Cont_runCont gopurs_runtime.Value
var once_Control_Monad_Cont_runCont sync.Once
func Get_Control_Monad_Cont_runCont() gopurs_runtime.Value {
	once_Control_Monad_Cont_runCont.Do(func() {
		cache_Control_Monad_Cont_runCont = gopurs_runtime.Func2(func(cc_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_runCont(cc_0_box, k_1_box)
})
	})
	return cache_Control_Monad_Cont_runCont
}

var cache_Control_Monad_Cont_mapCont gopurs_runtime.Value
var once_Control_Monad_Cont_mapCont sync.Once
func Get_Control_Monad_Cont_mapCont() gopurs_runtime.Value {
	once_Control_Monad_Cont_mapCont.Do(func() {
		cache_Control_Monad_Cont_mapCont = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_mapCont(f_0_box, v_1_box, k_2_box)
})
	})
	return cache_Control_Monad_Cont_mapCont
}

var cache_Control_Monad_Cont_cont gopurs_runtime.Value
var once_Control_Monad_Cont_cont sync.Once
func Get_Control_Monad_Cont_cont() gopurs_runtime.Value {
	once_Control_Monad_Cont_cont.Do(func() {
		cache_Control_Monad_Cont_cont = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_cont(f_0_box, c_1_box)
})
	})
	return cache_Control_Monad_Cont_cont
}

func Call_Control_Monad_Cont_withCont(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_2, x_3)
}))
_ = __local_var_3_0
return gopurs_runtime.Apply(v_1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_0, x_4)
}))
}

func Call_Control_Monad_Cont_runCont(cc_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cc_0 gopurs_runtime.Value = cc_0_loop
_ = cc_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(cc_0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, x_2)
}))
}

func Call_Control_Monad_Cont_mapCont(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, k_2))
}

func Call_Control_Monad_Cont_cont(f_0_loop gopurs_runtime.Value, c_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_1, x_2)
}))
}


