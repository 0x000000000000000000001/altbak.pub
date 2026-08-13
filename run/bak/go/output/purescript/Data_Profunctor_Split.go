package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Split_identity gopurs_runtime.Value
var once_Data_Profunctor_Split_identity sync.Once
func Get_Data_Profunctor_Split_identity() gopurs_runtime.Value {
	once_Data_Profunctor_Split_identity.Do(func() {
		cache_Data_Profunctor_Split_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Split_identity(x_0_box)
})
	})
	return cache_Data_Profunctor_Split_identity
}

var cache_Data_Profunctor_Split_SplitF gopurs_runtime.Value
var once_Data_Profunctor_Split_SplitF sync.Once
func Get_Data_Profunctor_Split_SplitF() gopurs_runtime.Value {
	once_Data_Profunctor_Split_SplitF.Do(func() {
		cache_Data_Profunctor_Split_SplitF = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1, value2})}
})
})
})
	})
	return cache_Data_Profunctor_Split_SplitF
}

var cache_Data_Profunctor_Split_Split gopurs_runtime.Value
var once_Data_Profunctor_Split_Split sync.Once
func Get_Data_Profunctor_Split_Split() gopurs_runtime.Value {
	once_Data_Profunctor_Split_Split.Do(func() {
		cache_Data_Profunctor_Split_Split = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Split_Split(x_0_box)
})
	})
	return cache_Data_Profunctor_Split_Split
}

var cache_Data_Profunctor_Split_unSplit gopurs_runtime.Value
var once_Data_Profunctor_Split_unSplit sync.Once
func Get_Data_Profunctor_Split_unSplit() gopurs_runtime.Value {
	once_Data_Profunctor_Split_unSplit.Do(func() {
		cache_Data_Profunctor_Split_unSplit = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Split_unSplit(f_0_box, v_1_box)
})
	})
	return cache_Data_Profunctor_Split_unSplit
}

var cache_Data_Profunctor_Split_split gopurs_runtime.Value
var once_Data_Profunctor_Split_split sync.Once
func Get_Data_Profunctor_Split_split() gopurs_runtime.Value {
	once_Data_Profunctor_Split_split.Do(func() {
		cache_Data_Profunctor_Split_split = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, fx_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Split_split(f_0_box, g_1_box, fx_2_box)
})
	})
	return cache_Data_Profunctor_Split_split
}

var cache_Data_Profunctor_Split_profunctorSplit gopurs_runtime.Value
var once_Data_Profunctor_Split_profunctorSplit sync.Once
func Get_Data_Profunctor_Split_profunctorSplit() gopurs_runtime.Value {
	once_Data_Profunctor_Split_profunctorSplit.Do(func() {
		cache_Data_Profunctor_Split_profunctorSplit = gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply((*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply(f_0, x_3))
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, gopurs_runtime.Apply((*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, x_3))
}), (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2})}
})
})
}))
	})
	return cache_Data_Profunctor_Split_profunctorSplit
}

var cache_Data_Profunctor_Split_lowerSplit gopurs_runtime.Value
var once_Data_Profunctor_Split_lowerSplit sync.Once
func Get_Data_Profunctor_Split_lowerSplit() gopurs_runtime.Value {
	once_Data_Profunctor_Split_lowerSplit.Do(func() {
		cache_Data_Profunctor_Split_lowerSplit = gopurs_runtime.Func2(func(dictInvariant_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Split_lowerSplit(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]](dictInvariant_0_box), v_1_box)
})
	})
	return cache_Data_Profunctor_Split_lowerSplit
}

var cache_Data_Profunctor_Split_liftSplit gopurs_runtime.Value
var once_Data_Profunctor_Split_liftSplit sync.Once
func Get_Data_Profunctor_Split_liftSplit() gopurs_runtime.Value {
	once_Data_Profunctor_Split_liftSplit.Do(func() {
		cache_Data_Profunctor_Split_liftSplit = gopurs_runtime.Func(func(fx_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Split_liftSplit(fx_0_box)
})
	})
	return cache_Data_Profunctor_Split_liftSplit
}

var cache_Data_Profunctor_Split_hoistSplit gopurs_runtime.Value
var once_Data_Profunctor_Split_hoistSplit sync.Once
func Get_Data_Profunctor_Split_hoistSplit() gopurs_runtime.Value {
	once_Data_Profunctor_Split_hoistSplit.Do(func() {
		cache_Data_Profunctor_Split_hoistSplit = gopurs_runtime.Func2(func(nat_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Split_hoistSplit(nat_0_box, v_1_box)
})
	})
	return cache_Data_Profunctor_Split_hoistSplit
}

