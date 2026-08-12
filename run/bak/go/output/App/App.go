package App

import (
	pkg_Bench "gopurs/output/Bench"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Array "gopurs/output/Data.Array"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Test_Ackermann "gopurs/output/Test.Ackermann"
	pkg_Test_ArrayOps "gopurs/output/Test.ArrayOps"
	pkg_Test_AstTree "gopurs/output/Test.AstTree"
	pkg_Test_Church "gopurs/output/Test.Church"
	pkg_Test_Fib "gopurs/output/Test.Fib"
	pkg_Test_LazyEvaluation "gopurs/output/Test.LazyEvaluation"
	pkg_Test_ListOps "gopurs/output/Test.ListOps"
	pkg_Test_Polymorphism "gopurs/output/Test.Polymorphism"
	pkg_Test_Primes "gopurs/output/Test.Primes"
	pkg_Test_RBTree "gopurs/output/Test.RBTree"
	pkg_Test_Records "gopurs/output/Test.Records"
	pkg_Test_StateMonad "gopurs/output/Test.StateMonad"
	pkg_Test_TCO "gopurs/output/Test.TCO"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_main gopurs_runtime.Value
var once_main sync.Once
func Get_main() gopurs_runtime.Value {
	once_main.Do(func() {
		cache_main = Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_AstTree.Get_describe(), pkg_Test_AstTree.Get_act()), gopurs_runtime.Func(func(t1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Fib.Get_describe(), pkg_Test_Fib.Get_act()), gopurs_runtime.Func(func(t2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_ListOps.Get_describe(), pkg_Test_ListOps.Get_act()), gopurs_runtime.Func(func(t3_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_TCO.Get_describe(), pkg_Test_TCO.Get_act()), gopurs_runtime.Func(func(t4_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Records.Get_describe(), pkg_Test_Records.Get_act()), gopurs_runtime.Func(func(t5_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Ackermann.Get_describe(), pkg_Test_Ackermann.Get_act()), gopurs_runtime.Func(func(t6_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Church.Get_describe(), pkg_Test_Church.Get_act()), gopurs_runtime.Func(func(t7_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Primes.Get_describe(), pkg_Test_Primes.Get_act()), gopurs_runtime.Func(func(t8_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_RBTree.Get_describe(), pkg_Test_RBTree.Get_act()), gopurs_runtime.Func(func(t9_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Polymorphism.Get_describe(), pkg_Test_Polymorphism.Get_act()), gopurs_runtime.Func(func(t10_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_StateMonad.Get_describe(), pkg_Test_StateMonad.Get_act()), gopurs_runtime.Func(func(t11_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_LazyEvaluation.Get_describe(), pkg_Test_LazyEvaluation.Get_act()), gopurs_runtime.Func(func(t12_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_ArrayOps.Get_describe(), pkg_Test_ArrayOps.Get_act()), gopurs_runtime.Func(func(t13_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("\x0a==================================================\x0a\x0aTotal exec time: "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Bench.Get_formatNumber(), gopurs_runtime.Float(Call_add__137136408(gopurs_runtime.Float(Call_add__137136408(gopurs_runtime.Float(Call_add__137136408(gopurs_runtime.Float(Call_add__137136408(gopurs_runtime.Float(Call_add__137136408(gopurs_runtime.Float(Call_add__137136408(gopurs_runtime.Float(Call_add__137136408(gopurs_runtime.Float(Call_add__137136408(gopurs_runtime.Float(Call_add__137136408(gopurs_runtime.Float(Call_add__137136408(gopurs_runtime.Float(Call_add__137136408(gopurs_runtime.Float(Call_add__137136408(gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t1_0.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal()), gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t2_1.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t3_2.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t4_3.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t5_4.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t6_5.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t7_6.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t8_7.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t9_8.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t10_9.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t11_10.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t12_11.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal()), gopurs_runtime.Float(Call_div__1002719800(gopurs_runtime.Float(t13_12.FloatVal()), gopurs_runtime.Float(1000.0)).FloatVal())).FloatVal())).StrVal()), gopurs_runtime.Str(" ms\x0a")).StrVal())).StrVal()))
}))
}))
}))
}))
}))
}))
}))
}))
}))
}))
}))
}))
}))
	})
	return cache_main
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__629383158 gopurs_runtime.Value
var once_pure__629383158 sync.Once
func Get_pure__629383158() gopurs_runtime.Value {
	once_pure__629383158.Do(func() {
		cache_pure__629383158 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__629383158(__eta0_0_box)
})
	})
	return cache_pure__629383158
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__3550378017 gopurs_runtime.Value
var once_bind__3550378017 sync.Once
func Get_bind__3550378017() gopurs_runtime.Value {
	once_bind__3550378017.Do(func() {
		cache_bind__3550378017 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3550378017(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_bind__3550378017
}

var cache_bind__1949526049 gopurs_runtime.Value
var once_bind__1949526049 sync.Once
func Get_bind__1949526049() gopurs_runtime.Value {
	once_bind__1949526049.Do(func() {
		cache_bind__1949526049 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1949526049(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_bind__1949526049
}

var cache_bind__3328406721 gopurs_runtime.Value
var once_bind__3328406721 sync.Once
func Get_bind__3328406721() gopurs_runtime.Value {
	once_bind__3328406721.Do(func() {
		cache_bind__3328406721 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3328406721(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_bind__3328406721
}

var cache_discard__203210016 gopurs_runtime.Value
var once_discard__203210016 sync.Once
func Get_discard__203210016() gopurs_runtime.Value {
	once_discard__203210016.Do(func() {
		cache_discard__203210016 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__203210016(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_discard__203210016
}

var cache_discard__317162198 gopurs_runtime.Value
var once_discard__317162198 sync.Once
func Get_discard__317162198() gopurs_runtime.Value {
	once_discard__317162198.Do(func() {
		cache_discard__317162198 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__317162198(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_discard__317162198
}

var cache_discardUnit__2687062302 gopurs_runtime.Value
var once_discardUnit__2687062302 sync.Once
func Get_discardUnit__2687062302() gopurs_runtime.Value {
	once_discardUnit__2687062302.Do(func() {
		cache_discardUnit__2687062302 = gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))
	})
	return cache_discardUnit__2687062302
}

var cache_filter__4047711382 gopurs_runtime.Value
var once_filter__4047711382 sync.Once
func Get_filter__4047711382() gopurs_runtime.Value {
	once_filter__4047711382.Do(func() {
		cache_filter__4047711382 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_filter__4047711382(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_filter__4047711382
}

var cache_filter__377906483 gopurs_runtime.Value
var once_filter__377906483 sync.Once
func Get_filter__377906483() gopurs_runtime.Value {
	once_filter__377906483.Do(func() {
		cache_filter__377906483 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_filter__377906483(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_filter__377906483
}

var cache_foldl__849397914 gopurs_runtime.Value
var once_foldl__849397914 sync.Once
func Get_foldl__849397914() gopurs_runtime.Value {
	once_foldl__849397914.Do(func() {
		cache_foldl__849397914 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl")
	})
	return cache_foldl__849397914
}

var cache_foldl__1469296346 gopurs_runtime.Value
var once_foldl__1469296346 sync.Once
func Get_foldl__1469296346() gopurs_runtime.Value {
	once_foldl__1469296346.Do(func() {
		cache_foldl__1469296346 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl")
	})
	return cache_foldl__1469296346
}

var cache_eq__789642299 gopurs_runtime.Value
var once_eq__789642299 sync.Once
func Get_eq__789642299() gopurs_runtime.Value {
	once_eq__789642299.Do(func() {
		cache_eq__789642299 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__789642299(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[bool]](dict_0_box))
})
	})
	return cache_eq__789642299
}

var cache_eq__1697837627 gopurs_runtime.Value
var once_eq__1697837627 sync.Once
func Get_eq__1697837627() gopurs_runtime.Value {
	once_eq__1697837627.Do(func() {
		cache_eq__1697837627 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__1697837627(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__1697837627
}

var cache_eq__2843686287 gopurs_runtime.Value
var once_eq__2843686287 sync.Once
func Get_eq__2843686287() gopurs_runtime.Value {
	once_eq__2843686287.Do(func() {
		cache_eq__2843686287 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2843686287(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__2843686287
}

var cache_eq__2276491096 gopurs_runtime.Value
var once_eq__2276491096 sync.Once
func Get_eq__2276491096() gopurs_runtime.Value {
	once_eq__2276491096.Do(func() {
		cache_eq__2276491096 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2276491096(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__2276491096
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_notEq__2843686287 gopurs_runtime.Value
var once_notEq__2843686287 sync.Once
func Get_notEq__2843686287() gopurs_runtime.Value {
	once_notEq__2843686287.Do(func() {
		cache_notEq__2843686287 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__2843686287(x_0_box, y_1_box))
})
	})
	return cache_notEq__2843686287
}

var cache_notEq__2384498378 gopurs_runtime.Value
var once_notEq__2384498378 sync.Once
func Get_notEq__2384498378() gopurs_runtime.Value {
	once_notEq__2384498378.Do(func() {
		cache_notEq__2384498378 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, y_2_box))
})
	})
	return cache_notEq__2384498378
}

var cache_div__1002719800 gopurs_runtime.Value
var once_div__1002719800 sync.Once
func Get_div__1002719800() gopurs_runtime.Value {
	once_div__1002719800.Do(func() {
		cache_div__1002719800 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div__1002719800(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_div__1002719800
}

var cache_div__2579358968 gopurs_runtime.Value
var once_div__2579358968 sync.Once
func Get_div__2579358968() gopurs_runtime.Value {
	once_div__2579358968.Do(func() {
		cache_div__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_div__2579358968
}

var cache_mod__2185172824 gopurs_runtime.Value
var once_mod__2185172824 sync.Once
func Get_mod__2185172824() gopurs_runtime.Value {
	once_mod__2185172824.Do(func() {
		cache_mod__2185172824 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__2185172824(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_mod__2185172824
}

var cache_mod__2579358968 gopurs_runtime.Value
var once_mod__2579358968 sync.Once
func Get_mod__2579358968() gopurs_runtime.Value {
	once_mod__2579358968.Do(func() {
		cache_mod__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mod__2579358968
}

var cache_compare__669572705 gopurs_runtime.Value
var once_compare__669572705 sync.Once
func Get_compare__669572705() gopurs_runtime.Value {
	once_compare__669572705.Do(func() {
		cache_compare__669572705 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__669572705(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__669572705
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_greaterThan__4087042607 gopurs_runtime.Value
var once_greaterThan__4087042607 sync.Once
func Get_greaterThan__4087042607() gopurs_runtime.Value {
	once_greaterThan__4087042607.Do(func() {
		cache_greaterThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__4087042607
}

var cache_greaterThan__1409282474 gopurs_runtime.Value
var once_greaterThan__1409282474 sync.Once
func Get_greaterThan__1409282474() gopurs_runtime.Value {
	once_greaterThan__1409282474.Do(func() {
		cache_greaterThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__1409282474
}

var cache_lessThan__4087042607 gopurs_runtime.Value
var once_lessThan__4087042607 sync.Once
func Get_lessThan__4087042607() gopurs_runtime.Value {
	once_lessThan__4087042607.Do(func() {
		cache_lessThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThan__4087042607
}

var cache_lessThan__1409282474 gopurs_runtime.Value
var once_lessThan__1409282474 sync.Once
func Get_lessThan__1409282474() gopurs_runtime.Value {
	once_lessThan__1409282474.Do(func() {
		cache_lessThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1409282474
}

var cache_sub__2927892844 gopurs_runtime.Value
var once_sub__2927892844 sync.Once
func Get_sub__2927892844() gopurs_runtime.Value {
	once_sub__2927892844.Do(func() {
		cache_sub__2927892844 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__2927892844(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[int64]](dict_0_box))
})
	})
	return cache_sub__2927892844
}

var cache_sub__1124926121 gopurs_runtime.Value
var once_sub__1124926121 sync.Once
func Get_sub__1124926121() gopurs_runtime.Value {
	once_sub__1124926121.Do(func() {
		cache_sub__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1124926121(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__1124926121
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1043827704
}

var cache_sub__1135378904 gopurs_runtime.Value
var once_sub__1135378904 sync.Once
func Get_sub__1135378904() gopurs_runtime.Value {
	once_sub__1135378904.Do(func() {
		cache_sub__1135378904 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1135378904(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1135378904
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_add__560788792
}

var cache_add__137136408 gopurs_runtime.Value
var once_add__137136408 sync.Once
func Get_add__137136408() gopurs_runtime.Value {
	once_add__137136408.Do(func() {
		cache_add__137136408 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__137136408(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_add__137136408
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_mul__560788792 gopurs_runtime.Value
var once_mul__560788792 sync.Once
func Get_mul__560788792() gopurs_runtime.Value {
	once_mul__560788792.Do(func() {
		cache_mul__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_mul__560788792
}

var cache_mul__1614463960 gopurs_runtime.Value
var once_mul__1614463960 sync.Once
func Get_mul__1614463960() gopurs_runtime.Value {
	once_mul__1614463960.Do(func() {
		cache_mul__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul__1614463960
}

var cache_show__3754018243 gopurs_runtime.Value
var once_show__3754018243 sync.Once
func Get_show__3754018243() gopurs_runtime.Value {
	once_show__3754018243.Do(func() {
		cache_show__3754018243 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__3754018243(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__3754018243
}

var cache_show__3984012258 gopurs_runtime.Value
var once_show__3984012258 sync.Once
func Get_show__3984012258() gopurs_runtime.Value {
	once_show__3984012258.Do(func() {
		cache_show__3984012258 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__3984012258(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_show__3984012258
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_logShow__2885109999 gopurs_runtime.Value
var once_logShow__2885109999 sync.Once
func Get_logShow__2885109999() gopurs_runtime.Value {
	once_logShow__2885109999.Do(func() {
		cache_logShow__2885109999 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_logShow__2885109999(a_0_box)
})
	})
	return cache_logShow__2885109999
}

var cache_logShow__339054415 gopurs_runtime.Value
var once_logShow__339054415 sync.Once
func Get_logShow__339054415() gopurs_runtime.Value {
	once_logShow__339054415.Do(func() {
		cache_logShow__339054415 = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_logShow__339054415(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dictShow_0_box), a_1_box)
})
	})
	return cache_logShow__339054415
}

var cache_applicativeEffect__284161122 gopurs_runtime.Value
var once_applicativeEffect__284161122 sync.Once
func Get_applicativeEffect__284161122() gopurs_runtime.Value {
	once_applicativeEffect__284161122.Do(func() {
		cache_applicativeEffect__284161122 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_pureE())
	})
	return cache_applicativeEffect__284161122
}

var cache_applyEffect__2014400020 gopurs_runtime.Value
var once_applyEffect__2014400020 sync.Once
func Get_applyEffect__2014400020() gopurs_runtime.Value {
	once_applyEffect__2014400020.Do(func() {
		cache_applyEffect__2014400020 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect__2014400020
}

var cache_bindEffect__2113658466 gopurs_runtime.Value
var once_bindEffect__2113658466 sync.Once
func Get_bindEffect__2113658466() gopurs_runtime.Value {
	once_bindEffect__2113658466.Do(func() {
		cache_bindEffect__2113658466 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_bindE())
	})
	return cache_bindEffect__2113658466
}

var cache_functorEffect__3107547953 gopurs_runtime.Value
var once_functorEffect__3107547953 sync.Once
func Get_functorEffect__3107547953() gopurs_runtime.Value {
	once_functorEffect__3107547953.Do(func() {
		cache_functorEffect__3107547953 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__3107547953
}

var cache_mulC__1746928225 gopurs_runtime.Value
var once_mulC__1746928225 sync.Once
func Get_mulC__1746928225() gopurs_runtime.Value {
	once_mulC__1746928225.Do(func() {
		cache_mulC__1746928225 = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mulC__1746928225(m_0_box, n_1_box, f_2_box, x_3_box)
})
	})
	return cache_mulC__1746928225
}

var cache_mulC__3596604257 gopurs_runtime.Value
var once_mulC__3596604257 sync.Once
func Get_mulC__3596604257() gopurs_runtime.Value {
	once_mulC__3596604257.Do(func() {
		cache_mulC__3596604257 = gopurs_runtime.Func4(func(m_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mulC__3596604257(m_0_box, n_1_box, f_2_box, x_3_box)
})
	})
	return cache_mulC__3596604257
}

var cache_succC__952275393 gopurs_runtime.Value
var once_succC__952275393 sync.Once
func Get_succC__952275393() gopurs_runtime.Value {
	once_succC__952275393.Do(func() {
		cache_succC__952275393 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succC__952275393(n_0_box, f_1_box, x_2_box)
})
	})
	return cache_succC__952275393
}

var cache_succC__1461826241 gopurs_runtime.Value
var once_succC__1461826241 sync.Once
func Get_succC__1461826241() gopurs_runtime.Value {
	once_succC__1461826241.Do(func() {
		cache_succC__1461826241 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succC__1461826241(n_0_box, f_1_box, x_2_box)
})
	})
	return cache_succC__1461826241
}

var cache_zeroC__4066693242 gopurs_runtime.Value
var once_zeroC__4066693242 sync.Once
func Get_zeroC__4066693242() gopurs_runtime.Value {
	once_zeroC__4066693242.Do(func() {
		cache_zeroC__4066693242 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_zeroC__4066693242(v_0_box, x_1_box.IntVal))
})
	})
	return cache_zeroC__4066693242
}

var cache_defer__3520065601 gopurs_runtime.Value
var once_defer__3520065601 sync.Once
func Get_defer__3520065601() gopurs_runtime.Value {
	once_defer__3520065601.Do(func() {
		cache_defer__3520065601 = pkg_Test_LazyEvaluation.Get_Lazy()
	})
	return cache_defer__3520065601
}

var cache_defer__3363737377 gopurs_runtime.Value
var once_defer__3363737377 sync.Once
func Get_defer__3363737377() gopurs_runtime.Value {
	once_defer__3363737377.Do(func() {
		cache_defer__3363737377 = pkg_Test_LazyEvaluation.Get_Lazy()
	})
	return cache_defer__3363737377
}

var cache_force__3902501304 gopurs_runtime.Value
var once_force__3902501304 sync.Once
func Get_force__3902501304() gopurs_runtime.Value {
	once_force__3902501304.Do(func() {
		cache_force__3902501304 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_force__3902501304(v_0_box)
})
	})
	return cache_force__3902501304
}

var cache_force__721037880 gopurs_runtime.Value
var once_force__721037880 sync.Once
func Get_force__721037880() gopurs_runtime.Value {
	once_force__721037880.Do(func() {
		cache_force__721037880 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_force__721037880(v_0_box)
})
	})
	return cache_force__721037880
}

var cache_foldl__3262866295 gopurs_runtime.Value
var once_foldl__3262866295 sync.Once
func Get_foldl__3262866295() gopurs_runtime.Value {
	once_foldl__3262866295.Do(func() {
		cache_foldl__3262866295 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3262866295(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Test_ListOps.Constructor_Cons[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_foldl__3262866295
}

var cache_foldl__1255354935 gopurs_runtime.Value
var once_foldl__1255354935 sync.Once
func Get_foldl__1255354935() gopurs_runtime.Value {
	once_foldl__1255354935.Do(func() {
		cache_foldl__1255354935 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1255354935(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Test_ListOps.Constructor_Cons[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_foldl__1255354935
}

var cache_mappend___1124926121 gopurs_runtime.Value
var once_mappend___1124926121 sync.Once
func Get_mappend___1124926121() gopurs_runtime.Value {
	once_mappend___1124926121.Do(func() {
		cache_mappend___1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mappend___1124926121(gopurs_runtime.CoerceToStruct[pkg_Test_Polymorphism.Constructor_Monoidish[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mappend___1124926121
}

var cache_mappend___3566619927 gopurs_runtime.Value
var once_mappend___3566619927 sync.Once
func Get_mappend___3566619927() gopurs_runtime.Value {
	once_mappend___3566619927.Do(func() {
		cache_mappend___3566619927 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mappend___3566619927(gopurs_runtime.CoerceToStruct[pkg_Test_Polymorphism.Constructor_Monoidish[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mappend___3566619927
}

var cache_mempty___1556010056 gopurs_runtime.Value
var once_mempty___1556010056 sync.Once
func Get_mempty___1556010056() gopurs_runtime.Value {
	once_mempty___1556010056.Do(func() {
		cache_mempty___1556010056 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty___1556010056(dict_0_box)
})
	})
	return cache_mempty___1556010056
}

var cache_mempty___1540866998 gopurs_runtime.Value
var once_mempty___1540866998 sync.Once
func Get_mempty___1540866998() gopurs_runtime.Value {
	once_mempty___1540866998.Do(func() {
		cache_mempty___1540866998 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty___1540866998(dict_0_box)
})
	})
	return cache_mempty___1540866998
}

var cache_polyLoop__1533381815 gopurs_runtime.Value
var once_polyLoop__1533381815 sync.Once
func Get_polyLoop__1533381815() gopurs_runtime.Value {
	once_polyLoop__1533381815.Do(func() {
		cache_polyLoop__1533381815 = gopurs_runtime.Func2(func(n_init_0_box gopurs_runtime.Value, acc_init_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_polyLoop__1533381815(n_init_0_box.IntVal, acc_init_1_box)
})
	})
	return cache_polyLoop__1533381815
}

var cache_polyLoop__2675791634 gopurs_runtime.Value
var once_polyLoop__2675791634 sync.Once
func Get_polyLoop__2675791634() gopurs_runtime.Value {
	once_polyLoop__2675791634.Do(func() {
		cache_polyLoop__2675791634 = gopurs_runtime.Func3(func(dictMonoidish_0_box gopurs_runtime.Value, n_init_1_box gopurs_runtime.Value, acc_init_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_polyLoop__2675791634(gopurs_runtime.CoerceToStruct[pkg_Test_Polymorphism.Constructor_Monoidish[gopurs_runtime.Value]](dictMonoidish_0_box), n_init_1_box.IntVal, acc_init_2_box)
})
	})
	return cache_polyLoop__2675791634
}

var cache_filter__1481233142 gopurs_runtime.Value
var once_filter__1481233142 sync.Once
func Get_filter__1481233142() gopurs_runtime.Value {
	once_filter__1481233142.Do(func() {
		cache_filter__1481233142 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, lst_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_filter__1481233142(p_0_box, gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](lst_1_box)))}
})
	})
	return cache_filter__1481233142
}

var cache_filter__37320371 gopurs_runtime.Value
var once_filter__37320371 sync.Once
func Get_filter__37320371() gopurs_runtime.Value {
	once_filter__37320371.Do(func() {
		cache_filter__37320371 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, lst_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_filter__37320371(p_0_box, gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](lst_1_box)))}
})
	})
	return cache_filter__37320371
}

var cache_reverse__3119428352 gopurs_runtime.Value
var once_reverse__3119428352 sync.Once
func Get_reverse__3119428352() gopurs_runtime.Value {
	once_reverse__3119428352.Do(func() {
		cache_reverse__3119428352 = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(Call_reverse__3119428352(gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](lst_0_box)))}
})
	})
	return cache_reverse__3119428352
}

var cache_bindState__567439955 gopurs_runtime.Value
var once_bindState__567439955 sync.Once
func Get_bindState__567439955() gopurs_runtime.Value {
	once_bindState__567439955.Do(func() {
		cache_bindState__567439955 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindState__567439955(v_0_box, g_1_box, s_2_box)
})
	})
	return cache_bindState__567439955
}

var cache_bindState__2171045075 gopurs_runtime.Value
var once_bindState__2171045075 sync.Once
func Get_bindState__2171045075() gopurs_runtime.Value {
	once_bindState__2171045075.Do(func() {
		cache_bindState__2171045075 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindState__2171045075(v_0_box, g_1_box, s_2_box)
})
	})
	return cache_bindState__2171045075
}

var cache_bindState__3267751411 gopurs_runtime.Value
var once_bindState__3267751411 sync.Once
func Get_bindState__3267751411() gopurs_runtime.Value {
	once_bindState__3267751411.Do(func() {
		cache_bindState__3267751411 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindState__3267751411(v_0_box, g_1_box, s_2_box)
})
	})
	return cache_bindState__3267751411
}

var cache_get__676984528 gopurs_runtime.Value
var once_get__676984528 sync.Once
func Get_get__676984528() gopurs_runtime.Value {
	once_get__676984528.Do(func() {
		cache_get__676984528 = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get__676984528(s_0_box)
})
	})
	return cache_get__676984528
}

var cache_modify__1175978184 gopurs_runtime.Value
var once_modify__1175978184 sync.Once
func Get_modify__1175978184() gopurs_runtime.Value {
	once_modify__1175978184.Do(func() {
		cache_modify__1175978184 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__1175978184(f_0_box, s_1_box)
})
	})
	return cache_modify__1175978184
}

var cache_modify__3050914184 gopurs_runtime.Value
var once_modify__3050914184 sync.Once
func Get_modify__3050914184() gopurs_runtime.Value {
	once_modify__3050914184.Do(func() {
		cache_modify__3050914184 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__3050914184(f_0_box, s_1_box)
})
	})
	return cache_modify__3050914184
}

var cache_pureState__608762702 gopurs_runtime.Value
var once_pureState__608762702 sync.Once
func Get_pureState__608762702() gopurs_runtime.Value {
	once_pureState__608762702.Do(func() {
		cache_pureState__608762702 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pureState__608762702(a_0_box, s_1_box)
})
	})
	return cache_pureState__608762702
}

var cache_pureState__1329830318 gopurs_runtime.Value
var once_pureState__1329830318 sync.Once
func Get_pureState__1329830318() gopurs_runtime.Value {
	once_pureState__1329830318.Do(func() {
		cache_pureState__1329830318 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pureState__1329830318(a_0_box, s_1_box)
})
	})
	return cache_pureState__1329830318
}

var cache_put__3685210848 gopurs_runtime.Value
var once_put__3685210848 sync.Once
func Get_put__3685210848() gopurs_runtime.Value {
	once_put__3685210848.Do(func() {
		cache_put__3685210848 = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_put__3685210848(s_0_box, v_1_box)
})
	})
	return cache_put__3685210848
}

var cache_runState__2373419117 gopurs_runtime.Value
var once_runState__2373419117 sync.Once
func Get_runState__2373419117() gopurs_runtime.Value {
	once_runState__2373419117.Do(func() {
		cache_runState__2373419117 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runState__2373419117(v_0_box, s_1_box)
})
	})
	return cache_runState__2373419117
}

var cache_runState__3059282509 gopurs_runtime.Value
var once_runState__3059282509 sync.Once
func Get_runState__3059282509() gopurs_runtime.Value {
	once_runState__3059282509.Do(func() {
		cache_runState__3059282509 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runState__3059282509(v_0_box, s_1_box)
})
	})
	return cache_runState__3059282509
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__629383158(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __eta0_0
})
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__3550378017(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_bind__1949526049(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_bind__3328406721(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Value{})
_ = __local_var_2_0
return gopurs_runtime.Apply(gopurs_runtime.Apply(__eta1_1, __local_var_2_0), gopurs_runtime.Value{})
})
}

func Call_discard__203210016(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), __eta0_0, __eta1_1)
}

func Call_discard__317162198(dict_0_loop *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_filter__4047711382(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_filterImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_filter__377906483(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_filterImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_eq__789642299(dict_0_loop *pkg_Data_Eq.Constructor_Eq[bool]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[bool] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__1697837627(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__2843686287(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool((__eta0_0.IntVal) == (__eta1_1.IntVal))
}

func Call_eq__2276491096(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) == ((__eta1_1.IntVal) != (0)))
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_notEq__2843686287(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
return ((x_0.IntVal) == (y_1.IntVal)) != (true)
}

func Call_notEq__2384498378(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) bool {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
return ((gopurs_runtime.Apply2(dictEq_0.V0, x_1, y_2).IntVal) != (0)) != (true)
}

func Call_div__1002719800(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Float((__eta0_0.FloatVal()) / (__eta1_1.FloatVal()))
}

func Call_div__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_mod__2185172824(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), __eta0_0, __eta1_1)
}

func Call_mod__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_compare__669572705(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_greaterThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_lessThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_lessThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_sub__2927892844(dict_0_loop *pkg_Data_Ring.Constructor_Ring[int64]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[int64] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__1124926121(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__1135378904(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Float((__eta0_0.FloatVal()) - (__eta1_1.FloatVal()))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) + (__eta1_1.IntVal))
}

func Call_add__137136408(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Float((__eta0_0.FloatVal()) + (__eta1_1.FloatVal()))
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mul__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) * (__eta1_1.IntVal))
}

func Call_mul__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_show__3754018243(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__3984012258(dict_0_loop *pkg_Data_Show.Constructor_Show[*pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[*pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_logShow__2885109999(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), a_0).StrVal()))
}

func Call_logShow__339054415(dictShow_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dictShow_0_loop
_ = dictShow_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(gopurs_runtime.Apply(dictShow_0.V0, a_1).StrVal()))
}

func Call_mulC__1746928225(m_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var n_1 gopurs_runtime.Value = n_1_loop
_ = n_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
return gopurs_runtime.Apply2(m_0, gopurs_runtime.Apply(n_1, f_2), x_3)
}

func Call_mulC__3596604257(m_0_loop gopurs_runtime.Value, n_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var m_0 gopurs_runtime.Value = m_0_loop
_ = m_0
var n_1 gopurs_runtime.Value = n_1_loop
_ = n_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
return gopurs_runtime.Apply2(m_0, gopurs_runtime.Apply(n_1, f_2), x_3)
}

func Call_succC__952275393(n_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(n_0, f_1, x_2))
}

func Call_succC__1461826241(n_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply2(n_0, f_1, x_2))
}

func Call_zeroC__4066693242(v_0_loop gopurs_runtime.Value, x_1_loop int64) int64 {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 int64 = x_1_loop
_ = x_1
return x_1
}

func Call_force__3902501304(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
}

func Call_force__721037880(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
}

func Call_foldl__3262866295(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Test_ListOps.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Test_ListOps.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply3(pkg_Test_ListOps.Get_foldl(), v_0, gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Test_ListOps.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*pkg_Test_ListOps.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_foldl__1255354935(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Test_ListOps.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Test_ListOps.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1127792131 && gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply3(pkg_Test_ListOps.Get_foldl(), v_0, gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Test_ListOps.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer((*pkg_Test_ListOps.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1127792131, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_mappend___1124926121(dict_0_loop *pkg_Test_Polymorphism.Constructor_Monoidish[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Test_Polymorphism.Constructor_Monoidish[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mappend___3566619927(dict_0_loop *pkg_Test_Polymorphism.Constructor_Monoidish[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Test_Polymorphism.Constructor_Monoidish[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mempty___1556010056(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_mempty___1540866998(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_polyLoop__1533381815(n_init_0_loop int64, acc_init_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_init_0 int64 = n_init_0_loop
_ = n_init_0
var acc_init_1 gopurs_runtime.Value = acc_init_1_loop
_ = acc_init_1
var go__go_2_0_0 gopurs_runtime.Value
go__go_2_0_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop int64 = v_3_loop_val.IntVal
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_0_0:
for {
if false { continue go__go_2_0_0 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if (v_3) == (0) {
__t1 = v1_4
goto end_branch_1
} else {

}
}
{
v_3_loop = (v_3) - (1)
v1_4_loop = gopurs_runtime.Int((v1_4.IntVal) + (1))
continue go__go_2_0_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply2(go__go_2_0_0, gopurs_runtime.Int(n_init_0), acc_init_1)
}

func Call_polyLoop__2675791634(dictMonoidish_0_loop *pkg_Test_Polymorphism.Constructor_Monoidish[gopurs_runtime.Value], n_init_1_loop int64, acc_init_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoidish_0 *pkg_Test_Polymorphism.Constructor_Monoidish[gopurs_runtime.Value] = dictMonoidish_0_loop
_ = dictMonoidish_0
var n_init_1 int64 = n_init_1_loop
_ = n_init_1
var acc_init_2 gopurs_runtime.Value = acc_init_2_loop
_ = acc_init_2
var go__go_3_0_1 gopurs_runtime.Value
go__go_3_0_1 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop int64 = v_4_loop_val.IntVal
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_1:
for {
if false { continue go__go_3_0_1 }
var v_4 int64 = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (v_4) == (0) {
__t1 = v1_5
goto end_branch_1
} else {

}
}
{
v_4_loop = (v_4) - (1)
v1_5_loop = gopurs_runtime.Apply2(dictMonoidish_0.V0, v1_5, dictMonoidish_0.V1)
continue go__go_3_0_1
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
return gopurs_runtime.Apply2(go__go_3_0_1, gopurs_runtime.Int(n_init_1), acc_init_2)
}

func Call_filter__1481233142(p_0_loop gopurs_runtime.Value, lst_1_loop *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]) *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var lst_1 *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = lst_1_loop
_ = lst_1
var go__go_2_0_2 gopurs_runtime.Value
go__go_2_0_2 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](v1_4_loop_val)
go__go_2_0_2:
for {
if false { continue go__go_2_0_2 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t4 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr == nil) {
var go__go_5_1_3 gopurs_runtime.Value
go__go_5_1_3 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](v1_7_loop_val)
go__go_5_1_3:
for {
if false { continue go__go_5_1_3 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = v1_7_loop
_ = v1_7
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v1_7)}
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V1)}
v1_7_loop = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0, v1_7})})
continue go__go_5_1_3
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_5_1_3, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v1_4)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(nil))})))}
goto end_branch_4
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr != nil) {
var __t3 *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(p_0, (*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
v1_4_loop = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, v1_4})})
continue go__go_2_0_2
__t3 = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
v1_4_loop = v1_4
continue go__go_2_0_2
__t3 = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t3)}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_2, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(nil))}))
}

func Call_filter__37320371(p_0_loop gopurs_runtime.Value, lst_1_loop *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]) *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var lst_1 *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = lst_1_loop
_ = lst_1
var go__go_2_0_4 gopurs_runtime.Value
go__go_2_0_4 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](v1_4_loop_val)
go__go_2_0_4:
for {
if false { continue go__go_2_0_4 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t4 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr == nil) {
var go__go_5_1_5 gopurs_runtime.Value
go__go_5_1_5 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](v1_7_loop_val)
go__go_5_1_5:
for {
if false { continue go__go_5_1_5 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = v1_7_loop
_ = v1_7
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v1_7)}
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 2390177629 && v_6.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V1)}
v1_7_loop = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0, v1_7})})
continue go__go_5_1_5
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_5_1_5, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v1_4)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(nil))})))}
goto end_branch_4
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 2390177629 && v_3.UnsafePtr != nil) {
var __t3 *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(p_0, (*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
v1_4_loop = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0, v1_4})})
continue go__go_2_0_4
__t3 = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
v1_4_loop = v1_4
continue go__go_2_0_4
__t3 = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(__t3)}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(nil))}))
}

