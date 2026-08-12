package Data_Profunctor_Split

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Functor_Invariant "gopurs/output/Data.Functor.Invariant"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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

var cache_SplitF gopurs_runtime.Value
var once_SplitF sync.Once
func Get_SplitF() gopurs_runtime.Value {
	once_SplitF.Do(func() {
		cache_SplitF = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, value2})}
})
})
})
	})
	return cache_SplitF
}

var cache_unSplit gopurs_runtime.Value
var once_unSplit sync.Once
func Get_unSplit() gopurs_runtime.Value {
	once_unSplit.Do(func() {
		cache_unSplit = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unSplit(f_0_box, v_1_box)
})
	})
	return cache_unSplit
}

var cache_split gopurs_runtime.Value
var once_split sync.Once
func Get_split() gopurs_runtime.Value {
	once_split.Do(func() {
		cache_split = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, fx_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_split(f_0_box, g_1_box, fx_2_box)
})
	})
	return cache_split
}

var cache_profunctorSplit gopurs_runtime.Value
var once_profunctorSplit sync.Once
func Get_profunctorSplit() gopurs_runtime.Value {
	once_profunctorSplit.Do(func() {
		cache_profunctorSplit = gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply(f_0, x_3))
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, gopurs_runtime.Apply((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, x_3))
}), (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2})}
})
})
}))
	})
	return cache_profunctorSplit
}

var cache_lowerSplit gopurs_runtime.Value
var once_lowerSplit sync.Once
func Get_lowerSplit() gopurs_runtime.Value {
	once_lowerSplit.Do(func() {
		cache_lowerSplit = gopurs_runtime.Func2(func(dictInvariant_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lowerSplit(gopurs_runtime.CoerceToStruct[pkg_Data_Functor_Invariant.Constructor_Invariant[gopurs_runtime.Value]](dictInvariant_0_box), v_1_box)
})
	})
	return cache_lowerSplit
}

var cache_liftSplit gopurs_runtime.Value
var once_liftSplit sync.Once
func Get_liftSplit() gopurs_runtime.Value {
	once_liftSplit.Do(func() {
		cache_liftSplit = gopurs_runtime.Func(func(fx_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftSplit(fx_0_box)
})
	})
	return cache_liftSplit
}

var cache_hoistSplit gopurs_runtime.Value
var once_hoistSplit sync.Once
func Get_hoistSplit() gopurs_runtime.Value {
	once_hoistSplit.Do(func() {
		cache_hoistSplit = gopurs_runtime.Func2(func(nat_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hoistSplit(nat_0_box, v_1_box)
})
	})
	return cache_hoistSplit
}

