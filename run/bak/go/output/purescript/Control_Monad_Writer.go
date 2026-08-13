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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Writer_writer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](x_0_box)))}
})
	})
	return cache_Control_Monad_Writer_writer
}

var cache_Control_Monad_Writer_runWriter gopurs_runtime.Value
var once_Control_Monad_Writer_runWriter sync.Once
func Get_Control_Monad_Writer_runWriter() gopurs_runtime.Value {
	once_Control_Monad_Writer_runWriter.Do(func() {
		cache_Control_Monad_Writer_runWriter = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Writer_runWriter(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](x_0_box)))}
})
	})
	return cache_Control_Monad_Writer_runWriter
}

var cache_Control_Monad_Writer_mapWriter gopurs_runtime.Value
var once_Control_Monad_Writer_mapWriter sync.Once
func Get_Control_Monad_Writer_mapWriter() gopurs_runtime.Value {
	once_Control_Monad_Writer_mapWriter.Do(func() {
		cache_Control_Monad_Writer_mapWriter = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Writer_mapWriter(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_1_box)))}
})
	})
	return cache_Control_Monad_Writer_mapWriter
}

var cache_Control_Monad_Writer_execWriter gopurs_runtime.Value
var once_Control_Monad_Writer_execWriter sync.Once
func Get_Control_Monad_Writer_execWriter() gopurs_runtime.Value {
	once_Control_Monad_Writer_execWriter.Do(func() {
		cache_Control_Monad_Writer_execWriter = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_execWriter(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](m_0_box))
})
	})
	return cache_Control_Monad_Writer_execWriter
}

var cache_Control_Monad_Writer_runWriter__3864461800 gopurs_runtime.Value
var once_Control_Monad_Writer_runWriter__3864461800 sync.Once
func Get_Control_Monad_Writer_runWriter__3864461800() gopurs_runtime.Value {
	once_Control_Monad_Writer_runWriter__3864461800.Do(func() {
		cache_Control_Monad_Writer_runWriter__3864461800 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Writer_runWriter__3864461800(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](x_0_box)))}
})
	})
	return cache_Control_Monad_Writer_runWriter__3864461800
}

func Call_Control_Monad_Writer_writer(x_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var x_0 *Constructor_Data_Tuple_Tuple = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Identity_applicativeIdentity(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(x_0)}))
}

func Call_Control_Monad_Writer_runWriter(x_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var x_0 *Constructor_Data_Tuple_Tuple = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Writer_mapWriter(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}))
}

func Call_Control_Monad_Writer_execWriter(m_0_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var m_0 *Constructor_Data_Tuple_Tuple = m_0_loop
_ = m_0
return (m_0).V1
}

func Call_Control_Monad_Writer_runWriter__3864461800(x_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var x_0 *Constructor_Data_Tuple_Tuple = x_0_loop
_ = x_0
return x_0
}


