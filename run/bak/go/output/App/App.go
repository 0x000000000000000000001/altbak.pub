package App

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
)

var main gopurs_runtime.Value
var once_main sync.Once
func Get_main() gopurs_runtime.Value {
	once_main.Do(func() {
		main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
_dollar__unused_0_0 := gopurs_runtime.Apply(pkg_Test_AstTree.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_0_0
_dollar__unused_1_1 := gopurs_runtime.Apply(pkg_Test_AstTree.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_1_1
_dollar__unused_2_2 := gopurs_runtime.Apply(pkg_Test_Fib.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_2_2
_dollar__unused_3_3 := gopurs_runtime.Apply(pkg_Test_Fib.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_3_3
_dollar__unused_4_4 := gopurs_runtime.Apply(pkg_Test_ListOps.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_4_4
_dollar__unused_5_5 := gopurs_runtime.Apply(pkg_Test_ListOps.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_5_5
_dollar__unused_6_6 := gopurs_runtime.Apply(pkg_Test_TCO.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_6_6
_dollar__unused_7_7 := gopurs_runtime.Apply(pkg_Test_TCO.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_7_7
_dollar__unused_8_8 := gopurs_runtime.Apply(pkg_Test_Records.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_8_8
_dollar__unused_9_9 := gopurs_runtime.Apply(pkg_Test_Records.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_9_9
_dollar__unused_10_10 := gopurs_runtime.Apply(pkg_Test_Ackermann.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_10_10
_dollar__unused_11_11 := gopurs_runtime.Apply(pkg_Test_Ackermann.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_11_11
_dollar__unused_12_12 := gopurs_runtime.Apply(pkg_Test_Church.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_12_12
_dollar__unused_13_13 := gopurs_runtime.Apply(pkg_Test_Church.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_13_13
_dollar__unused_14_14 := gopurs_runtime.Apply(pkg_Test_Primes.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_14_14
_dollar__unused_15_15 := gopurs_runtime.Apply(pkg_Test_Primes.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_15_15
_dollar__unused_16_16 := gopurs_runtime.Apply(pkg_Test_RBTree.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_16_16
_dollar__unused_17_17 := gopurs_runtime.Apply(pkg_Test_RBTree.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_17_17
_dollar__unused_18_18 := gopurs_runtime.Apply(pkg_Test_Polymorphism.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_18_18
_dollar__unused_19_19 := gopurs_runtime.Apply(pkg_Test_Polymorphism.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_19_19
_dollar__unused_20_20 := gopurs_runtime.Apply(pkg_Test_StateMonad.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_20_20
_dollar__unused_21_21 := gopurs_runtime.Apply(pkg_Test_StateMonad.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_21_21
_dollar__unused_22_22 := gopurs_runtime.Apply(pkg_Test_LazyEvaluation.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_22_22
return gopurs_runtime.Apply(pkg_Test_LazyEvaluation.Get_act(), gopurs_runtime.Value{})
})
	})
	return main
}


