package Control_Monad_Reader

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_Reader_Trans "gopurs/output/Control.Monad.Reader.Trans"
)

var cache_withReader gopurs_runtime.Value
var once_withReader sync.Once
func Get_withReader() gopurs_runtime.Value {
	once_withReader.Do(func() {
		cache_withReader = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(interface{}) interface{}, inner_arg1 func(interface{}) interface{}, inner_arg2 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply3(pkg_Control_Monad_Reader_Trans.Get_withReaderT__func_func_interface____interface____func_interface____interface____interface____interface___3586686264(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg2)))
}(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg1, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(arg2)))
})
	})
	return cache_withReader
}

var cache_runReader gopurs_runtime.Value
var once_runReader sync.Once
func Get_runReader() gopurs_runtime.Value {
	once_runReader.Do(func() {
		cache_runReader = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_runReader(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_1_box)))
})
	})
	return cache_runReader
}

var cache_mapReader gopurs_runtime.Value
var once_mapReader sync.Once
func Get_mapReader() gopurs_runtime.Value {
	once_mapReader.Do(func() {
		cache_mapReader = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_mapReader(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_2_box)))
})
	})
	return cache_mapReader
}

func Call_runReader(v_0_loop func(interface{}) interface{}, x_1_loop interface{}) interface{} {
var v_0 func(interface{}) interface{} = v_0_loop
_ = v_0
var x_1 interface{} = x_1_loop
_ = x_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_0(x_1)))
}

func Call_mapReader(f_0_loop func(interface{}) interface{}, v_1_loop func(interface{}) interface{}, x_2_loop interface{}) interface{} {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
var v_1 func(interface{}) interface{} = v_1_loop
_ = v_1
var x_2 interface{} = x_2_loop
_ = x_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_1(x_2))))))
}
