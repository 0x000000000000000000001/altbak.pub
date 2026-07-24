package Control_Extend

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var extendFn gopurs_runtime.Value
var once_extendFn sync.Once
func Get_extendFn() gopurs_runtime.Value {
	once_extendFn.Do(func() {
		extendFn = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Func(func(w_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), w_3, w_prime_4))
}))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}))
}()
})
	})
	return extendFn
}

var extendArray gopurs_runtime.Value
var once_extendArray sync.Once
func Get_extendArray() gopurs_runtime.Value {
	once_extendArray.Do(func() {
		extendArray = gopurs_runtime.RecordDict2("extend", "Functor0", Get_arrayExtend(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}))
	})
	return extendArray
}

var extend gopurs_runtime.Value
var once_extend sync.Once
func Get_extend() gopurs_runtime.Value {
	once_extend.Do(func() {
		extend = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "extend")
}()
})
	})
	return extend
}

var extendFlipped gopurs_runtime.Value
var once_extendFlipped sync.Once
func Get_extendFlipped() gopurs_runtime.Value {
	once_extendFlipped.Do(func() {
		extendFlipped = gopurs_runtime.Func3(Call_extendFlipped)
	})
	return extendFlipped
}

var duplicate gopurs_runtime.Value
var once_duplicate sync.Once
func Get_duplicate() gopurs_runtime.Value {
	once_duplicate.Do(func() {
		duplicate = gopurs_runtime.Func(func(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0_loop, "extend"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}()
})
	})
	return duplicate
}

var composeCoKleisliFlipped gopurs_runtime.Value
var once_composeCoKleisliFlipped sync.Once
func Get_composeCoKleisliFlipped() gopurs_runtime.Value {
	once_composeCoKleisliFlipped.Do(func() {
		composeCoKleisliFlipped = gopurs_runtime.Func4(Call_composeCoKleisliFlipped)
	})
	return composeCoKleisliFlipped
}

var composeCoKleisli gopurs_runtime.Value
var once_composeCoKleisli sync.Once
func Get_composeCoKleisli() gopurs_runtime.Value {
	once_composeCoKleisli.Do(func() {
		composeCoKleisli = gopurs_runtime.Func4(Call_composeCoKleisli)
	})
	return composeCoKleisli
}

func Call_extendFlipped(dictExtend_0_loop gopurs_runtime.Value, w_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
var w_1 gopurs_runtime.Value = w_1_loop
_ = w_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0_loop, "extend"), f_2_loop, w_1_loop)
}

func Call_composeCoKleisliFlipped(dictExtend_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, w_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var w_3 gopurs_runtime.Value = w_3_loop
_ = w_3
return gopurs_runtime.Apply(f_1_loop, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0_loop, "extend"), g_2_loop, w_3_loop))
}

func Call_composeCoKleisli(dictExtend_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, w_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var w_3 gopurs_runtime.Value = w_3_loop
_ = w_3
return gopurs_runtime.Apply(g_2_loop, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0_loop, "extend"), f_1_loop, w_3_loop))
}

func Get_arrayExtend() gopurs_runtime.Value {
	return _Gopurs_ArrayExtend
}
