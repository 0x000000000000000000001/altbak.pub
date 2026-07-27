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
return Call_withCont(func(inner_arg0 func(interface{}) interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg1)))
})
})
	})
	return cache_withCont
}

var cache_runCont gopurs_runtime.Value
var once_runCont sync.Once
func Get_runCont() gopurs_runtime.Value {
	once_runCont.Do(func() {
		cache_runCont = gopurs_runtime.Func2(func(cc_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_runCont(func(inner_arg0 func(interface{}) interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(cc_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
})))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(k_1_box, gopurs_runtime.Any(inner_arg0)))
}))
})
	})
	return cache_runCont
}

var cache_mapCont gopurs_runtime.Value
var once_mapCont sync.Once
func Get_mapCont() gopurs_runtime.Value {
	once_mapCont.Do(func() {
		cache_mapCont = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_mapCont(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 func(interface{}) interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_1_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
})))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(k_2_box, gopurs_runtime.Any(inner_arg0)))
}))
})
	})
	return cache_mapCont
}

var cache_cont gopurs_runtime.Value
var once_cont sync.Once
func Get_cont() gopurs_runtime.Value {
	once_cont.Do(func() {
		cache_cont = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_cont(func(inner_arg0 func(interface{}) interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
})))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(c_1_box, gopurs_runtime.Any(inner_arg0)))
}))
})
	})
	return cache_cont
}

func Call_withCont(f_0_loop func(func(interface{}) interface{}, interface{}) interface{}) gopurs_runtime.Value {
var f_0 func(func(interface{}) interface{}, interface{}) interface{} = f_0_loop
_ = f_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"), pkg_Data_Identity.Get_Identity())
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"), pkg_Unsafe_Coerce.Get_unsafeCoerce())
_ = __local_var_2_1
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(arg1)))
}), gopurs_runtime.Apply(__local_var_2_1, k_4))))
})
}

func Call_runCont(cc_0_loop func(func(interface{}) interface{}) interface{}, k_1_loop func(interface{}) interface{}) interface{} {
var cc_0 func(func(interface{}) interface{}) interface{} = cc_0_loop
_ = cc_0
var k_1 func(interface{}) interface{} = k_1_loop
_ = k_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(cc_0(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(k_1(gopurs_runtime.UnboxAny(x_2)))
}), gopurs_runtime.Any(inner_arg0)))
})))
}

func Call_mapCont(f_0_loop func(interface{}) interface{}, v_1_loop func(func(interface{}) interface{}) interface{}, k_2_loop func(interface{}) interface{}) interface{} {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
var v_1 func(func(interface{}) interface{}) interface{} = v_1_loop
_ = v_1
var k_2 func(interface{}) interface{} = k_2_loop
_ = k_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_1(k_2))))))
}

func Call_cont(f_0_loop func(func(interface{}) interface{}) interface{}, c_1_loop func(interface{}) interface{}) interface{} {
var f_0 func(func(interface{}) interface{}) interface{} = f_0_loop
_ = f_0
var c_1 func(interface{}) interface{} = c_1_loop
_ = c_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(c_1(gopurs_runtime.UnboxAny(x_2)))
}), gopurs_runtime.Any(inner_arg0)))
})))
}
