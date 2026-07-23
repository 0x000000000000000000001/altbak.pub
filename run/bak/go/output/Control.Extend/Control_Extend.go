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
		extendFn = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Func(func(w_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), w_3, w_prime_4))
}))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}))
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
		extend = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "extend")
})
	})
	return extend
}

var extendFlipped gopurs_runtime.Value
var once_extendFlipped sync.Once
func Get_extendFlipped() gopurs_runtime.Value {
	once_extendFlipped.Do(func() {
		extendFlipped = gopurs_runtime.Func3(func(dictExtend_0 gopurs_runtime.Value, w_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), f_2, w_1)
})
	})
	return extendFlipped
}

var duplicate gopurs_runtime.Value
var once_duplicate sync.Once
func Get_duplicate() gopurs_runtime.Value {
	once_duplicate.Do(func() {
		duplicate = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return duplicate
}

var composeCoKleisliFlipped gopurs_runtime.Value
var once_composeCoKleisliFlipped sync.Once
func Get_composeCoKleisliFlipped() gopurs_runtime.Value {
	once_composeCoKleisliFlipped.Do(func() {
		composeCoKleisliFlipped = gopurs_runtime.Func4(func(dictExtend_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), g_2, w_3))
})
	})
	return composeCoKleisliFlipped
}

var composeCoKleisli gopurs_runtime.Value
var once_composeCoKleisli sync.Once
func Get_composeCoKleisli() gopurs_runtime.Value {
	once_composeCoKleisli.Do(func() {
		composeCoKleisli = gopurs_runtime.Func4(func(dictExtend_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), f_1, w_3))
})
	})
	return composeCoKleisli
}

func Get_arrayExtend() gopurs_runtime.Value {
	return _Gopurs_ArrayExtend
}
