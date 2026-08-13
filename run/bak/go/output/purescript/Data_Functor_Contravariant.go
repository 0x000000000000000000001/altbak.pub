package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Contravariant_Contravariant_dollarDict gopurs_runtime.Value
var once_Data_Functor_Contravariant_Contravariant_dollarDict sync.Once
func Get_Data_Functor_Contravariant_Contravariant_dollarDict() gopurs_runtime.Value {
	once_Data_Functor_Contravariant_Contravariant_dollarDict.Do(func() {
		cache_Data_Functor_Contravariant_Contravariant_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Contravariant_Contravariant_dollarDict(x_0_box)
})
	})
	return cache_Data_Functor_Contravariant_Contravariant_dollarDict
}

var cache_Data_Functor_Contravariant_contravariantConst gopurs_runtime.Value
var once_Data_Functor_Contravariant_contravariantConst sync.Once
func Get_Data_Functor_Contravariant_contravariantConst() gopurs_runtime.Value {
	once_Data_Functor_Contravariant_contravariantConst.Do(func() {
		cache_Data_Functor_Contravariant_contravariantConst = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_Data_Functor_Contravariant_contravariantConst
}

var cache_Data_Functor_Contravariant_cmap gopurs_runtime.Value
var once_Data_Functor_Contravariant_cmap sync.Once
func Get_Data_Functor_Contravariant_cmap() gopurs_runtime.Value {
	once_Data_Functor_Contravariant_cmap.Do(func() {
		cache_Data_Functor_Contravariant_cmap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Contravariant_cmap(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](dict_0_box))
})
	})
	return cache_Data_Functor_Contravariant_cmap
}

var cache_Data_Functor_Contravariant_cmapFlipped gopurs_runtime.Value
var once_Data_Functor_Contravariant_cmapFlipped sync.Once
func Get_Data_Functor_Contravariant_cmapFlipped() gopurs_runtime.Value {
	once_Data_Functor_Contravariant_cmapFlipped.Do(func() {
		cache_Data_Functor_Contravariant_cmapFlipped = gopurs_runtime.Func3(func(dictContravariant_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Contravariant_cmapFlipped(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](dictContravariant_0_box), x_1_box, f_2_box)
})
	})
	return cache_Data_Functor_Contravariant_cmapFlipped
}

var cache_Data_Functor_Contravariant_coerce gopurs_runtime.Value
var once_Data_Functor_Contravariant_coerce sync.Once
func Get_Data_Functor_Contravariant_coerce() gopurs_runtime.Value {
	once_Data_Functor_Contravariant_coerce.Do(func() {
		cache_Data_Functor_Contravariant_coerce = gopurs_runtime.Func3(func(dictContravariant_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Contravariant_coerce(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](dictContravariant_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_1_box), a_2_box)
})
	})
	return cache_Data_Functor_Contravariant_coerce
}

var cache_Data_Functor_Contravariant_imapC gopurs_runtime.Value
var once_Data_Functor_Contravariant_imapC sync.Once
func Get_Data_Functor_Contravariant_imapC() gopurs_runtime.Value {
	once_Data_Functor_Contravariant_imapC.Do(func() {
		cache_Data_Functor_Contravariant_imapC = gopurs_runtime.Func3(func(dictContravariant_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Contravariant_imapC(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](dictContravariant_0_box), v_1_box, f_2_box)
})
	})
	return cache_Data_Functor_Contravariant_imapC
}

var cache_Data_Functor_Contravariant_cmap__326373820 gopurs_runtime.Value
var once_Data_Functor_Contravariant_cmap__326373820 sync.Once
func Get_Data_Functor_Contravariant_cmap__326373820() gopurs_runtime.Value {
	once_Data_Functor_Contravariant_cmap__326373820.Do(func() {
		cache_Data_Functor_Contravariant_cmap__326373820 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Contravariant_cmap__326373820(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](dict_0_box))
})
	})
	return cache_Data_Functor_Contravariant_cmap__326373820
}

var cache_Data_Functor_Contravariant_cmap__1884541340 gopurs_runtime.Value
var once_Data_Functor_Contravariant_cmap__1884541340 sync.Once
func Get_Data_Functor_Contravariant_cmap__1884541340() gopurs_runtime.Value {
	once_Data_Functor_Contravariant_cmap__1884541340.Do(func() {
		cache_Data_Functor_Contravariant_cmap__1884541340 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Contravariant_cmap__1884541340(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Contravariant_Contravariant](dict_0_box))
})
	})
	return cache_Data_Functor_Contravariant_cmap__1884541340
}

type Constructor_Data_Functor_Contravariant_Contravariant struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[85171506] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Functor_Contravariant_Contravariant)(ptr)
		_ = c
		switch key {
		case "cmap": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Functor_Contravariant_Contravariant: " + key)
		}
	}
}


func Call_Data_Functor_Contravariant_Contravariant_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Functor_Contravariant_cmap(dict_0_loop *Constructor_Data_Functor_Contravariant_Contravariant) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Contravariant_Contravariant = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_Contravariant_cmapFlipped(dictContravariant_0_loop *Constructor_Data_Functor_Contravariant_Contravariant, x_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 *Constructor_Data_Functor_Contravariant_Contravariant = dictContravariant_0_loop
_ = dictContravariant_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictContravariant_0.V0), f_2, x_1)
}

func Call_Data_Functor_Contravariant_coerce(dictContravariant_0_loop *Constructor_Data_Functor_Contravariant_Contravariant, dictFunctor_1_loop *Constructor_Data_Functor_Functor, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 *Constructor_Data_Functor_Contravariant_Contravariant = dictContravariant_0_loop
_ = dictContravariant_0
var dictFunctor_1 *Constructor_Data_Functor_Functor = dictFunctor_1_loop
_ = dictFunctor_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_1.V0), Get_Data_Void_absurd(), gopurs_runtime.Apply2(gopurs_runtime.Box(dictContravariant_0.V0), Get_Data_Void_absurd(), a_2))
}

func Call_Data_Functor_Contravariant_imapC(dictContravariant_0_loop *Constructor_Data_Functor_Contravariant_Contravariant, v_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 *Constructor_Data_Functor_Contravariant_Contravariant = dictContravariant_0_loop
_ = dictContravariant_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictContravariant_0.V0), f_2)
}

func Call_Data_Functor_Contravariant_cmap__326373820(dict_0_loop *Constructor_Data_Functor_Contravariant_Contravariant) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Contravariant_Contravariant = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_Contravariant_cmap__1884541340(dict_0_loop *Constructor_Data_Functor_Contravariant_Contravariant) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Contravariant_Contravariant = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}


