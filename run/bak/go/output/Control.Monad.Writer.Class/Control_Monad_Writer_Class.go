package Control_Monad_Writer_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	unsafe "unsafe"
)

var cache_tell gopurs_runtime.Value
var once_tell sync.Once
func Get_tell() gopurs_runtime.Value {
	once_tell.Do(func() {
		cache_tell = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tell(dict_0_box)
})
	})
	return cache_tell
}

var cache_pass gopurs_runtime.Value
var once_pass sync.Once
func Get_pass() gopurs_runtime.Value {
	once_pass.Do(func() {
		cache_pass = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pass(dict_0_box)
})
	})
	return cache_pass
}

var cache_pass__gopurs_runtime_Value_1541745649 gopurs_runtime.Value
var once_pass__gopurs_runtime_Value_1541745649 sync.Once
func Get_pass__gopurs_runtime_Value_1541745649() gopurs_runtime.Value {
	once_pass__gopurs_runtime_Value_1541745649.Do(func() {
		cache_pass__gopurs_runtime_Value_1541745649 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pass__gopurs_runtime_Value_1541745649(dict_0_box)
})
	})
	return cache_pass__gopurs_runtime_Value_1541745649
}

var cache_listen gopurs_runtime.Value
var once_listen sync.Once
func Get_listen() gopurs_runtime.Value {
	once_listen.Do(func() {
		cache_listen = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listen(dict_0_box)
})
	})
	return cache_listen
}

var cache_listen__gopurs_runtime_Value_1541745649 gopurs_runtime.Value
var once_listen__gopurs_runtime_Value_1541745649 sync.Once
func Get_listen__gopurs_runtime_Value_1541745649() gopurs_runtime.Value {
	once_listen__gopurs_runtime_Value_1541745649.Do(func() {
		cache_listen__gopurs_runtime_Value_1541745649 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listen__gopurs_runtime_Value_1541745649(dict_0_box)
})
	})
	return cache_listen__gopurs_runtime_Value_1541745649
}

var cache_listens gopurs_runtime.Value
var once_listens sync.Once
func Get_listens() gopurs_runtime.Value {
	once_listens.Do(func() {
		cache_listens = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listens(dictMonadWriter_0_box)
})
	})
	return cache_listens
}

var cache_censor gopurs_runtime.Value
var once_censor sync.Once
func Get_censor() gopurs_runtime.Value {
	once_censor.Do(func() {
		cache_censor = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_censor(dictMonadWriter_0_box)
})
	})
	return cache_censor
}

func Call_tell(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "tell")
}

func Call_pass(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "pass")
}

func Call_pass__gopurs_runtime_Value_1541745649(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "pass")
}

func Call_listen(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "listen")
}

func Call_listen__gopurs_runtime_Value_1541745649(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "listen")
}

func Call_listens(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{}), "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), m_3), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})})
}))
})
}

func Call_censor(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{}), "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), m_3, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_4, f_2})})
})))
})
}


