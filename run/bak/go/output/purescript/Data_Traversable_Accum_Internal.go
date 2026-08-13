package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Traversable_Accum_Internal_StateR gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_StateR sync.Once
func Get_Data_Traversable_Accum_Internal_StateR() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_StateR.Do(func() {
		cache_Data_Traversable_Accum_Internal_StateR = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Accum_Internal_StateR(x_0_box)
})
	})
	return cache_Data_Traversable_Accum_Internal_StateR
}

var cache_Data_Traversable_Accum_Internal_StateL gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_StateL sync.Once
func Get_Data_Traversable_Accum_Internal_StateL() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_StateL.Do(func() {
		cache_Data_Traversable_Accum_Internal_StateL = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Accum_Internal_StateL(x_0_box)
})
	})
	return cache_Data_Traversable_Accum_Internal_StateL
}

var cache_Data_Traversable_Accum_Internal_stateR gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_stateR sync.Once
func Get_Data_Traversable_Accum_Internal_stateR() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_stateR.Do(func() {
		cache_Data_Traversable_Accum_Internal_stateR = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Accum_Internal_stateR(v_0_box)
})
	})
	return cache_Data_Traversable_Accum_Internal_stateR
}

var cache_Data_Traversable_Accum_Internal_stateL gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_stateL sync.Once
func Get_Data_Traversable_Accum_Internal_stateL() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_stateL.Do(func() {
		cache_Data_Traversable_Accum_Internal_stateL = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Accum_Internal_stateL(v_0_box)
})
	})
	return cache_Data_Traversable_Accum_Internal_stateL
}

var cache_Data_Traversable_Accum_Internal_functorStateR gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_functorStateR sync.Once
func Get_Data_Traversable_Accum_Internal_functorStateR() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_functorStateR.Do(func() {
		cache_Data_Traversable_Accum_Internal_functorStateR = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_functorStateR
}

var cache_Data_Traversable_Accum_Internal_functorStateL gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_functorStateL sync.Once
func Get_Data_Traversable_Accum_Internal_functorStateL() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_functorStateL.Do(func() {
		cache_Data_Traversable_Accum_Internal_functorStateL = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_functorStateL
}

var cache_Data_Traversable_Accum_Internal_applyStateR gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applyStateR sync.Once
func Get_Data_Traversable_Accum_Internal_applyStateR() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applyStateR.Do(func() {
		cache_Data_Traversable_Accum_Internal_applyStateR = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Traversable_Accum_Internal_functorStateR()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply(x_1, s_2)
_ = v_3_0
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v1_4_1, "value"), gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_applyStateR
}

var cache_Data_Traversable_Accum_Internal_applyStateL gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applyStateL sync.Once
func Get_Data_Traversable_Accum_Internal_applyStateL() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applyStateL.Do(func() {
		cache_Data_Traversable_Accum_Internal_applyStateL = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Traversable_Accum_Internal_functorStateL()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply(f_0, s_2)
_ = v_3_0
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply(x_1, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v_3_0, "value"), gopurs_runtime.RecordGet(v1_4_1, "value")))
})
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_applyStateL
}

