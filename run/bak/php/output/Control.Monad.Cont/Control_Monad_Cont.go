package Control_Monad_Cont

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var cache_withCont gopurs_runtime.Value
var once_withCont sync.Once
func Get_withCont() gopurs_runtime.Value {
	once_withCont.Do(func() {
		cache_withCont = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withCont(f_0_box)
})
	})
	return cache_withCont
}

var cache_runCont gopurs_runtime.Value
var once_runCont sync.Once
func Get_runCont() gopurs_runtime.Value {
	once_runCont.Do(func() {
		cache_runCont = gopurs_runtime.Func2(func(cc_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runCont(cc_0_box, k_1_box)
})
	})
	return cache_runCont
}

var cache_mapCont gopurs_runtime.Value
var once_mapCont sync.Once
func Get_mapCont() gopurs_runtime.Value {
	once_mapCont.Do(func() {
		cache_mapCont = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapCont(f_0_box)
})
	})
	return cache_mapCont
}

var cache_cont gopurs_runtime.Value
var once_cont sync.Once
func Get_cont() gopurs_runtime.Value {
	once_cont.Do(func() {
		cache_cont = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cont(f_0_box, c_1_box)
})
	})
	return cache_cont
}

func Call_withCont(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
__local_var_1_0 := gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"), pkg_Data_Identity.Get_Identity()), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), f_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"), pkg_Unsafe_Coerce.Get_unsafeCoerce())))
_ = __local_var_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply(__local_var_1_0, k_3))
})
}

func Call_runCont(cc_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cc_0 gopurs_runtime.Value = cc_0_loop
_ = cc_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(cc_0, gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pkg_Data_Identity.Get_Identity(), k_1))
}

func Call_mapCont(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
__local_var_1_0 := gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pkg_Data_Identity.Get_Identity(), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), f_0, pkg_Unsafe_Coerce.Get_unsafeCoerce()))
_ = __local_var_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(v_2, k_3))
})
}

func Call_cont(f_0_loop gopurs_runtime.Value, c_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pkg_Unsafe_Coerce.Get_unsafeCoerce(), c_1))
}


