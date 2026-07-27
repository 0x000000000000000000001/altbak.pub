package Data_Functor_Contravariant

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Void "gopurs/output/Data.Void"
)

var cache_contravariantConst gopurs_runtime.Value
var once_contravariantConst sync.Once
func Get_contravariantConst() gopurs_runtime.Value {
	once_contravariantConst.Do(func() {
		cache_contravariantConst = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}))))
	})
	return cache_contravariantConst
}

var cache_cmap gopurs_runtime.Value
var once_cmap sync.Once
func Get_cmap() gopurs_runtime.Value {
	once_cmap.Do(func() {
		cache_cmap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cmap(dict_0_box)
})
	})
	return cache_cmap
}

var cache_cmap__func_gopurs_runtime_Value__func_interface____interface____interface____interface___3810731752 gopurs_runtime.Value
var once_cmap__func_gopurs_runtime_Value__func_interface____interface____interface____interface___3810731752 sync.Once
func Get_cmap__func_gopurs_runtime_Value__func_interface____interface____interface____interface___3810731752() gopurs_runtime.Value {
	once_cmap__func_gopurs_runtime_Value__func_interface____interface____interface____interface___3810731752.Do(func() {
		cache_cmap__func_gopurs_runtime_Value__func_interface____interface____interface____interface___3810731752 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cmap__func_gopurs_runtime_Value__func_interface____interface____interface____interface___3810731752(dict_0_box)
})
	})
	return cache_cmap__func_gopurs_runtime_Value__func_interface____interface____interface____interface___3810731752
}

var cache_cmapFlipped gopurs_runtime.Value
var once_cmapFlipped sync.Once
func Get_cmapFlipped() gopurs_runtime.Value {
	once_cmapFlipped.Do(func() {
		cache_cmapFlipped = gopurs_runtime.Func3(func(dictContravariant_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_cmapFlipped(dictContravariant_0_box, gopurs_runtime.UnboxAny(x_1_box), func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_2_box, gopurs_runtime.Any(inner_arg0)))
}))
})
	})
	return cache_cmapFlipped
}

var cache_coerce gopurs_runtime.Value
var once_coerce sync.Once
func Get_coerce() gopurs_runtime.Value {
	once_coerce.Do(func() {
		cache_coerce = gopurs_runtime.Func3(func(dictContravariant_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_coerce(dictContravariant_0_box, dictFunctor_1_box, gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_coerce
}

var cache_imapC gopurs_runtime.Value
var once_imapC sync.Once
func Get_imapC() gopurs_runtime.Value {
	once_imapC.Do(func() {
		cache_imapC = gopurs_runtime.Func3(func(dictContravariant_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_imapC(dictContravariant_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_1_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_2_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_imapC
}

func Call_cmap(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "cmap")
}

func Call_cmap__func_gopurs_runtime_Value__func_interface____interface____interface____interface___3810731752(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "cmap")
}

func Call_cmapFlipped(dictContravariant_0_loop gopurs_runtime.Value, x_1_loop interface{}, f_2_loop func(interface{}) interface{}) interface{} {
var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
_ = dictContravariant_0
var x_1 interface{} = x_1_loop
_ = x_1
var f_2 func(interface{}) interface{} = f_2_loop
_ = f_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_2(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(x_1)))
}

func Call_coerce(dictContravariant_0_loop gopurs_runtime.Value, dictFunctor_1_loop gopurs_runtime.Value, a_2_loop interface{}) interface{} {
var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
_ = dictContravariant_0
var dictFunctor_1 gopurs_runtime.Value = dictFunctor_1_loop
_ = dictFunctor_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_1, "map"), pkg_Data_Void.Get_absurd(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), pkg_Data_Void.Get_absurd(), gopurs_runtime.Any(a_2))))
}

func Call_imapC(dictContravariant_0_loop gopurs_runtime.Value, v_1_loop func(interface{}) interface{}, f_2_loop func(interface{}) interface{}) gopurs_runtime.Value {
var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
_ = dictContravariant_0
var v_1 func(interface{}) interface{} = v_1_loop
_ = v_1
var f_2 func(interface{}) interface{} = f_2_loop
_ = f_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_2(gopurs_runtime.UnboxAny(arg0)))
}))
}