func Call_reverse__3119428352(lst_0_loop *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]) *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] {
var lst_0 *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = lst_0_loop
_ = lst_0
var go__go_1_0_6 gopurs_runtime.Value
go__go_1_0_6 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
go__go_1_0_6:
for {
if false { continue go__go_1_0_6 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(v1_3)}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2390177629 && v_2.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(&pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3})})
continue go__go_1_0_6
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_1_0_6, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2390177629, UnsafePtr: unsafe.Pointer((*pkg_Test_Primes.Constructor_Cons[gopurs_runtime.Value])(nil))}))
}

func Call_bindState__567439955(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_bindState__2171045075(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_bindState__3267751411(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_get__676984528(s_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
return gopurs_runtime.RecordDict2("state", "val", s_0, s_0)
}

func Call_modify__1175978184(f_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Apply(f_0, s_1), pkg_Data_Unit.Get_unit())
}

func Call_modify__3050914184(f_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Apply(f_0, s_1), pkg_Data_Unit.Get_unit())
}

func Call_pureState__608762702(a_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", s_1, a_0)
}

func Call_pureState__1329830318(a_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", s_1, a_0)
}

func Call_put__3685210848(s_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.RecordDict2("state", "val", s_0, pkg_Data_Unit.Get_unit())
}

func Call_runState__2373419117(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(v_0, s_1)
}

func Call_runState__3059282509(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(v_0, s_1)
}


