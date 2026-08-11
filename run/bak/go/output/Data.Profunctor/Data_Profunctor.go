package Data_Profunctor

import (
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_identity1 gopurs_runtime.Value
var once_identity1 sync.Once
func Get_identity1() gopurs_runtime.Value {
	once_identity1.Do(func() {
		cache_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity1(x_0_box)
})
	})
	return cache_identity1
}

var cache_profunctorFn gopurs_runtime.Value
var once_profunctorFn sync.Once
func Get_profunctorFn() gopurs_runtime.Value {
	once_profunctorFn.Do(func() {
		cache_profunctorFn = gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(a2b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c2d_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b2c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c2d_1, gopurs_runtime.Apply(b2c_2, gopurs_runtime.Apply(a2b_0, x_3)))
})
})
})
}))
	})
	return cache_profunctorFn
}

var cache_profunctorFn__gopurs_runtime_Value_3736629211 gopurs_runtime.Value
var once_profunctorFn__gopurs_runtime_Value_3736629211 sync.Once
func Get_profunctorFn__gopurs_runtime_Value_3736629211() gopurs_runtime.Value {
	once_profunctorFn__gopurs_runtime_Value_3736629211.Do(func() {
		cache_profunctorFn__gopurs_runtime_Value_3736629211 = gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(a2b_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c2d_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b2c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c2d_1, gopurs_runtime.Apply(b2c_2, gopurs_runtime.Apply(a2b_0, x_3)))
})
})
})
}))
	})
	return cache_profunctorFn__gopurs_runtime_Value_3736629211
}

var cache_dimap gopurs_runtime.Value
var once_dimap sync.Once
func Get_dimap() gopurs_runtime.Value {
	once_dimap.Do(func() {
		cache_dimap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dimap(dict_0_box)
})
	})
	return cache_dimap
}

var cache_dimap__gopurs_runtime_Value_1466332548 gopurs_runtime.Value
var once_dimap__gopurs_runtime_Value_1466332548 sync.Once
func Get_dimap__gopurs_runtime_Value_1466332548() gopurs_runtime.Value {
	once_dimap__gopurs_runtime_Value_1466332548.Do(func() {
		cache_dimap__gopurs_runtime_Value_1466332548 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dimap__gopurs_runtime_Value_1466332548(dict_0_box)
})
	})
	return cache_dimap__gopurs_runtime_Value_1466332548
}

var cache_lcmap gopurs_runtime.Value
var once_lcmap sync.Once
func Get_lcmap() gopurs_runtime.Value {
	once_lcmap.Do(func() {
		cache_lcmap = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, a2b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lcmap(dictProfunctor_0_box, a2b_1_box)
})
	})
	return cache_lcmap
}

var cache_lcmap__gopurs_runtime_Value_1762133278 gopurs_runtime.Value
var once_lcmap__gopurs_runtime_Value_1762133278 sync.Once
func Get_lcmap__gopurs_runtime_Value_1762133278() gopurs_runtime.Value {
	once_lcmap__gopurs_runtime_Value_1762133278.Do(func() {
		cache_lcmap__gopurs_runtime_Value_1762133278 = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, a2b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lcmap__gopurs_runtime_Value_1762133278(dictProfunctor_0_box, a2b_1_box)
})
	})
	return cache_lcmap__gopurs_runtime_Value_1762133278
}

var cache_rmap gopurs_runtime.Value
var once_rmap sync.Once
func Get_rmap() gopurs_runtime.Value {
	once_rmap.Do(func() {
		cache_rmap = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, b2c_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rmap(dictProfunctor_0_box, b2c_1_box)
})
	})
	return cache_rmap
}

var cache_rmap__gopurs_runtime_Value_1762133278 gopurs_runtime.Value
var once_rmap__gopurs_runtime_Value_1762133278 sync.Once
func Get_rmap__gopurs_runtime_Value_1762133278() gopurs_runtime.Value {
	once_rmap__gopurs_runtime_Value_1762133278.Do(func() {
		cache_rmap__gopurs_runtime_Value_1762133278 = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, b2c_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rmap__gopurs_runtime_Value_1762133278(dictProfunctor_0_box, b2c_1_box)
})
	})
	return cache_rmap__gopurs_runtime_Value_1762133278
}

var cache_unwrapIso gopurs_runtime.Value
var once_unwrapIso sync.Once
func Get_unwrapIso() gopurs_runtime.Value {
	once_unwrapIso.Do(func() {
		cache_unwrapIso = gopurs_runtime.Func2(func(dictProfunctor_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrapIso(dictProfunctor_0_box, _dollar__unused_1_box)
})
	})
	return cache_unwrapIso
}

var cache_wrapIso gopurs_runtime.Value
var once_wrapIso sync.Once
func Get_wrapIso() gopurs_runtime.Value {
	once_wrapIso.Do(func() {
		cache_wrapIso = gopurs_runtime.Func3(func(dictProfunctor_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_wrapIso(dictProfunctor_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return cache_wrapIso
}

var cache_arr gopurs_runtime.Value
var once_arr sync.Once
func Get_arr() gopurs_runtime.Value {
	once_arr.Do(func() {
		cache_arr = gopurs_runtime.Func(func(dictCategory_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_arr(dictCategory_0_box)
})
	})
	return cache_arr
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_dimap(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "dimap")
}

func Call_dimap__gopurs_runtime_Value_1466332548(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "dimap")
}

func Call_lcmap(dictProfunctor_0_loop gopurs_runtime.Value, a2b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
var a2b_1 gopurs_runtime.Value = a2b_1_loop
_ = a2b_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), a2b_1, Get_identity())
}

func Call_lcmap__gopurs_runtime_Value_1762133278(dictProfunctor_0_loop gopurs_runtime.Value, a2b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
var a2b_1 gopurs_runtime.Value = a2b_1_loop
_ = a2b_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), a2b_1, Get_identity())
}

func Call_rmap(dictProfunctor_0_loop gopurs_runtime.Value, b2c_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
var b2c_1 gopurs_runtime.Value = b2c_1_loop
_ = b2c_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), Get_identity1(), b2c_1)
}

func Call_rmap__gopurs_runtime_Value_1762133278(dictProfunctor_0_loop gopurs_runtime.Value, b2c_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
var b2c_1 gopurs_runtime.Value = b2c_1_loop
_ = b2c_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), Get_identity1(), b2c_1)
}

func Call_unwrapIso(dictProfunctor_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), pkg_Unsafe_Coerce.Get_unsafeCoerce(), pkg_Unsafe_Coerce.Get_unsafeCoerce())
}

func Call_wrapIso(dictProfunctor_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), pkg_Unsafe_Coerce.Get_unsafeCoerce(), pkg_Unsafe_Coerce.Get_unsafeCoerce())
}

func Call_arr(dictCategory_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCategory_0 gopurs_runtime.Value = dictCategory_0_loop
_ = dictCategory_0
identity2_1_0 := gopurs_runtime.RecordGet(dictCategory_0, "identity")
_ = identity2_1_0
return gopurs_runtime.Func(func(dictProfunctor_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_2, "dimap"), Get_identity1(), f_3, identity2_1_0)
})
})
}


