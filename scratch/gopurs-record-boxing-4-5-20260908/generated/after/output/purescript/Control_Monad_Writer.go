package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Writer_unwrap gopurs_runtime.Value
var once_Control_Monad_Writer_unwrap sync.Once
func Get_Control_Monad_Writer_unwrap() gopurs_runtime.Value {
	once_Control_Monad_Writer_unwrap.Do(func() {
		cache_Control_Monad_Writer_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Control_Monad_Writer_unwrap
}

var cache_Control_Monad_Writer_writer gopurs_runtime.Value
var once_Control_Monad_Writer_writer sync.Once
func Get_Control_Monad_Writer_writer() gopurs_runtime.Value {
	once_Control_Monad_Writer_writer.Do(func() {
		cache_Control_Monad_Writer_writer = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Monad_Writer_writer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box))
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Control_Monad_Writer_writer
}

var cache_Control_Monad_Writer_runWriter gopurs_runtime.Value
var once_Control_Monad_Writer_runWriter sync.Once
func Get_Control_Monad_Writer_runWriter() gopurs_runtime.Value {
	once_Control_Monad_Writer_runWriter.Do(func() {
		cache_Control_Monad_Writer_runWriter = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Monad_Writer_runWriter(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box))
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Control_Monad_Writer_runWriter
}

var cache_Control_Monad_Writer_mapWriter gopurs_runtime.Value
var once_Control_Monad_Writer_mapWriter sync.Once
func Get_Control_Monad_Writer_mapWriter() gopurs_runtime.Value {
	once_Control_Monad_Writer_mapWriter.Do(func() {
		cache_Control_Monad_Writer_mapWriter = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Monad_Writer_mapWriter(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Control_Monad_Writer_mapWriter
}

var cache_Control_Monad_Writer_execWriter gopurs_runtime.Value
var once_Control_Monad_Writer_execWriter sync.Once
func Get_Control_Monad_Writer_execWriter() gopurs_runtime.Value {
	once_Control_Monad_Writer_execWriter.Do(func() {
		cache_Control_Monad_Writer_execWriter = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_execWriter(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](m_0_box))
})
	})
	return cache_Control_Monad_Writer_execWriter
}

func Call_Control_Monad_Writer_writer(x_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var x_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = x_0_loop
_ = x_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(x_0)}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Control_Monad_Writer_runWriter(x_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var x_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = x_0_loop
_ = x_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(x_0)}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Control_Monad_Writer_mapWriter(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)})
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Control_Monad_Writer_execWriter(m_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var m_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = m_0_loop
_ = m_0
return (m_0).V1
}


