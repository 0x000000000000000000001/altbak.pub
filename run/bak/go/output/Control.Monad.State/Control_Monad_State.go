package Control_Monad_State

import (
	pkg_Control_Monad_State_Trans "gopurs/output/Control.Monad.State.Trans"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_withState gopurs_runtime.Value
var once_withState sync.Once
func Get_withState() gopurs_runtime.Value {
	once_withState.Do(func() {
		cache_withState = pkg_Control_Monad_State_Trans.Get_withStateT()
	})
	return cache_withState
}

var cache_runState gopurs_runtime.Value
var once_runState sync.Once
func Get_runState() gopurs_runtime.Value {
	once_runState.Do(func() {
		cache_runState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_runState(v_0_box, x_1_box))}
})
	})
	return cache_runState
}

var cache_mapState gopurs_runtime.Value
var once_mapState sync.Once
func Get_mapState() gopurs_runtime.Value {
	once_mapState.Do(func() {
		cache_mapState = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_mapState(f_0_box, v_1_box, x_2_box))}
})
	})
	return cache_mapState
}

var cache_execState gopurs_runtime.Value
var once_execState sync.Once
func Get_execState() gopurs_runtime.Value {
	once_execState.Do(func() {
		cache_execState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_execState(v_0_box, s_1_box)
})
	})
	return cache_execState
}

var cache_evalState gopurs_runtime.Value
var once_evalState sync.Once
func Get_evalState() gopurs_runtime.Value {
	once_evalState.Do(func() {
		cache_evalState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_evalState(v_0_box, s_1_box)
})
	})
	return cache_evalState
}

var cache_mapStateT__1987836370 gopurs_runtime.Value
var once_mapStateT__1987836370 sync.Once
func Get_mapStateT__1987836370() gopurs_runtime.Value {
	once_mapStateT__1987836370.Do(func() {
		cache_mapStateT__1987836370 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapStateT__1987836370(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_mapStateT__1987836370
}

var cache_mapStateT__3766252210 gopurs_runtime.Value
var once_mapStateT__3766252210 sync.Once
func Get_mapStateT__3766252210() gopurs_runtime.Value {
	once_mapStateT__3766252210.Do(func() {
		cache_mapStateT__3766252210 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapStateT__3766252210(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_mapStateT__3766252210
}

var cache_withStateT__2250856667 gopurs_runtime.Value
var once_withStateT__2250856667 sync.Once
func Get_withStateT__2250856667() gopurs_runtime.Value {
	once_withStateT__2250856667.Do(func() {
		cache_withStateT__2250856667 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_withStateT__2250856667(f_0_box, v_1_box, x_2_box))}
})
	})
	return cache_withStateT__2250856667
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

func Call_runState(v_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(v_0, x_1))
}

func Call_mapState(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2)))
}

func Call_execState(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(v_0, s_1).UnsafePtr).V1
}

func Call_evalState(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(v_0, s_1).UnsafePtr).V0
}

func Call_mapStateT__1987836370(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_mapStateT__3766252210(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}

func Call_withStateT__2250856667(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) *pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2)))
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


