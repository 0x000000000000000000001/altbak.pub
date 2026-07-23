package Data_Symbol

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var reifySymbol gopurs_runtime.Value
var once_reifySymbol sync.Once
func Get_reifySymbol() gopurs_runtime.Value {
	once_reifySymbol.Do(func() {
		reifySymbol = gopurs_runtime.Func2(func(s_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_unsafeCoerce(), gopurs_runtime.Func(func(dictIsSymbol_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, dictIsSymbol_2)
}), gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return s_0
})), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Proxy")))
})
	})
	return reifySymbol
}

var reflectSymbol gopurs_runtime.Value
var once_reflectSymbol sync.Once
func Get_reflectSymbol() gopurs_runtime.Value {
	once_reflectSymbol.Do(func() {
		reflectSymbol = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "reflectSymbol")
})
	})
	return reflectSymbol
}

func Get_unsafeCoerce() gopurs_runtime.Value {
	return _Gopurs_UnsafeCoerce
}