var cache_Data_Traversable_Accum_Internal_applicativeStateR gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applicativeStateR sync.Once
func Get_Data_Traversable_Accum_Internal_applicativeStateR() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applicativeStateR.Do(func() {
		cache_Data_Traversable_Accum_Internal_applicativeStateR = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Traversable_Accum_Internal_applyStateR()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_applicativeStateR
}

var cache_Data_Traversable_Accum_Internal_applicativeStateL gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applicativeStateL sync.Once
func Get_Data_Traversable_Accum_Internal_applicativeStateL() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applicativeStateL.Do(func() {
		cache_Data_Traversable_Accum_Internal_applicativeStateL = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Traversable_Accum_Internal_applyStateL()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_applicativeStateL
}

var cache_Data_Traversable_Accum_Internal_applicativeStateL__2039640491 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applicativeStateL__2039640491 sync.Once
func Get_Data_Traversable_Accum_Internal_applicativeStateL__2039640491() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applicativeStateL__2039640491.Do(func() {
		cache_Data_Traversable_Accum_Internal_applicativeStateL__2039640491 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Traversable_Accum_Internal_applyStateL()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_applicativeStateL__2039640491
}

var cache_Data_Traversable_Accum_Internal_applicativeStateL__3057114219 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applicativeStateL__3057114219 sync.Once
func Get_Data_Traversable_Accum_Internal_applicativeStateL__3057114219() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applicativeStateL__3057114219.Do(func() {
		cache_Data_Traversable_Accum_Internal_applicativeStateL__3057114219 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Traversable_Accum_Internal_applyStateL()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_applicativeStateL__3057114219
}

var cache_Data_Traversable_Accum_Internal_applicativeStateR__2039640491 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applicativeStateR__2039640491 sync.Once
func Get_Data_Traversable_Accum_Internal_applicativeStateR__2039640491() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applicativeStateR__2039640491.Do(func() {
		cache_Data_Traversable_Accum_Internal_applicativeStateR__2039640491 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Traversable_Accum_Internal_applyStateR()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_applicativeStateR__2039640491
}

var cache_Data_Traversable_Accum_Internal_applicativeStateR__3057114219 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applicativeStateR__3057114219 sync.Once
func Get_Data_Traversable_Accum_Internal_applicativeStateR__3057114219() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applicativeStateR__3057114219.Do(func() {
		cache_Data_Traversable_Accum_Internal_applicativeStateR__3057114219 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_Traversable_Accum_Internal_applyStateR()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_applicativeStateR__3057114219
}

var cache_Data_Traversable_Accum_Internal_applyStateL__2377133483 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applyStateL__2377133483 sync.Once
func Get_Data_Traversable_Accum_Internal_applyStateL__2377133483() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applyStateL__2377133483.Do(func() {
		cache_Data_Traversable_Accum_Internal_applyStateL__2377133483 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Traversable_Accum_Internal_functorStateL()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply(f_0, s_2)
_ = v_3_0
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply(x_1, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v_3_0, "value"), gopurs_runtime.RecordGet(v1_4_1, "value")))
})
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_applyStateL__2377133483
}

var cache_Data_Traversable_Accum_Internal_applyStateL__1203051627 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applyStateL__1203051627 sync.Once
func Get_Data_Traversable_Accum_Internal_applyStateL__1203051627() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applyStateL__1203051627.Do(func() {
		cache_Data_Traversable_Accum_Internal_applyStateL__1203051627 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Traversable_Accum_Internal_functorStateL()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply(f_0, s_2)
_ = v_3_0
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply(x_1, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v_3_0, "value"), gopurs_runtime.RecordGet(v1_4_1, "value")))
})
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_applyStateL__1203051627
}

var cache_Data_Traversable_Accum_Internal_applyStateR__2377133483 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applyStateR__2377133483 sync.Once
func Get_Data_Traversable_Accum_Internal_applyStateR__2377133483() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applyStateR__2377133483.Do(func() {
		cache_Data_Traversable_Accum_Internal_applyStateR__2377133483 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Traversable_Accum_Internal_functorStateR()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply(x_1, s_2)
_ = v_3_0
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v1_4_1, "value"), gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_applyStateR__2377133483
}

var cache_Data_Traversable_Accum_Internal_applyStateR__1203051627 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_applyStateR__1203051627 sync.Once
func Get_Data_Traversable_Accum_Internal_applyStateR__1203051627() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_applyStateR__1203051627.Do(func() {
		cache_Data_Traversable_Accum_Internal_applyStateR__1203051627 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Traversable_Accum_Internal_functorStateR()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply(x_1, s_2)
_ = v_3_0
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v1_4_1, "value"), gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_applyStateR__1203051627
}

var cache_Data_Traversable_Accum_Internal_functorStateL__2836346332 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_functorStateL__2836346332 sync.Once
func Get_Data_Traversable_Accum_Internal_functorStateL__2836346332() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_functorStateL__2836346332.Do(func() {
		cache_Data_Traversable_Accum_Internal_functorStateL__2836346332 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_functorStateL__2836346332
}

var cache_Data_Traversable_Accum_Internal_functorStateL__2692529692 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_functorStateL__2692529692 sync.Once
func Get_Data_Traversable_Accum_Internal_functorStateL__2692529692() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_functorStateL__2692529692.Do(func() {
		cache_Data_Traversable_Accum_Internal_functorStateL__2692529692 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_functorStateL__2692529692
}

var cache_Data_Traversable_Accum_Internal_functorStateR__2836346332 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_functorStateR__2836346332 sync.Once
func Get_Data_Traversable_Accum_Internal_functorStateR__2836346332() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_functorStateR__2836346332.Do(func() {
		cache_Data_Traversable_Accum_Internal_functorStateR__2836346332 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_functorStateR__2836346332
}

var cache_Data_Traversable_Accum_Internal_functorStateR__2692529692 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_functorStateR__2692529692 sync.Once
func Get_Data_Traversable_Accum_Internal_functorStateR__2692529692() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_functorStateR__2692529692.Do(func() {
		cache_Data_Traversable_Accum_Internal_functorStateR__2692529692 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
})})}
	})
	return cache_Data_Traversable_Accum_Internal_functorStateR__2692529692
}

