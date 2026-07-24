package Control_Monad_Writer

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
return x_0
})
	})
	return runWriter
}

var mapWriter gopurs_runtime.Value
var once_mapWriter sync.Once
func Get_mapWriter() gopurs_runtime.Value {
	once_mapWriter.Do(func() {
		mapWriter = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v_1)
})
	})
	return mapWriter
}

var execWriter gopurs_runtime.Value
var once_execWriter sync.Once
func Get_execWriter() gopurs_runtime.Value {
	once_execWriter.Do(func() {
		execWriter = gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(m_0.UnsafePtr)[1]
})
	})
	return execWriter
}


