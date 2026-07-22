package Control_Monad_Writer

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var writer gopurs_runtime.Value
var once_writer sync.Once
func Get_writer() gopurs_runtime.Value {
	once_writer.Do(func() {
		writer = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return writer
}

var runWriter gopurs_runtime.Value
var once_runWriter sync.Once
func Get_runWriter() gopurs_runtime.Value {
	once_runWriter.Do(func() {
		runWriter = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), x_0)
})
	})
	return runWriter
}

var mapWriter gopurs_runtime.Value
var once_mapWriter sync.Once
func Get_mapWriter() gopurs_runtime.Value {
	once_mapWriter.Do(func() {
		mapWriter = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), v_1))
})
})
	})
	return mapWriter
}

var execWriter gopurs_runtime.Value
var once_execWriter sync.Once
func Get_execWriter() gopurs_runtime.Value {
	once_execWriter.Do(func() {
		execWriter = gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), m_0).PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
	})
	return execWriter
}


