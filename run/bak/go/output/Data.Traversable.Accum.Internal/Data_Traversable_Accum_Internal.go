package Data_Traversable_Accum_Internal

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
return Call_stateR(v_0_box)
})
	})
	return cache_stateR
}

var cache_stateR__gopurs_runtime_Value_1334064830 gopurs_runtime.Value
var once_stateR__gopurs_runtime_Value_1334064830 sync.Once
func Get_stateR__gopurs_runtime_Value_1334064830() gopurs_runtime.Value {
	once_stateR__gopurs_runtime_Value_1334064830.Do(func() {
		cache_stateR__gopurs_runtime_Value_1334064830 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateR__gopurs_runtime_Value_1334064830(v_0_box)
})
	})
	return cache_stateR__gopurs_runtime_Value_1334064830
}

var cache_stateL gopurs_runtime.Value
var once_stateL sync.Once
func Get_stateL() gopurs_runtime.Value {
	once_stateL.Do(func() {
		cache_stateL = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateL(v_0_box)
})
	})
	return cache_stateL
}

var cache_stateL__gopurs_runtime_Value_1334064830 gopurs_runtime.Value
var once_stateL__gopurs_runtime_Value_1334064830 sync.Once
func Get_stateL__gopurs_runtime_Value_1334064830() gopurs_runtime.Value {
	once_stateL__gopurs_runtime_Value_1334064830.Do(func() {
		cache_stateL__gopurs_runtime_Value_1334064830 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateL__gopurs_runtime_Value_1334064830(v_0_box)
})
	})
	return cache_stateL__gopurs_runtime_Value_1334064830
}

var cache_functorStateR gopurs_runtime.Value
var once_functorStateR sync.Once
func Get_functorStateR() gopurs_runtime.Value {
	once_functorStateR.Do(func() {
		cache_functorStateR = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
}))
	})
	return cache_functorStateR
}

var cache_functorStateR__gopurs_runtime_Value_830241200 gopurs_runtime.Value
var once_functorStateR__gopurs_runtime_Value_830241200 sync.Once
func Get_functorStateR__gopurs_runtime_Value_830241200() gopurs_runtime.Value {
	once_functorStateR__gopurs_runtime_Value_830241200.Do(func() {
		cache_functorStateR__gopurs_runtime_Value_830241200 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
}))
	})
	return cache_functorStateR__gopurs_runtime_Value_830241200
}

var cache_functorStateL gopurs_runtime.Value
var once_functorStateL sync.Once
func Get_functorStateL() gopurs_runtime.Value {
	once_functorStateL.Do(func() {
		cache_functorStateL = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
}))
	})
	return cache_functorStateL
}

var cache_functorStateL__gopurs_runtime_Value_830241200 gopurs_runtime.Value
var once_functorStateL__gopurs_runtime_Value_830241200 sync.Once
func Get_functorStateL__gopurs_runtime_Value_830241200() gopurs_runtime.Value {
	once_functorStateL__gopurs_runtime_Value_830241200.Do(func() {
		cache_functorStateL__gopurs_runtime_Value_830241200 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
}))
	})
	return cache_functorStateL__gopurs_runtime_Value_830241200
}

var cache_applyStateR gopurs_runtime.Value
var once_applyStateR sync.Once
func Get_applyStateR() gopurs_runtime.Value {
	once_applyStateR.Do(func() {
		cache_applyStateR = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorStateR()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(x_1, s_2)
_ = v_3_0
v1_4_1 := gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v1_4_1, "value"), gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
}))
	})
	return cache_applyStateR
}

var cache_applyStateR__gopurs_runtime_Value_1243455060 gopurs_runtime.Value
var once_applyStateR__gopurs_runtime_Value_1243455060 sync.Once
func Get_applyStateR__gopurs_runtime_Value_1243455060() gopurs_runtime.Value {
	once_applyStateR__gopurs_runtime_Value_1243455060.Do(func() {
		cache_applyStateR__gopurs_runtime_Value_1243455060 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorStateR__gopurs_runtime_Value_830241200()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(x_1, s_2)
_ = v_3_0
v1_4_1 := gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v1_4_1, "value"), gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
}))
	})
	return cache_applyStateR__gopurs_runtime_Value_1243455060
}

var cache_applyStateL gopurs_runtime.Value
var once_applyStateL sync.Once
func Get_applyStateL() gopurs_runtime.Value {
	once_applyStateL.Do(func() {
		cache_applyStateL = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorStateL()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(f_0, s_2)
_ = v_3_0
v1_4_1 := gopurs_runtime.Apply(x_1, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v_3_0, "value"), gopurs_runtime.RecordGet(v1_4_1, "value")))
})
})
}))
	})
	return cache_applyStateL
}

var cache_applyStateL__gopurs_runtime_Value_1243455060 gopurs_runtime.Value
var once_applyStateL__gopurs_runtime_Value_1243455060 sync.Once
func Get_applyStateL__gopurs_runtime_Value_1243455060() gopurs_runtime.Value {
	once_applyStateL__gopurs_runtime_Value_1243455060.Do(func() {
		cache_applyStateL__gopurs_runtime_Value_1243455060 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorStateL__gopurs_runtime_Value_830241200()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(f_0, s_2)
_ = v_3_0
v1_4_1 := gopurs_runtime.Apply(x_1, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v_3_0, "value"), gopurs_runtime.RecordGet(v1_4_1, "value")))
})
})
}))
	})
	return cache_applyStateL__gopurs_runtime_Value_1243455060
}

var cache_applicativeStateR gopurs_runtime.Value
var once_applicativeStateR sync.Once
func Get_applicativeStateR() gopurs_runtime.Value {
	once_applicativeStateR.Do(func() {
		cache_applicativeStateR = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyStateR()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
}))
	})
	return cache_applicativeStateR
}

var cache_applicativeStateR__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__2039640491 gopurs_runtime.Value
var once_applicativeStateR__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__2039640491 sync.Once
func Get_applicativeStateR__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__2039640491() gopurs_runtime.Value {
	once_applicativeStateR__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__2039640491.Do(func() {
		cache_applicativeStateR__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__2039640491 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyStateR()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
})})}
	})
	return cache_applicativeStateR__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__2039640491
}

var cache_applicativeStateL gopurs_runtime.Value
var once_applicativeStateL sync.Once
func Get_applicativeStateL() gopurs_runtime.Value {
	once_applicativeStateL.Do(func() {
		cache_applicativeStateL = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyStateL()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
}))
	})
	return cache_applicativeStateL
}

var cache_applicativeStateL__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__2039640491 gopurs_runtime.Value
var once_applicativeStateL__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__2039640491 sync.Once
func Get_applicativeStateL__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__2039640491() gopurs_runtime.Value {
	once_applicativeStateL__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__2039640491.Do(func() {
		cache_applicativeStateL__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__2039640491 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyStateL()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
})})}
	})
	return cache_applicativeStateL__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__2039640491
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

func Call_stateR(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_stateR__gopurs_runtime_Value_1334064830(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_stateL(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_stateL__gopurs_runtime_Value_1334064830(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}


