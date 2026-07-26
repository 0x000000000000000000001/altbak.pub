package Control_Monad_Writer

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Control_Monad_Writer_Trans "gopurs/output/Control.Monad.Writer.Trans"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var cache_writer gopurs_runtime.Value
var once_writer sync.Once
func Get_writer() gopurs_runtime.Value {
	once_writer.Do(func() {
		cache_writer = gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pkg_Control_Monad_Writer_Trans.Get_WriterT(), gopurs_runtime.RecordGet(pkg_Data_Identity.Get_applicativeIdentity(), "pure"))
	})
	return cache_writer
}

var cache_runWriter gopurs_runtime.Value
var once_runWriter sync.Once
func Get_runWriter() gopurs_runtime.Value {
	once_runWriter.Do(func() {
		cache_runWriter = gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pkg_Unsafe_Coerce.Get_unsafeCoerce(), pkg_Control_Monad_Writer_Trans.Get_runWriterT())
	})
	return cache_runWriter
}

var cache_mapWriter gopurs_runtime.Value
var once_mapWriter sync.Once
func Get_mapWriter() gopurs_runtime.Value {
	once_mapWriter.Do(func() {
		cache_mapWriter = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWriter(f_0_box)
})
	})
	return cache_mapWriter
}

var cache_execWriter gopurs_runtime.Value
var once_execWriter sync.Once
func Get_execWriter() gopurs_runtime.Value {
	once_execWriter.Do(func() {
		cache_execWriter = gopurs_runtime.Func(func(m_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_execWriter(m_0_box)
})
	})
	return cache_execWriter
}

func Call_mapWriter(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pkg_Data_Identity.Get_Identity(), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), f_0, pkg_Unsafe_Coerce.Get_unsafeCoerce()))
}

func Call_execWriter(m_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(Get_runWriter(), m_0).UnsafePtr).V1
}


