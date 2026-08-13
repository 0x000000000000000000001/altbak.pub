package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_RWS_pure gopurs_runtime.Value
var once_Control_Monad_RWS_pure sync.Once
func Get_Control_Monad_RWS_pure() gopurs_runtime.Value {
	once_Control_Monad_RWS_pure.Do(func() {
		cache_Control_Monad_RWS_pure = gopurs_runtime.RecordGet(Get_Data_Identity_applicativeIdentity(), "pure")
	})
	return cache_Control_Monad_RWS_pure
}

var cache_Control_Monad_RWS_unwrap gopurs_runtime.Value
var once_Control_Monad_RWS_unwrap sync.Once
func Get_Control_Monad_RWS_unwrap() gopurs_runtime.Value {
	once_Control_Monad_RWS_unwrap.Do(func() {
		cache_Control_Monad_RWS_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Control_Monad_RWS_unwrap
}

var cache_Control_Monad_RWS_unwrap1 gopurs_runtime.Value
var once_Control_Monad_RWS_unwrap1 sync.Once
func Get_Control_Monad_RWS_unwrap1() gopurs_runtime.Value {
	once_Control_Monad_RWS_unwrap1.Do(func() {
		cache_Control_Monad_RWS_unwrap1 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Control_Monad_RWS_unwrap1
}

var cache_Control_Monad_RWS_unwrap2 gopurs_runtime.Value
var once_Control_Monad_RWS_unwrap2 sync.Once
func Get_Control_Monad_RWS_unwrap2() gopurs_runtime.Value {
	once_Control_Monad_RWS_unwrap2.Do(func() {
		cache_Control_Monad_RWS_unwrap2 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Control_Monad_RWS_unwrap2
}

var cache_Control_Monad_RWS_withRWS gopurs_runtime.Value
var once_Control_Monad_RWS_withRWS sync.Once
func Get_Control_Monad_RWS_withRWS() gopurs_runtime.Value {
	once_Control_Monad_RWS_withRWS.Do(func() {
		cache_Control_Monad_RWS_withRWS = Get_Control_Monad_RWS_Trans_withRWST()
	})
	return cache_Control_Monad_RWS_withRWS
}

var cache_Control_Monad_RWS_rws gopurs_runtime.Value
var once_Control_Monad_RWS_rws sync.Once
func Get_Control_Monad_RWS_rws() gopurs_runtime.Value {
	once_Control_Monad_RWS_rws.Do(func() {
		cache_Control_Monad_RWS_rws = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Call_Control_Monad_RWS_rws(f_0_box, r_1_box, s_2_box))}
})
	})
	return cache_Control_Monad_RWS_rws
}

var cache_Control_Monad_RWS_runRWS gopurs_runtime.Value
var once_Control_Monad_RWS_runRWS sync.Once
func Get_Control_Monad_RWS_runRWS() gopurs_runtime.Value {
	once_Control_Monad_RWS_runRWS.Do(func() {
		cache_Control_Monad_RWS_runRWS = gopurs_runtime.Func3(func(m_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Call_Control_Monad_RWS_runRWS(m_0_box, r_1_box, s_2_box))}
})
	})
	return cache_Control_Monad_RWS_runRWS
}

var cache_Control_Monad_RWS_mapRWS gopurs_runtime.Value
var once_Control_Monad_RWS_mapRWS sync.Once
func Get_Control_Monad_RWS_mapRWS() gopurs_runtime.Value {
	once_Control_Monad_RWS_mapRWS.Do(func() {
		cache_Control_Monad_RWS_mapRWS = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(Call_Control_Monad_RWS_mapRWS(f_0_box, v_1_box, r_2_box, s_3_box))}
})
	})
	return cache_Control_Monad_RWS_mapRWS
}

var cache_Control_Monad_RWS_execRWS gopurs_runtime.Value
var once_Control_Monad_RWS_execRWS sync.Once
func Get_Control_Monad_RWS_execRWS() gopurs_runtime.Value {
	once_Control_Monad_RWS_execRWS.Do(func() {
		cache_Control_Monad_RWS_execRWS = gopurs_runtime.Func3(func(m_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_RWS_execRWS(m_0_box, r_1_box, s_2_box))}
})
	})
	return cache_Control_Monad_RWS_execRWS
}

var cache_Control_Monad_RWS_evalRWS gopurs_runtime.Value
var once_Control_Monad_RWS_evalRWS sync.Once
func Get_Control_Monad_RWS_evalRWS() gopurs_runtime.Value {
	once_Control_Monad_RWS_evalRWS.Do(func() {
		cache_Control_Monad_RWS_evalRWS = gopurs_runtime.Func3(func(m_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_RWS_evalRWS(m_0_box, r_1_box, s_2_box))}
})
	})
	return cache_Control_Monad_RWS_evalRWS
}

func Call_Control_Monad_RWS_rws(f_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *Constructor_Control_Monad_RWS_Trans_RWSResult {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_RWS_Trans_RWSResult](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Identity_applicativeIdentity(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_RWS_Trans_RWSResult](gopurs_runtime.Apply2(f_0, r_1, s_2)))}))
}

func Call_Control_Monad_RWS_runRWS(m_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *Constructor_Control_Monad_RWS_Trans_RWSResult {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_RWS_Trans_RWSResult](gopurs_runtime.Apply2(m_0, r_1, s_2))
}

func Call_Control_Monad_RWS_mapRWS(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) *Constructor_Control_Monad_RWS_Trans_RWSResult {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_RWS_Trans_RWSResult](gopurs_runtime.Apply(f_0, gopurs_runtime.Apply2(v_1, r_2, s_3)))
}

func Call_Control_Monad_RWS_execRWS(m_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
// TAST (Let): Applicative0_3_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Identity_monadIdentity(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Identity_monadIdentity(), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(m_0, r_1, s_2), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_4.UnsafePtr).V0, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_4.UnsafePtr).V2})})
})))
}

func Call_Control_Monad_RWS_evalRWS(m_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
// TAST (Let): Applicative0_3_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Identity_monadIdentity(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Identity_monadIdentity(), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(m_0, r_1, s_2), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_4.UnsafePtr).V1, (*Constructor_Control_Monad_RWS_Trans_RWSResult)(v1_4.UnsafePtr).V2})})
})))
}


