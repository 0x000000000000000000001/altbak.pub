package Control_Monad_Reader

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_Reader_Trans "gopurs/output/Control.Monad.Reader.Trans"
)

var withReader gopurs_runtime.Value
var once_withReader sync.Once
func Get_withReader() gopurs_runtime.Value {
	once_withReader.Do(func() {
		withReader = pkg_Control_Monad_Reader_Trans.Get_withReaderT()
	})
	return withReader
}

var runReader gopurs_runtime.Value
var once_runReader sync.Once
func Get_runReader() gopurs_runtime.Value {
	once_runReader.Do(func() {
		runReader = gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, x_1)
})
	})
	return runReader
}

var mapReader gopurs_runtime.Value
var once_mapReader sync.Once
func Get_mapReader() gopurs_runtime.Value {
	once_mapReader.Do(func() {
		mapReader = gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
})
	})
	return mapReader
}




