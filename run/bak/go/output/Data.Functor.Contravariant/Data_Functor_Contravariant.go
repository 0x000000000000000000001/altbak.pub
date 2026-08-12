package Data_Functor_Contravariant

import (
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Void "gopurs/output/Data.Void"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_contravariantConst gopurs_runtime.Value
var once_contravariantConst sync.Once
func Get_contravariantConst() gopurs_runtime.Value {
	once_contravariantConst.Do(func() {
		cache_contravariantConst = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_contravariantConst
}

var cache_cmap gopurs_runtime.Value
var once_cmap sync.Once
func Get_cmap() gopurs_runtime.Value {
	once_cmap.Do(func() {
		cache_cmap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cmap(gopurs_runtime.CoerceToStruct[Constructor_Contravariant[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_cmap
}

var cache_cmap__gopurs_runtime_Value_326373820 gopurs_runtime.Value
var once_cmap__gopurs_runtime_Value_326373820 sync.Once
func Get_cmap__gopurs_runtime_Value_326373820() gopurs_runtime.Value {
	once_cmap__gopurs_runtime_Value_326373820.Do(func() {
		cache_cmap__gopurs_runtime_Value_326373820 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cmap__gopurs_runtime_Value_326373820(gopurs_runtime.CoerceToStruct[Constructor_Contravariant[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_cmap__gopurs_runtime_Value_326373820
}

var cache_cmapFlipped gopurs_runtime.Value
var once_cmapFlipped sync.Once
func Get_cmapFlipped() gopurs_runtime.Value {
	once_cmapFlipped.Do(func() {
		cache_cmapFlipped = gopurs_runtime.Func3(func(dictContravariant_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cmapFlipped(gopurs_runtime.CoerceToStruct[Constructor_Contravariant[gopurs_runtime.Value]](dictContravariant_0_box), x_1_box, f_2_box)
})
	})
	return cache_cmapFlipped
}

var cache_coerce gopurs_runtime.Value
var once_coerce sync.Once
func Get_coerce() gopurs_runtime.Value {
	once_coerce.Do(func() {
		cache_coerce = gopurs_runtime.Func3(func(dictContravariant_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coerce(gopurs_runtime.CoerceToStruct[Constructor_Contravariant[gopurs_runtime.Value]](dictContravariant_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_1_box), a_2_box)
})
	})
	return cache_coerce
}

var cache_imapC gopurs_runtime.Value
var once_imapC sync.Once
func Get_imapC() gopurs_runtime.Value {
	once_imapC.Do(func() {
		cache_imapC = gopurs_runtime.Func3(func(dictContravariant_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_imapC(gopurs_runtime.CoerceToStruct[Constructor_Contravariant[gopurs_runtime.Value]](dictContravariant_0_box), v_1_box, f_2_box)
})
	})
	return cache_imapC
}

type Constructor_Contravariant[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[85171506] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Contravariant[gopurs_runtime.Value])(ptr)
		switch key {
		case "cmap": return c.V0
		default: panic("Key not found in dictionary Constructor_Contravariant: " + key)
		}
	}
}


func Call_cmap(dict_0_loop *Constructor_Contravariant[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Contravariant[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_cmap__gopurs_runtime_Value_326373820(dict_0_loop *Constructor_Contravariant[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Contravariant[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_cmapFlipped(dictContravariant_0_loop *Constructor_Contravariant[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 *Constructor_Contravariant[gopurs_runtime.Value] = dictContravariant_0_loop
_ = dictContravariant_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictContravariant_0.V0, f_2, x_1)
}

func Call_coerce(dictContravariant_0_loop *Constructor_Contravariant[gopurs_runtime.Value], dictFunctor_1_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 *Constructor_Contravariant[gopurs_runtime.Value] = dictContravariant_0_loop
_ = dictContravariant_0
var dictFunctor_1 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(dictFunctor_1.V0, pkg_Data_Void.Get_absurd(), gopurs_runtime.Apply2(dictContravariant_0.V0, pkg_Data_Void.Get_absurd(), a_2))
}

func Call_imapC(dictContravariant_0_loop *Constructor_Contravariant[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 *Constructor_Contravariant[gopurs_runtime.Value] = dictContravariant_0_loop
_ = dictContravariant_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply(dictContravariant_0.V0, f_2)
}


