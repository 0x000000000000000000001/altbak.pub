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
		cache_tell = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "tell")
}()
})
	})
	return cache_tell
}

var cache_pass gopurs_runtime.Value
var once_pass sync.Once
func Get_pass() gopurs_runtime.Value {
	once_pass.Do(func() {
		cache_pass = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "pass")
}()
})
	})
	return cache_pass
}

var cache_listen gopurs_runtime.Value
var once_listen sync.Once
func Get_listen() gopurs_runtime.Value {
	once_listen.Do(func() {
		cache_listen = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "listen")
}()
})
	})
	return cache_listen
}

var cache_listens gopurs_runtime.Value
var once_listens sync.Once
func Get_listens() gopurs_runtime.Value {
	once_listens.Do(func() {
		cache_listens = gopurs_runtime.Func(func(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{}), "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), m_3), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{(*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)})})
}))
})
}()
})
	})
	return cache_listens
}

var cache_censor gopurs_runtime.Value
var once_censor sync.Once
func Get_censor() gopurs_runtime.Value {
	once_censor.Do(func() {
		cache_censor = gopurs_runtime.Func(func(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{}), "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), m_3, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{a_4, f_2})})
})))
})
}()
})
	})
	return cache_censor
}




