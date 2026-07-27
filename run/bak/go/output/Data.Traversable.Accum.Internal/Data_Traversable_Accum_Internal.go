package Data_Traversable_Accum_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_StateR gopurs_runtime.Value
var once_StateR sync.Once
func Get_StateR() gopurs_runtime.Value {
	once_StateR.Do(func() {
		cache_StateR = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_StateR(x_0_box)
})
	})
	return cache_StateR
}

var cache_StateL gopurs_runtime.Value
var once_StateL sync.Once
func Get_StateL() gopurs_runtime.Value {
	once_StateL.Do(func() {
		cache_StateL = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_StateL(x_0_box)
})
	})
	return cache_StateL
}

var cache_stateR gopurs_runtime.Value
var once_stateR sync.Once
func Get_stateR() gopurs_runtime.Value {
	once_stateR.Do(func() {
		cache_stateR = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateR(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_stateR
}

var cache_stateR__func_func_interface____interface____interface____interface___2078380730 gopurs_runtime.Value
var once_stateR__func_func_interface____interface____interface____interface___2078380730 sync.Once
func Get_stateR__func_func_interface____interface____interface____interface___2078380730() gopurs_runtime.Value {
	once_stateR__func_func_interface____interface____interface____interface___2078380730.Do(func() {
		cache_stateR__func_func_interface____interface____interface____interface___2078380730 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateR__func_func_interface____interface____interface____interface___2078380730(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_stateR__func_func_interface____interface____interface____interface___2078380730
}

var cache_stateL gopurs_runtime.Value
var once_stateL sync.Once
func Get_stateL() gopurs_runtime.Value {
	once_stateL.Do(func() {
		cache_stateL = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateL(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_stateL
}

var cache_stateL__func_func_interface____interface____interface____interface___2078380730 gopurs_runtime.Value
var once_stateL__func_func_interface____interface____interface____interface___2078380730 sync.Once
func Get_stateL__func_func_interface____interface____interface____interface___2078380730() gopurs_runtime.Value {
	once_stateL__func_func_interface____interface____interface____interface___2078380730.Do(func() {
		cache_stateL__func_func_interface____interface____interface____interface___2078380730 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateL__func_func_interface____interface____interface____interface___2078380730(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_stateL__func_func_interface____interface____interface____interface___2078380730
}

var cache_functorStateR gopurs_runtime.Value
var once_functorStateR sync.Once
func Get_functorStateR() gopurs_runtime.Value {
	once_functorStateR.Do(func() {
		cache_functorStateR = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
}))))
	})
	return cache_functorStateR
}

var cache_functorStateL gopurs_runtime.Value
var once_functorStateL sync.Once
func Get_functorStateL() gopurs_runtime.Value {
	once_functorStateL.Do(func() {
		cache_functorStateL = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
}))))
	})
	return cache_functorStateL
}

var cache_applyStateR gopurs_runtime.Value
var once_applyStateR sync.Once
func Get_applyStateR() gopurs_runtime.Value {
	once_applyStateR.Do(func() {
		cache_applyStateR = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorStateR()
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(x_1, s_2)
_ = v_3_0
v1_4_1 := gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v1_4_1, "value"), gopurs_runtime.RecordGet(v_3_0, "value")))
}))))
	})
	return cache_applyStateR
}

var cache_applyStateL gopurs_runtime.Value
var once_applyStateL sync.Once
func Get_applyStateL() gopurs_runtime.Value {
	once_applyStateL.Do(func() {
		cache_applyStateL = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorStateL()
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(f_0, s_2)
_ = v_3_0
v1_4_1 := gopurs_runtime.Apply(x_1, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v_3_0, "value"), gopurs_runtime.RecordGet(v1_4_1, "value")))
}))))
	})
	return cache_applyStateL
}

var cache_applicativeStateR gopurs_runtime.Value
var once_applicativeStateR sync.Once
func Get_applicativeStateR() gopurs_runtime.Value {
	once_applicativeStateR.Do(func() {
		cache_applicativeStateR = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyStateR()
}), gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
}))))
	})
	return cache_applicativeStateR
}

var cache_applicativeStateL gopurs_runtime.Value
var once_applicativeStateL sync.Once
func Get_applicativeStateL() gopurs_runtime.Value {
	once_applicativeStateL.Do(func() {
		cache_applicativeStateL = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyStateL()
}), gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
}))))
	})
	return cache_applicativeStateL
}

func Call_StateR(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_StateL(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_stateR(v_0_loop func(interface{}) interface{}) gopurs_runtime.Value {
var v_0 func(interface{}) interface{} = v_0_loop
_ = v_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(v_0(gopurs_runtime.UnboxAny(arg0)))
})
}

func Call_stateR__func_func_interface____interface____interface____interface___2078380730(v_0_loop func(interface{}) interface{}) gopurs_runtime.Value {
var v_0 func(interface{}) interface{} = v_0_loop
_ = v_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(v_0(gopurs_runtime.UnboxAny(arg0)))
})
}

func Call_stateL(v_0_loop func(interface{}) interface{}) gopurs_runtime.Value {
var v_0 func(interface{}) interface{} = v_0_loop
_ = v_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(v_0(gopurs_runtime.UnboxAny(arg0)))
})
}

func Call_stateL__func_func_interface____interface____interface____interface___2078380730(v_0_loop func(interface{}) interface{}) gopurs_runtime.Value {
var v_0 func(interface{}) interface{} = v_0_loop
_ = v_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(v_0(gopurs_runtime.UnboxAny(arg0)))
})
}
