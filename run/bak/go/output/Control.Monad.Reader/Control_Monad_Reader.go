package Control_Monad_Reader

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_Reader_Trans "gopurs/output/Control.Monad.Reader.Trans"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Data_Identity "gopurs/output/Data.Identity"
)

var cache_withReader gopurs_runtime.Value
var once_withReader sync.Once
func Get_withReader() gopurs_runtime.Value {
	once_withReader.Do(func() {
		cache_withReader = pkg_Control_Monad_Reader_Trans.Get_withReaderT__gopurs_runtime_Value()
	})
	return cache_withReader
}

var cache_runReader gopurs_runtime.Value
var once_runReader sync.Once
func Get_runReader() gopurs_runtime.Value {
	once_runReader.Do(func() {
		cache_runReader = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runReader(v_0_box)
})
	})
	return cache_runReader
}

var cache_mapReader gopurs_runtime.Value
var once_mapReader sync.Once
func Get_mapReader() gopurs_runtime.Value {
	once_mapReader.Do(func() {
		cache_mapReader = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapReader(f_0_box)
})
	})
	return cache_mapReader
}

func Call_runReader(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pkg_Unsafe_Coerce.Get_unsafeCoerce(), v_0)
}

func Call_mapReader(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(pkg_Control_Monad_Reader_Trans.Get_mapReaderT(), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pkg_Data_Identity.Get_Identity(), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), f_0, pkg_Unsafe_Coerce.Get_unsafeCoerce())))
}


