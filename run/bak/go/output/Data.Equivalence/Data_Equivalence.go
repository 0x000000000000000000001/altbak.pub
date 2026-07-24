package Data_Equivalence

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var Equivalence gopurs_runtime.Value
var once_Equivalence sync.Once
func Get_Equivalence() gopurs_runtime.Value {
	once_Equivalence.Do(func() {
		Equivalence = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return Equivalence
}

var semigroupEquivalence gopurs_runtime.Value
var once_semigroupEquivalence sync.Once
func Get_semigroupEquivalence() gopurs_runtime.Value {
	once_semigroupEquivalence.Do(func() {
		semigroupEquivalence = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func4(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(v_0, a_2, b_3).IntVal != 0 && gopurs_runtime.Apply2(v1_1, a_2, b_3).IntVal != 0)
}))
	})
	return semigroupEquivalence
}

var newtypeEquivalence gopurs_runtime.Value
var once_newtypeEquivalence sync.Once
func Get_newtypeEquivalence() gopurs_runtime.Value {
	once_newtypeEquivalence.Do(func() {
		newtypeEquivalence = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeEquivalence
}

var monoidEquivalence gopurs_runtime.Value
var once_monoidEquivalence sync.Once
func Get_monoidEquivalence() gopurs_runtime.Value {
	once_monoidEquivalence.Do(func() {
		monoidEquivalence = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupEquivalence()
}))
	})
	return monoidEquivalence
}

var defaultEquivalence gopurs_runtime.Value
var once_defaultEquivalence sync.Once
func Get_defaultEquivalence() gopurs_runtime.Value {
	once_defaultEquivalence.Do(func() {
		defaultEquivalence = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordGet(dictEq_0, "eq")
}()
})
	})
	return defaultEquivalence
}

var contravariantEquivalence gopurs_runtime.Value
var once_contravariantEquivalence sync.Once
func Get_contravariantEquivalence() gopurs_runtime.Value {
	once_contravariantEquivalence.Do(func() {
		contravariantEquivalence = gopurs_runtime.RecordDict1("cmap", gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(v_1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3))
}))
	})
	return contravariantEquivalence
}

var comparisonEquivalence gopurs_runtime.Value
var once_comparisonEquivalence sync.Once
func Get_comparisonEquivalence() gopurs_runtime.Value {
	once_comparisonEquivalence.Do(func() {
		comparisonEquivalence = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_comparisonEquivalence(v_0_box, a_1_box, b_2_box))
})
	})
	return comparisonEquivalence
}

func Call_comparisonEquivalence(v_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) bool {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return (gopurs_runtime.Apply2(v_0, a_1, b_2).Type == 9 && gopurs_runtime.Apply2(v_0, a_1, b_2).IntVal == 1111389260)
}