var cache_functorSplit gopurs_runtime.Value
var once_functorSplit sync.Once
func Get_functorSplit() gopurs_runtime.Value {
	once_functorSplit.Do(func() {
		cache_functorSplit = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, x_2))
}), (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2})}
})
}))
	})
	return cache_functorSplit
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_mkExists__689758297 gopurs_runtime.Value
var once_mkExists__689758297 sync.Once
func Get_mkExists__689758297() gopurs_runtime.Value {
	once_mkExists__689758297.Do(func() {
		cache_mkExists__689758297 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_mkExists__689758297
}

var cache_mkExists__1788455359 gopurs_runtime.Value
var once_mkExists__1788455359 sync.Once
func Get_mkExists__1788455359() gopurs_runtime.Value {
	once_mkExists__1788455359.Do(func() {
		cache_mkExists__1788455359 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_mkExists__1788455359
}

var cache_mkExists__3241495993 gopurs_runtime.Value
var once_mkExists__3241495993 sync.Once
func Get_mkExists__3241495993() gopurs_runtime.Value {
	once_mkExists__3241495993.Do(func() {
		cache_mkExists__3241495993 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_mkExists__3241495993
}

var cache_runExists__3717333575 gopurs_runtime.Value
var once_runExists__3717333575 sync.Once
func Get_runExists__3717333575() gopurs_runtime.Value {
	once_runExists__3717333575.Do(func() {
		cache_runExists__3717333575 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_runExists__3717333575
}

var cache_runExists__2888357825 gopurs_runtime.Value
var once_runExists__2888357825 sync.Once
func Get_runExists__2888357825() gopurs_runtime.Value {
	once_runExists__2888357825.Do(func() {
		cache_runExists__2888357825 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_runExists__2888357825
}

var cache_runExists__1921822855 gopurs_runtime.Value
var once_runExists__1921822855 sync.Once
func Get_runExists__1921822855() gopurs_runtime.Value {
	once_runExists__1921822855.Do(func() {
		cache_runExists__1921822855 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_runExists__1921822855
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__1093730912 gopurs_runtime.Value
var once_flip__1093730912 sync.Once
func Get_flip__1093730912() gopurs_runtime.Value {
	once_flip__1093730912.Do(func() {
		cache_flip__1093730912 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1093730912(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1093730912
}

var cache_split__1533821423 gopurs_runtime.Value
var once_split__1533821423 sync.Once
func Get_split__1533821423() gopurs_runtime.Value {
	once_split__1533821423.Do(func() {
		cache_split__1533821423 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, fx_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_split__1533821423(f_0_box, g_1_box, fx_2_box)
})
	})
	return cache_split__1533821423
}

var cache_split__569735913 gopurs_runtime.Value
var once_split__569735913 sync.Once
func Get_split__569735913() gopurs_runtime.Value {
	once_split__569735913.Do(func() {
		cache_split__569735913 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, fx_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_split__569735913(f_0_box, g_1_box, fx_2_box)
})
	})
	return cache_split__569735913
}

var cache_unSplit__597066880 gopurs_runtime.Value
var once_unSplit__597066880 sync.Once
func Get_unSplit__597066880() gopurs_runtime.Value {
	once_unSplit__597066880.Do(func() {
		cache_unSplit__597066880 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unSplit__597066880(f_0_box, v_1_box)
})
	})
	return cache_unSplit__597066880
}

var cache_unSplit__3745310982 gopurs_runtime.Value
var once_unSplit__3745310982 sync.Once
func Get_unSplit__3745310982() gopurs_runtime.Value {
	once_unSplit__3745310982.Do(func() {
		cache_unSplit__3745310982 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unSplit__3745310982(f_0_box, v_1_box)
})
	})
	return cache_unSplit__3745310982
}

var cache_unSplit__3644369478 gopurs_runtime.Value
var once_unSplit__3644369478 sync.Once
func Get_unSplit__3644369478() gopurs_runtime.Value {
	once_unSplit__3644369478.Do(func() {
		cache_unSplit__3644369478 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unSplit__3644369478(f_0_box, v_1_box)
})
	})
	return cache_unSplit__3644369478
}

type Constructor_SplitF[T_f any, T_a any, T_b any, T_x any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_unSplit(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(f_0, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
}

func Call_split(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, fx_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var fx_2 gopurs_runtime.Value = fx_2_loop
_ = fx_2
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, f_0, g_1, fx_2})}
}

func Call_lowerSplit(dictInvariant_0_loop *pkg_Data_Functor_Invariant.Constructor_Invariant[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictInvariant_0 *pkg_Data_Functor_Invariant.Constructor_Invariant[gopurs_runtime.Value] = dictInvariant_0_loop
_ = dictInvariant_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(dictInvariant_0.V0, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
}

func Call_liftSplit(fx_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fx_0 gopurs_runtime.Value = fx_0_loop
_ = fx_0
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_identity(), Get_identity(), fx_0})}
}

func Call_hoistSplit(nat_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var nat_0 gopurs_runtime.Value = nat_0_loop
_ = nat_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, gopurs_runtime.Apply(nat_0, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)})}
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1093730912(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_split__1533821423(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, fx_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var fx_2 gopurs_runtime.Value = fx_2_loop
_ = fx_2
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, f_0, g_1, fx_2})}
}

func Call_split__569735913(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, fx_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var fx_2 gopurs_runtime.Value = fx_2_loop
_ = fx_2
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, f_0, g_1, fx_2})}
}

func Call_unSplit__597066880(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(f_0, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
}

func Call_unSplit__3745310982(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(f_0, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
}

func Call_unSplit__3644369478(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(f_0, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
}


