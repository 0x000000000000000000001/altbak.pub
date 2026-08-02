package AppX

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Test_AstTree "gopurs/output/Test.AstTree"
	pkg_Test_Fib "gopurs/output/Test.Fib"
	pkg_Test_ListOps "gopurs/output/Test.ListOps"
	pkg_Test_TCO "gopurs/output/Test.TCO"
	pkg_Test_Records "gopurs/output/Test.Records"
	pkg_Test_Ackermann "gopurs/output/Test.Ackermann"
	pkg_Test_Church "gopurs/output/Test.Church"
	pkg_Test_Primes "gopurs/output/Test.Primes"
	pkg_Test_RBTree "gopurs/output/Test.RBTree"
	pkg_Test_Polymorphism "gopurs/output/Test.Polymorphism"
	pkg_Test_StateMonad "gopurs/output/Test.StateMonad"
	pkg_Test_LazyEvaluation "gopurs/output/Test.LazyEvaluation"
	pkg_Test_ArrayOps "gopurs/output/Test.ArrayOps"
	pkg_Test_FileOps "gopurs/output/Test.FileOps"
	pkg_Test_STArray "gopurs/output/Test.STArray"
	pkg_Test_StringOps "gopurs/output/Test.StringOps"
	pkg_Test_AffOperations "gopurs/output/Test.AffOperations"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
)

var cache_main gopurs_runtime.Value
var once_main sync.Once
func Get_main() gopurs_runtime.Value {
	once_main.Do(func() {
		cache_main = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_AstTree.Get_describe(), pkg_Test_AstTree.Get_act()), gopurs_runtime.Func(func(t1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Fib.Get_describe(), pkg_Test_Fib.Get_act()), gopurs_runtime.Func(func(t2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_ListOps.Get_describe(), pkg_Test_ListOps.Get_act()), gopurs_runtime.Func(func(t3_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_TCO.Get_describe(), pkg_Test_TCO.Get_act()), gopurs_runtime.Func(func(t4_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Records.Get_describe(), pkg_Test_Records.Get_act()), gopurs_runtime.Func(func(t5_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Ackermann.Get_describe(), pkg_Test_Ackermann.Get_act()), gopurs_runtime.Func(func(t6_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Church.Get_describe(), pkg_Test_Church.Get_act()), gopurs_runtime.Func(func(t7_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Primes.Get_describe(), pkg_Test_Primes.Get_act()), gopurs_runtime.Func(func(t8_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_RBTree.Get_describe(), pkg_Test_RBTree.Get_act()), gopurs_runtime.Func(func(t9_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Polymorphism.Get_describe(), pkg_Test_Polymorphism.Get_act()), gopurs_runtime.Func(func(t10_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_StateMonad.Get_describe(), pkg_Test_StateMonad.Get_act()), gopurs_runtime.Func(func(t11_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_LazyEvaluation.Get_describe(), pkg_Test_LazyEvaluation.Get_act()), gopurs_runtime.Func(func(t12_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_ArrayOps.Get_describe(), pkg_Test_ArrayOps.Get_act()), gopurs_runtime.Func(func(t13_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_FileOps.Get_describe(), pkg_Test_FileOps.Get_act()), gopurs_runtime.Func(func(t14_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_STArray.Get_describe(), pkg_Test_STArray.Get_act()), gopurs_runtime.Func(func(t15_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_StringOps.Get_describe(), pkg_Test_StringOps.Get_act()), gopurs_runtime.Func(func(t16_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_AffOperations.Get_describe(), pkg_Test_AffOperations.Get_act()), gopurs_runtime.Func(func(t17_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("Total exec time: "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(pkg_Bench.Get_formatNumber(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t1_0, gopurs_runtime.Float(1000.0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t2_1, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t3_2, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t4_3, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t5_4, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t6_5, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t7_6, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t8_7, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t9_8, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t10_9, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t11_10, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t12_11, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t13_12, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t14_13, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t15_14, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t16_15, gopurs_runtime.Float(1000.0))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingNumber(), "div"), t17_16, gopurs_runtime.Float(1000.0)))), gopurs_runtime.Str(" ms\x0a"))))
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
}))
}))
}))
}))
	})
	return cache_main
}