var cache_Data_Traversable_Accum_Internal_stateL__1334064830 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_stateL__1334064830 sync.Once
func Get_Data_Traversable_Accum_Internal_stateL__1334064830() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_stateL__1334064830.Do(func() {
		cache_Data_Traversable_Accum_Internal_stateL__1334064830 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Accum_Internal_stateL__1334064830(v_0_box)
})
	})
	return cache_Data_Traversable_Accum_Internal_stateL__1334064830
}

var cache_Data_Traversable_Accum_Internal_stateL__1412771550 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_stateL__1412771550 sync.Once
func Get_Data_Traversable_Accum_Internal_stateL__1412771550() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_stateL__1412771550.Do(func() {
		cache_Data_Traversable_Accum_Internal_stateL__1412771550 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Accum_Internal_stateL__1412771550(v_0_box)
})
	})
	return cache_Data_Traversable_Accum_Internal_stateL__1412771550
}

var cache_Data_Traversable_Accum_Internal_stateL__123903742 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_stateL__123903742 sync.Once
func Get_Data_Traversable_Accum_Internal_stateL__123903742() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_stateL__123903742.Do(func() {
		cache_Data_Traversable_Accum_Internal_stateL__123903742 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Accum_Internal_stateL__123903742(v_0_box)
})
	})
	return cache_Data_Traversable_Accum_Internal_stateL__123903742
}

var cache_Data_Traversable_Accum_Internal_stateL__3020604382 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_stateL__3020604382 sync.Once
func Get_Data_Traversable_Accum_Internal_stateL__3020604382() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_stateL__3020604382.Do(func() {
		cache_Data_Traversable_Accum_Internal_stateL__3020604382 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Accum_Internal_stateL__3020604382(v_0_box)
})
	})
	return cache_Data_Traversable_Accum_Internal_stateL__3020604382
}

var cache_Data_Traversable_Accum_Internal_stateR__1334064830 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_stateR__1334064830 sync.Once
func Get_Data_Traversable_Accum_Internal_stateR__1334064830() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_stateR__1334064830.Do(func() {
		cache_Data_Traversable_Accum_Internal_stateR__1334064830 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Accum_Internal_stateR__1334064830(v_0_box)
})
	})
	return cache_Data_Traversable_Accum_Internal_stateR__1334064830
}

var cache_Data_Traversable_Accum_Internal_stateR__1412771550 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_stateR__1412771550 sync.Once
func Get_Data_Traversable_Accum_Internal_stateR__1412771550() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_stateR__1412771550.Do(func() {
		cache_Data_Traversable_Accum_Internal_stateR__1412771550 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Accum_Internal_stateR__1412771550(v_0_box)
})
	})
	return cache_Data_Traversable_Accum_Internal_stateR__1412771550
}

var cache_Data_Traversable_Accum_Internal_stateR__123903742 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_stateR__123903742 sync.Once
func Get_Data_Traversable_Accum_Internal_stateR__123903742() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_stateR__123903742.Do(func() {
		cache_Data_Traversable_Accum_Internal_stateR__123903742 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Accum_Internal_stateR__123903742(v_0_box)
})
	})
	return cache_Data_Traversable_Accum_Internal_stateR__123903742
}

var cache_Data_Traversable_Accum_Internal_stateR__3020604382 gopurs_runtime.Value
var once_Data_Traversable_Accum_Internal_stateR__3020604382 sync.Once
func Get_Data_Traversable_Accum_Internal_stateR__3020604382() gopurs_runtime.Value {
	once_Data_Traversable_Accum_Internal_stateR__3020604382.Do(func() {
		cache_Data_Traversable_Accum_Internal_stateR__3020604382 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Accum_Internal_stateR__3020604382(v_0_box)
})
	})
	return cache_Data_Traversable_Accum_Internal_stateR__3020604382
}

func Call_Data_Traversable_Accum_Internal_StateR(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Traversable_Accum_Internal_StateL(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Traversable_Accum_Internal_stateR(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Traversable_Accum_Internal_stateL(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Traversable_Accum_Internal_stateL__1334064830(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Traversable_Accum_Internal_stateL__1412771550(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Traversable_Accum_Internal_stateL__123903742(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Traversable_Accum_Internal_stateL__3020604382(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Traversable_Accum_Internal_stateR__1334064830(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Traversable_Accum_Internal_stateR__1412771550(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Traversable_Accum_Internal_stateR__123903742(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Traversable_Accum_Internal_stateR__3020604382(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}


