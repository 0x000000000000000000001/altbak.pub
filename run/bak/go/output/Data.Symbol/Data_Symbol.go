package Data_Symbol

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Type_Proxy "gopurs/output/Type.Proxy"
	unsafe "unsafe"
)

var reifySymbol gopurs_runtime.Value
var once_reifySymbol sync.Once
func Get_reifySymbol() gopurs_runtime.Value {
	once_reifySymbol.Do(func() {
		reifySymbol = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reifySymbol(s_0_box, f_1_box)
})
	})
	return reifySymbol
}

var reflectSymbol gopurs_runtime.Value
var once_reflectSymbol sync.Once
func Get_reflectSymbol() gopurs_runtime.Value {
	once_reflectSymbol.Do(func() {
		reflectSymbol = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "reflectSymbol")
}()
})
	})
	return reflectSymbol
}

func Call_reifySymbol(s_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply3(Get_unsafeCoerce(), gopurs_runtime.Func(func(dictIsSymbol_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, dictIsSymbol_2)
}), gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return s_0
})), gopurs_runtime.Value{Type: 9, IntVal: 3178699476, UnsafePtr: unsafe.Pointer(&pkg_Type_Proxy.Data_Type_Proxy_Proxy{})})
}

func Get_unsafeCoerce() gopurs_runtime.Value {
	return _Gopurs_UnsafeCoerce
}
