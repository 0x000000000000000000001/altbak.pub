package Control_Monad_RWS

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_RWS_Trans "gopurs/output/Control.Monad.RWS.Trans"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	unsafe "unsafe"
)

var cache_execRWST gopurs_runtime.Value
var once_execRWST sync.Once
func Get_execRWST() gopurs_runtime.Value {
	once_execRWST.Do(func() {
		cache_execRWST = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_execRWST(v_0_box, r_1_box, s_2_box)
})
	})
	return cache_execRWST
}

var cache_evalRWST gopurs_runtime.Value
var once_evalRWST sync.Once
func Get_evalRWST() gopurs_runtime.Value {
	once_evalRWST.Do(func() {
		cache_evalRWST = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_evalRWST(v_0_box, r_1_box, s_2_box)
})
	})
	return cache_evalRWST
}

var cache_withRWS gopurs_runtime.Value
var once_withRWS sync.Once
func Get_withRWS() gopurs_runtime.Value {
	once_withRWS.Do(func() {
		cache_withRWS = pkg_Control_Monad_RWS_Trans.Get_withRWST()
	})
	return cache_withRWS
}

var cache_rws gopurs_runtime.Value
var once_rws sync.Once
func Get_rws() gopurs_runtime.Value {
	once_rws.Do(func() {
		cache_rws = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rws(f_0_box, r_1_box, s_2_box)
})
	})
	return cache_rws
}

var cache_runRWS gopurs_runtime.Value
var once_runRWS sync.Once
func Get_runRWS() gopurs_runtime.Value {
	once_runRWS.Do(func() {
		cache_runRWS = gopurs_runtime.Func3(func(m_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runRWS(m_0_box, r_1_box, s_2_box)
})
	})
	return cache_runRWS
}

var cache_mapRWS gopurs_runtime.Value
var once_mapRWS sync.Once
func Get_mapRWS() gopurs_runtime.Value {
	once_mapRWS.Do(func() {
		cache_mapRWS = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapRWS(f_0_box, v_1_box, r_2_box, s_3_box)
})
	})
	return cache_mapRWS
}

var cache_execRWS gopurs_runtime.Value
var once_execRWS sync.Once
func Get_execRWS() gopurs_runtime.Value {
	once_execRWS.Do(func() {
		cache_execRWS = gopurs_runtime.Func3(func(m_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_execRWS(m_0_box, r_1_box, s_2_box)
})
	})
	return cache_execRWS
}

var cache_evalRWS gopurs_runtime.Value
var once_evalRWS sync.Once
func Get_evalRWS() gopurs_runtime.Value {
	once_evalRWS.Do(func() {
		cache_evalRWS = gopurs_runtime.Func3(func(m_0_box gopurs_runtime.Value, r_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_evalRWS(m_0_box, r_1_box, s_2_box)
})
	})
	return cache_evalRWS
}

func Call_execRWST(v_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_monadIdentity(), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(v_0, r_1, s_2), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_monadIdentity(), "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V2})})
}))
}

func Call_evalRWST(v_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_monadIdentity(), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(v_0, r_1, s_2), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_monadIdentity(), "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1, (*pkg_Control_Monad_RWS_Trans.Constructor_RWSResult[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V2})})
}))
}

func Call_rws(f_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_applicativeIdentity(), "pure"), gopurs_runtime.Apply2(f_0, r_1, s_2))
}

func Call_runRWS(m_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return gopurs_runtime.Apply2(m_0, r_1, s_2)
}

func Call_mapRWS(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply2(v_1, r_2, s_3))
}

func Call_execRWS(m_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return Call_execRWST(m_0, r_1, s_2)
}

func Call_evalRWS(m_0_loop gopurs_runtime.Value, r_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var r_1 gopurs_runtime.Value = r_1_loop
_ = r_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return Call_evalRWST(m_0, r_1, s_2)
}


