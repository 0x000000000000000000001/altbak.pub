package Data_Functor_Contravariant

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Void "gopurs/output/Data.Void"
)

var contravariantConst gopurs_runtime.Value
var once_contravariantConst sync.Once
func Get_contravariantConst() gopurs_runtime.Value {
	once_contravariantConst.Do(func() {
		contravariantConst = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}))
	})
	return contravariantConst
}

var cmap gopurs_runtime.Value
var once_cmap sync.Once
func Get_cmap() gopurs_runtime.Value {
	once_cmap.Do(func() {
		cmap = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "cmap")
})
	})
	return cmap
}

var cmapFlipped gopurs_runtime.Value
var once_cmapFlipped sync.Once
func Get_cmapFlipped() gopurs_runtime.Value {
	once_cmapFlipped.Do(func() {
		cmapFlipped = gopurs_runtime.Func3(func(dictContravariant_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), f_2, x_1)
})
	})
	return cmapFlipped
}

var coerce gopurs_runtime.Value
var once_coerce sync.Once
func Get_coerce() gopurs_runtime.Value {
	once_coerce.Do(func() {
		coerce = gopurs_runtime.Func3(func(dictContravariant_0 gopurs_runtime.Value, dictFunctor_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_1, "map"), pkg_Data_Void.Get_absurd(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), pkg_Data_Void.Get_absurd(), a_2))
})
	})
	return coerce
}

var imapC gopurs_runtime.Value
var once_imapC sync.Once
func Get_imapC() gopurs_runtime.Value {
	once_imapC.Do(func() {
		imapC = gopurs_runtime.Func3(func(dictContravariant_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), f_2)
})
	})
	return imapC
}