var cache_Data_Profunctor_Split_functorSplit gopurs_runtime.Value
var once_Data_Profunctor_Split_functorSplit sync.Once
func Get_Data_Profunctor_Split_functorSplit() gopurs_runtime.Value {
	once_Data_Profunctor_Split_functorSplit.Do(func() {
		cache_Data_Profunctor_Split_functorSplit = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply((*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, x_2))
}), (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2})}
})
}))
	})
	return cache_Data_Profunctor_Split_functorSplit
}

var cache_Data_Profunctor_Split_split__1533821423 gopurs_runtime.Value
var once_Data_Profunctor_Split_split__1533821423 sync.Once
func Get_Data_Profunctor_Split_split__1533821423() gopurs_runtime.Value {
	once_Data_Profunctor_Split_split__1533821423.Do(func() {
		cache_Data_Profunctor_Split_split__1533821423 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, fx_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Split_split__1533821423(f_0_box, g_1_box, fx_2_box)
})
	})
	return cache_Data_Profunctor_Split_split__1533821423
}

var cache_Data_Profunctor_Split_split__569735913 gopurs_runtime.Value
var once_Data_Profunctor_Split_split__569735913 sync.Once
func Get_Data_Profunctor_Split_split__569735913() gopurs_runtime.Value {
	once_Data_Profunctor_Split_split__569735913.Do(func() {
		cache_Data_Profunctor_Split_split__569735913 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, fx_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Split_split__569735913(f_0_box, g_1_box, fx_2_box)
})
	})
	return cache_Data_Profunctor_Split_split__569735913
}

var cache_Data_Profunctor_Split_unSplit__597066880 gopurs_runtime.Value
var once_Data_Profunctor_Split_unSplit__597066880 sync.Once
func Get_Data_Profunctor_Split_unSplit__597066880() gopurs_runtime.Value {
	once_Data_Profunctor_Split_unSplit__597066880.Do(func() {
		cache_Data_Profunctor_Split_unSplit__597066880 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Split_unSplit__597066880(f_0_box, v_1_box)
})
	})
	return cache_Data_Profunctor_Split_unSplit__597066880
}

var cache_Data_Profunctor_Split_unSplit__3745310982 gopurs_runtime.Value
var once_Data_Profunctor_Split_unSplit__3745310982 sync.Once
func Get_Data_Profunctor_Split_unSplit__3745310982() gopurs_runtime.Value {
	once_Data_Profunctor_Split_unSplit__3745310982.Do(func() {
		cache_Data_Profunctor_Split_unSplit__3745310982 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Split_unSplit__3745310982(f_0_box, v_1_box)
})
	})
	return cache_Data_Profunctor_Split_unSplit__3745310982
}

var cache_Data_Profunctor_Split_unSplit__3644369478 gopurs_runtime.Value
var once_Data_Profunctor_Split_unSplit__3644369478 sync.Once
func Get_Data_Profunctor_Split_unSplit__3644369478() gopurs_runtime.Value {
	once_Data_Profunctor_Split_unSplit__3644369478.Do(func() {
		cache_Data_Profunctor_Split_unSplit__3644369478 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Split_unSplit__3644369478(f_0_box, v_1_box)
})
	})
	return cache_Data_Profunctor_Split_unSplit__3644369478
}

type Constructor_Data_Profunctor_Split_SplitF[T_f any, T_a any, T_b any, T_x any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func Call_Data_Profunctor_Split_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Profunctor_Split_Split(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Profunctor_Split_unSplit(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(f_0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
}

func Call_Data_Profunctor_Split_split(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, fx_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var fx_2 gopurs_runtime.Value = fx_2_loop
_ = fx_2
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, f_0, g_1, fx_2})}
}

func Call_Data_Profunctor_Split_lowerSplit(dictInvariant_0_loop *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictInvariant_0 *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value] = dictInvariant_0_loop
_ = dictInvariant_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictInvariant_0.V0), (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
}

func Call_Data_Profunctor_Split_liftSplit(fx_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fx_0 gopurs_runtime.Value = fx_0_loop
_ = fx_0
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Profunctor_Split_identity(), Get_Data_Profunctor_Split_identity(), fx_0})}
}

func Call_Data_Profunctor_Split_hoistSplit(nat_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var nat_0 gopurs_runtime.Value = nat_0_loop
_ = nat_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, gopurs_runtime.Apply(nat_0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)})}
}

func Call_Data_Profunctor_Split_split__1533821423(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, fx_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var fx_2 gopurs_runtime.Value = fx_2_loop
_ = fx_2
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, f_0, g_1, fx_2})}
}

func Call_Data_Profunctor_Split_split__569735913(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, fx_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var fx_2 gopurs_runtime.Value = fx_2_loop
_ = fx_2
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, f_0, g_1, fx_2})}
}

func Call_Data_Profunctor_Split_unSplit__597066880(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(f_0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
}

func Call_Data_Profunctor_Split_unSplit__3745310982(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(f_0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
}

func Call_Data_Profunctor_Split_unSplit__3644369478(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(f_0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1, (*Constructor_Data_Profunctor_Split_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)
}


