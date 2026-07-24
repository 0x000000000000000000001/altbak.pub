package Data_Traversable_Accum_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var StateR gopurs_runtime.Value
var once_StateR sync.Once
func Get_StateR() gopurs_runtime.Value {
	once_StateR.Do(func() {
		StateR = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return StateR
}

var StateL gopurs_runtime.Value
var once_StateL sync.Once
func Get_StateL() gopurs_runtime.Value {
	once_StateL.Do(func() {
		StateL = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return StateL
}

var stateR gopurs_runtime.Value
var once_stateR sync.Once
func Get_stateR() gopurs_runtime.Value {
	once_stateR.Do(func() {
		stateR = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0_loop
}()
})
	})
	return stateR
}

var stateL gopurs_runtime.Value
var once_stateL sync.Once
func Get_stateL() gopurs_runtime.Value {
	once_stateL.Do(func() {
		stateL = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0_loop
}()
})
	})
	return stateL
}

var functorStateR gopurs_runtime.Value
var once_functorStateR sync.Once
func Get_functorStateR() gopurs_runtime.Value {
	once_functorStateR.Do(func() {
		functorStateR = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
}))
	})
	return functorStateR
}

var functorStateL gopurs_runtime.Value
var once_functorStateL sync.Once
func Get_functorStateL() gopurs_runtime.Value {
	once_functorStateL.Do(func() {
		functorStateL = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
}))
	})
	return functorStateL
}

var applyStateR gopurs_runtime.Value
var once_applyStateR sync.Once
func Get_applyStateR() gopurs_runtime.Value {
	once_applyStateR.Do(func() {
		applyStateR = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(x_1, s_2)
_ = v_3_0
v1_4_1 := gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v1_4_1, "value"), gopurs_runtime.RecordGet(v_3_0, "value")))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorStateR()
}))
	})
	return applyStateR
}

var applyStateL gopurs_runtime.Value
var once_applyStateL sync.Once
func Get_applyStateL() gopurs_runtime.Value {
	once_applyStateL.Do(func() {
		applyStateL = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(f_0, s_2)
_ = v_3_0
v1_4_1 := gopurs_runtime.Apply(x_1, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v_3_0, "value"), gopurs_runtime.RecordGet(v1_4_1, "value")))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorStateL()
}))
	})
	return applyStateL
}

var applicativeStateR gopurs_runtime.Value
var once_applicativeStateR sync.Once
func Get_applicativeStateR() gopurs_runtime.Value {
	once_applicativeStateR.Do(func() {
		applicativeStateR = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyStateR()
}))
	})
	return applicativeStateR
}

var applicativeStateL gopurs_runtime.Value
var once_applicativeStateL sync.Once
func Get_applicativeStateL() gopurs_runtime.Value {
	once_applicativeStateL.Do(func() {
		applicativeStateL = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyStateL()
}))
	})
	return applicativeStateL
}




