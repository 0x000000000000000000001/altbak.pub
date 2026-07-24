package Data_Profunctor_Split

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var SplitF gopurs_runtime.Value
var once_SplitF sync.Once
func Get_SplitF() gopurs_runtime.Value {
	once_SplitF.Do(func() {
		SplitF = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor3("SplitF", value0, value1, value2)
})
})
})
	})
	return SplitF
}

var unSplit gopurs_runtime.Value
var once_unSplit sync.Once
func Get_unSplit() gopurs_runtime.Value {
	once_unSplit.Do(func() {
		unSplit = gopurs_runtime.Func2(Call_unSplit)
	})
	return unSplit
}

var split gopurs_runtime.Value
var once_split sync.Once
func Get_split() gopurs_runtime.Value {
	once_split.Do(func() {
		split = gopurs_runtime.Func3(Call_split)
	})
	return split
}

var profunctorSplit gopurs_runtime.Value
var once_profunctorSplit sync.Once
func Get_profunctorSplit() gopurs_runtime.Value {
	once_profunctorSplit.Do(func() {
		profunctorSplit = gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor3("SplitF", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], gopurs_runtime.Apply(f_0, x_3))
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1], x_3))
}), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[2])
}))
	})
	return profunctorSplit
}

var lowerSplit gopurs_runtime.Value
var once_lowerSplit sync.Once
func Get_lowerSplit() gopurs_runtime.Value {
	once_lowerSplit.Do(func() {
		lowerSplit = gopurs_runtime.Func2(Call_lowerSplit)
	})
	return lowerSplit
}

var liftSplit gopurs_runtime.Value
var once_liftSplit sync.Once
func Get_liftSplit() gopurs_runtime.Value {
	once_liftSplit.Do(func() {
		liftSplit = gopurs_runtime.Func(func(fx_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var fx_0 gopurs_runtime.Value = fx_0_loop
_ = fx_0
return gopurs_runtime.Constructor3("SplitF", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), fx_0_loop)
}()
})
	})
	return liftSplit
}

var hoistSplit gopurs_runtime.Value
var once_hoistSplit sync.Once
func Get_hoistSplit() gopurs_runtime.Value {
	once_hoistSplit.Do(func() {
		hoistSplit = gopurs_runtime.Func2(Call_hoistSplit)
	})
	return hoistSplit
}

var functorSplit gopurs_runtime.Value
var once_functorSplit sync.Once
func Get_functorSplit() gopurs_runtime.Value {
	once_functorSplit.Do(func() {
		functorSplit = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor3("SplitF", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1], x_2))
}), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[2])
}))
	})
	return functorSplit
}

func Call_unSplit(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(f_0_loop, (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[2])
}

func Call_split(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, fx_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var fx_2 gopurs_runtime.Value = fx_2_loop
_ = fx_2
return gopurs_runtime.Constructor3("SplitF", f_0_loop, g_1_loop, fx_2_loop)
}

func Call_lowerSplit(dictInvariant_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictInvariant_0 gopurs_runtime.Value = dictInvariant_0_loop
_ = dictInvariant_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictInvariant_0_loop, "imap"), (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[2])
}

func Call_hoistSplit(nat_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var nat_0 gopurs_runtime.Value = nat_0_loop
_ = nat_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Constructor3("SplitF", (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[1], gopurs_runtime.Apply(nat_0_loop, (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[2]))
}


