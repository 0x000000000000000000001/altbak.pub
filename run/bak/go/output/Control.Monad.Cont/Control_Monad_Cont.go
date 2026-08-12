package Control_Monad_Cont

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
		cache_mapCont = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapCont(f_0_box, v_1_box, k_2_box)
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

var cache_mapContT__1876830945 gopurs_runtime.Value
var once_mapContT__1876830945 sync.Once
func Get_mapContT__1876830945() gopurs_runtime.Value {
	once_mapContT__1876830945.Do(func() {
		cache_mapContT__1876830945 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapContT__1876830945(f_0_box, v_1_box, k_2_box)
})
	})
	return cache_mapContT__1876830945
}

var cache_runContT__3370980748 gopurs_runtime.Value
var once_runContT__3370980748 sync.Once
func Get_runContT__3370980748() gopurs_runtime.Value {
	once_runContT__3370980748.Do(func() {
		cache_runContT__3370980748 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runContT__3370980748(v_0_box, k_1_box)
})
	})
	return cache_runContT__3370980748
}

var cache_withContT__3360252697 gopurs_runtime.Value
var once_withContT__3360252697 sync.Once
func Get_withContT__3360252697() gopurs_runtime.Value {
	once_withContT__3360252697.Do(func() {
		cache_withContT__3360252697 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withContT__3360252697(f_0_box, v_1_box, k_2_box)
})
	})
	return cache_withContT__3360252697
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_unwrap__3267718003 gopurs_runtime.Value
var once_unwrap__3267718003 sync.Once
func Get_unwrap__3267718003() gopurs_runtime.Value {
	once_unwrap__3267718003.Do(func() {
		cache_unwrap__3267718003 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__3267718003(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__3267718003
}

func Call_withCont(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"), pkg_Data_Identity.Get_Identity())
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"), pkg_Unsafe_Coerce.Get_unsafeCoerce())
_ = __local_var_2_3
__local_var_2_2 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_2_3, x_3))
})
_ = __local_var_2_2
__local_var_1_0 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, gopurs_runtime.Apply(__local_var_2_2, x_3))
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Apply(__local_var_1_0, k_3))
})
})
}

func Call_runCont(cc_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cc_0 gopurs_runtime.Value = cc_0_loop
_ = cc_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(cc_0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_1, x_2)
}))
}

func Call_mapCont(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, k_2))
}

func Call_cont(f_0_loop gopurs_runtime.Value, c_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_1, x_2)
}))
}

func Call_mapContT__1876830945(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, k_2))
}

func Call_runContT__3370980748(v_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(v_0, k_1)
}

func Call_withContT__3360252697(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, k_2))
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_unwrap__3267718003(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}


