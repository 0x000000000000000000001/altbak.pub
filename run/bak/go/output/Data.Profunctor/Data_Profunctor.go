package Data_Profunctor

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var profunctorFn gopurs_runtime.Value
var once_profunctorFn sync.Once
func Get_profunctorFn() gopurs_runtime.Value {
	once_profunctorFn.Do(func() {
		profunctorFn = gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func4(func(a2b_0 gopurs_runtime.Value, c2d_1 gopurs_runtime.Value, b2c_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c2d_1, gopurs_runtime.Apply(b2c_2, gopurs_runtime.Apply(a2b_0, x_3)))
}))
	})
	return profunctorFn
}

var dimap gopurs_runtime.Value
var once_dimap sync.Once
func Get_dimap() gopurs_runtime.Value {
	once_dimap.Do(func() {
		dimap = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "dimap")
})
	})
	return dimap
}

var lcmap gopurs_runtime.Value
var once_lcmap sync.Once
func Get_lcmap() gopurs_runtime.Value {
	once_lcmap.Do(func() {
		lcmap = gopurs_runtime.Func2(func(dictProfunctor_0 gopurs_runtime.Value, a2b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), a2b_1, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return lcmap
}

var rmap gopurs_runtime.Value
var once_rmap sync.Once
func Get_rmap() gopurs_runtime.Value {
	once_rmap.Do(func() {
		rmap = gopurs_runtime.Func2(func(dictProfunctor_0 gopurs_runtime.Value, b2c_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), b2c_1)
})
	})
	return rmap
}

var unwrapIso gopurs_runtime.Value
var once_unwrapIso sync.Once
func Get_unwrapIso() gopurs_runtime.Value {
	once_unwrapIso.Do(func() {
		unwrapIso = gopurs_runtime.Func2(func(dictProfunctor_0 gopurs_runtime.Value, _dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), pkg_Unsafe_Coerce.Get_unsafeCoerce(), pkg_Unsafe_Coerce.Get_unsafeCoerce())
})
	})
	return unwrapIso
}

var wrapIso gopurs_runtime.Value
var once_wrapIso sync.Once
func Get_wrapIso() gopurs_runtime.Value {
	once_wrapIso.Do(func() {
		wrapIso = gopurs_runtime.Func3(func(dictProfunctor_0 gopurs_runtime.Value, _dollar__unused_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), pkg_Unsafe_Coerce.Get_unsafeCoerce(), pkg_Unsafe_Coerce.Get_unsafeCoerce())
})
	})
	return wrapIso
}

var arr gopurs_runtime.Value
var once_arr sync.Once
func Get_arr() gopurs_runtime.Value {
	once_arr.Do(func() {
		arr = gopurs_runtime.Func(func(dictCategory_0 gopurs_runtime.Value) gopurs_runtime.Value {
identity1_1_0 := gopurs_runtime.RecordGet(dictCategory_0, "identity")
_ = identity1_1_0
return gopurs_runtime.Func2(func(dictProfunctor_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_2, "dimap"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), f_3, identity1_1_0)
})
})
	})
	return arr
}


