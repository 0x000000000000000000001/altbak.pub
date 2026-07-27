package Data_Exists

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var cache_runExists gopurs_runtime.Value
var once_runExists sync.Once
func Get_runExists() gopurs_runtime.Value {
	once_runExists.Do(func() {
		cache_runExists = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(interface{}) interface{}, inner_arg1 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), inner_arg1))
}(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)))
}, arg1))
})
	})
	return cache_runExists
}

var cache_runExists__func_func_interface____interface____gopurs_runtime_Value__interface___1222692984 gopurs_runtime.Value
var once_runExists__func_func_interface____interface____gopurs_runtime_Value__interface___1222692984 sync.Once
func Get_runExists__func_func_interface____interface____gopurs_runtime_Value__interface___1222692984() gopurs_runtime.Value {
	once_runExists__func_func_interface____interface____gopurs_runtime_Value__interface___1222692984.Do(func() {
		cache_runExists__func_func_interface____interface____gopurs_runtime_Value__interface___1222692984 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(interface{}) interface{}, inner_arg1 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), inner_arg1))
}(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)))
}, arg1))
})
	})
	return cache_runExists__func_func_interface____interface____gopurs_runtime_Value__interface___1222692984
}

var cache_mkExists gopurs_runtime.Value
var once_mkExists sync.Once
func Get_mkExists() gopurs_runtime.Value {
	once_mkExists.Do(func() {
		cache_mkExists = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return func(inner_arg0 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Any(inner_arg0))
}(gopurs_runtime.UnboxAny(arg0))
})
	})
	return cache_mkExists
}

var cache_mkExists__func_interface____gopurs_runtime_Value_2466284825 gopurs_runtime.Value
var once_mkExists__func_interface____gopurs_runtime_Value_2466284825 sync.Once
func Get_mkExists__func_interface____gopurs_runtime_Value_2466284825() gopurs_runtime.Value {
	once_mkExists__func_interface____gopurs_runtime_Value_2466284825.Do(func() {
		cache_mkExists__func_interface____gopurs_runtime_Value_2466284825 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return func(inner_arg0 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Any(inner_arg0))
}(gopurs_runtime.UnboxAny(arg0))
})
	})
	return cache_mkExists__func_interface____gopurs_runtime_Value_2466284825
}


