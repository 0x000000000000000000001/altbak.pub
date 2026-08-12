package Test_ArrayOps

import (
	pkg_Bench "gopurs/output/Bench"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Array "gopurs/output/Data.Array"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_go__range gopurs_runtime.Value
var once_go__range sync.Once
func Get_go__range() gopurs_runtime.Value {
	once_go__range.Do(func() {
		cache_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_go__range(start_0_box.IntVal, end_1_box.IntVal)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_go__range
}

var cache_filterEvens gopurs_runtime.Value
var once_filterEvens sync.Once
func Get_filterEvens() gopurs_runtime.Value {
	once_filterEvens.Do(func() {
		cache_filterEvens = gopurs_runtime.Func(func(arr_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_filterEvens(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(arr_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_filterEvens
}

var cache_sumEvens gopurs_runtime.Value
var once_sumEvens sync.Once
func Get_sumEvens() gopurs_runtime.Value {
	once_sumEvens.Do(func() {
		cache_sumEvens = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_sumEvens(n_0_box.IntVal))
})
	})
	return cache_sumEvens
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Array Processing (900 elements):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(Get_bind__3550378017(), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(900)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_sumEvens(dummy_0.IntVal))))
}))
	})
	return cache_act
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
		cache_bind__3550378017 = gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind")
	})
	return cache_bind__3550378017
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

var cache_eq__2843686287 gopurs_runtime.Value
var once_eq__2843686287 sync.Once
func Get_eq__2843686287() gopurs_runtime.Value {
	once_eq__2843686287.Do(func() {
		cache_eq__2843686287 = pkg_Data_Eq.Get_eqIntImpl()
	})
	return cache_eq__2843686287
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

var cache_mod__2185172824 gopurs_runtime.Value
var once_mod__2185172824 sync.Once
func Get_mod__2185172824() gopurs_runtime.Value {
	once_mod__2185172824.Do(func() {
		cache_mod__2185172824 = gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod")
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

func Call_go__range(start_0_loop int64, end_1_loop int64) []int64 {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_rangeImpl(), gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_filterEvens(arr_0_loop []int64) []int64 {
var arr_0 []int64 = arr_0_loop
_ = arr_0
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_filter__4047711382(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Apply2(Get_mod__2185172824(), x_1, gopurs_runtime.Int(2)), gopurs_runtime.Int(0)).IntVal) != (0))
}), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := arr_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_sumEvens(n_0_loop int64) int64 {
var n_0 int64 = n_0_loop
_ = n_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl"), pkg_Data_Semiring.Get_intAdd(), gopurs_runtime.Int(0), gopurs_runtime.Array(Call_filter__4047711382(gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Apply2(Get_mod__2185172824(), x_1, gopurs_runtime.Int(2)), gopurs_runtime.Int(0)).IntVal) != (0))
}), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_rangeImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(n_0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mod__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}


