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
		cmap = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "cmap")
}()
})
	})
	return cmap
}

var cmapFlipped gopurs_runtime.Value
var once_cmapFlipped sync.Once
func Get_cmapFlipped() gopurs_runtime.Value {
	once_cmapFlipped.Do(func() {
		cmapFlipped = gopurs_runtime.Func3(Call_cmapFlipped)
	})
	return cmapFlipped
}

var coerce gopurs_runtime.Value
var once_coerce sync.Once
func Get_coerce() gopurs_runtime.Value {
	once_coerce.Do(func() {
		coerce = gopurs_runtime.Func3(Call_coerce)
	})
	return coerce
}

var imapC gopurs_runtime.Value
var once_imapC sync.Once
func Get_imapC() gopurs_runtime.Value {
	once_imapC.Do(func() {
		imapC = gopurs_runtime.Func3(Call_imapC)
	})
	return imapC
}

func Call_cmapFlipped(dictContravariant_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
_ = dictContravariant_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictContravariant_0_loop, "cmap"), f_2_loop, x_1_loop)
}

func Call_coerce(dictContravariant_0_loop gopurs_runtime.Value, dictFunctor_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
_ = dictContravariant_0
var dictFunctor_1 gopurs_runtime.Value = dictFunctor_1_loop
_ = dictFunctor_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_1_loop, "map"), pkg_Data_Void.Get_absurd(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictContravariant_0_loop, "cmap"), pkg_Data_Void.Get_absurd(), a_2_loop))
}

func Call_imapC(dictContravariant_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
_ = dictContravariant_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictContravariant_0_loop, "cmap"), f_2_loop)
}


