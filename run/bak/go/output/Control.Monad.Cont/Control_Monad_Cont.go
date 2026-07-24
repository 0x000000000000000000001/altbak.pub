package Control_Monad_Cont

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var withCont gopurs_runtime.Value
var once_withCont sync.Once
func Get_withCont() gopurs_runtime.Value {
	once_withCont.Do(func() {
		withCont = gopurs_runtime.Func3(Call_withCont)
	})
	return withCont
}

var runCont gopurs_runtime.Value
var once_runCont sync.Once
func Get_runCont() gopurs_runtime.Value {
	once_runCont.Do(func() {
		runCont = gopurs_runtime.Func2(Call_runCont)
	})
	return runCont
}

var mapCont gopurs_runtime.Value
var once_mapCont sync.Once
func Get_mapCont() gopurs_runtime.Value {
	once_mapCont.Do(func() {
		mapCont = gopurs_runtime.Func3(Call_mapCont)
	})
	return mapCont
}

var cont gopurs_runtime.Value
var once_cont sync.Once
func Get_cont() gopurs_runtime.Value {
	once_cont.Do(func() {
		cont = gopurs_runtime.Func2(Call_cont)
	})
	return cont
}

func Call_withCont(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(v_1_loop, gopurs_runtime.Apply(f_0_loop, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_2_loop, x_3)
})))
}

func Call_runCont(cc_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cc_0 gopurs_runtime.Value = cc_0_loop
_ = cc_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(cc_0_loop, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1_loop, x_2)
}))
}

func Call_mapCont(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(f_0_loop, gopurs_runtime.Apply(v_1_loop, k_2_loop))
}

func Call_cont(f_0_loop gopurs_runtime.Value, c_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
return gopurs_runtime.Apply(f_0_loop, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_1_loop, x_2)
}))
}


