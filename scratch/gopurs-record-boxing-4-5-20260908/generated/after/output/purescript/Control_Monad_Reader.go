package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Control_Monad_Reader_unwrap gopurs_runtime.Value
var once_Control_Monad_Reader_unwrap sync.Once
func Get_Control_Monad_Reader_unwrap() gopurs_runtime.Value {
	once_Control_Monad_Reader_unwrap.Do(func() {
		cache_Control_Monad_Reader_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Control_Monad_Reader_unwrap
}

var cache_Control_Monad_Reader_withReader gopurs_runtime.Value
var once_Control_Monad_Reader_withReader sync.Once
func Get_Control_Monad_Reader_withReader() gopurs_runtime.Value {
	once_Control_Monad_Reader_withReader.Do(func() {
		cache_Control_Monad_Reader_withReader = Get_Control_Monad_Reader_Trans_withReaderT()
	})
	return cache_Control_Monad_Reader_withReader
}

var cache_Control_Monad_Reader_runReader gopurs_runtime.Value
var once_Control_Monad_Reader_runReader sync.Once
func Get_Control_Monad_Reader_runReader() gopurs_runtime.Value {
	once_Control_Monad_Reader_runReader.Do(func() {
		cache_Control_Monad_Reader_runReader = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_runReader(v_0_box, x_1_box)
})
	})
	return cache_Control_Monad_Reader_runReader
}

var cache_Control_Monad_Reader_mapReader gopurs_runtime.Value
var once_Control_Monad_Reader_mapReader sync.Once
func Get_Control_Monad_Reader_mapReader() gopurs_runtime.Value {
	once_Control_Monad_Reader_mapReader.Do(func() {
		cache_Control_Monad_Reader_mapReader = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_mapReader(f_0_box, v_1_box, x_2_box)
})
	})
	return cache_Control_Monad_Reader_mapReader
}

func Call_Control_Monad_Reader_runReader(v_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(v_0, x_1)
}

func Call_Control_Monad_Reader_mapReader(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
}


