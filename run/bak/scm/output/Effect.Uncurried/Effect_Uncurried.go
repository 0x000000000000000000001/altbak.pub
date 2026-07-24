package Effect_Uncurried

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var semigroupEffectFn9 gopurs_runtime.Value
var once_semigroupEffectFn9 sync.Once
func Get_semigroupEffectFn9() gopurs_runtime.Value {
	once_semigroupEffectFn9.Do(func() {
		semigroupEffectFn9 = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_1 gopurs_runtime.Value, f2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_12_0 := gopurs_runtime.UncurriedApp(f1_1, a_3, b_4, c_5, d_6, e_7, f_8, g_9, h_10, i_11)
_ = a_prime_12_0
a_prime_13_1 := gopurs_runtime.UncurriedApp(f2_2, a_3, b_4, c_5, d_6, e_7, f_8, g_9, h_10, i_11)
_ = a_prime_13_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), a_prime_12_0, a_prime_13_1)
})
})
})
})
})
})
})
})
})
})
}))
}()
})
	})
	return semigroupEffectFn9
}

var semigroupEffectFn8 gopurs_runtime.Value
var once_semigroupEffectFn8 sync.Once
func Get_semigroupEffectFn8() gopurs_runtime.Value {
	once_semigroupEffectFn8.Do(func() {
		semigroupEffectFn8 = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_1 gopurs_runtime.Value, f2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_11_0 := gopurs_runtime.UncurriedApp(f1_1, a_3, b_4, c_5, d_6, e_7, f_8, g_9, h_10)
_ = a_prime_11_0
a_prime_12_1 := gopurs_runtime.UncurriedApp(f2_2, a_3, b_4, c_5, d_6, e_7, f_8, g_9, h_10)
_ = a_prime_12_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), a_prime_11_0, a_prime_12_1)
})
})
})
})
})
})
})
})
})
}))
}()
})
	})
	return semigroupEffectFn8
}

var semigroupEffectFn7 gopurs_runtime.Value
var once_semigroupEffectFn7 sync.Once
func Get_semigroupEffectFn7() gopurs_runtime.Value {
	once_semigroupEffectFn7.Do(func() {
		semigroupEffectFn7 = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_1 gopurs_runtime.Value, f2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_10_0 := gopurs_runtime.UncurriedApp(f1_1, a_3, b_4, c_5, d_6, e_7, f_8, g_9)
_ = a_prime_10_0
a_prime_11_1 := gopurs_runtime.UncurriedApp(f2_2, a_3, b_4, c_5, d_6, e_7, f_8, g_9)
_ = a_prime_11_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), a_prime_10_0, a_prime_11_1)
})
})
})
})
})
})
})
})
}))
}()
})
	})
	return semigroupEffectFn7
}

var semigroupEffectFn6 gopurs_runtime.Value
var once_semigroupEffectFn6 sync.Once
func Get_semigroupEffectFn6() gopurs_runtime.Value {
	once_semigroupEffectFn6.Do(func() {
		semigroupEffectFn6 = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_1 gopurs_runtime.Value, f2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_9_0 := gopurs_runtime.UncurriedApp(f1_1, a_3, b_4, c_5, d_6, e_7, f_8)
_ = a_prime_9_0
a_prime_10_1 := gopurs_runtime.UncurriedApp(f2_2, a_3, b_4, c_5, d_6, e_7, f_8)
_ = a_prime_10_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), a_prime_9_0, a_prime_10_1)
})
})
})
})
})
})
})
}))
}()
})
	})
	return semigroupEffectFn6
}

var semigroupEffectFn5 gopurs_runtime.Value
var once_semigroupEffectFn5 sync.Once
func Get_semigroupEffectFn5() gopurs_runtime.Value {
	once_semigroupEffectFn5.Do(func() {
		semigroupEffectFn5 = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_1 gopurs_runtime.Value, f2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func5(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, c_5 gopurs_runtime.Value, d_6 gopurs_runtime.Value, e_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_8_0 := gopurs_runtime.UncurriedApp5(f1_1, a_3, b_4, c_5, d_6, e_7)
_ = a_prime_8_0
a_prime_9_1 := gopurs_runtime.UncurriedApp5(f2_2, a_3, b_4, c_5, d_6, e_7)
_ = a_prime_9_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), a_prime_8_0, a_prime_9_1)
})
})
}))
}()
})
	})
	return semigroupEffectFn5
}

var semigroupEffectFn4 gopurs_runtime.Value
var once_semigroupEffectFn4 sync.Once
func Get_semigroupEffectFn4() gopurs_runtime.Value {
	once_semigroupEffectFn4.Do(func() {
		semigroupEffectFn4 = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_1 gopurs_runtime.Value, f2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func4(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, c_5 gopurs_runtime.Value, d_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_7_0 := gopurs_runtime.UncurriedApp4(f1_1, a_3, b_4, c_5, d_6)
_ = a_prime_7_0
a_prime_8_1 := gopurs_runtime.UncurriedApp4(f2_2, a_3, b_4, c_5, d_6)
_ = a_prime_8_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), a_prime_7_0, a_prime_8_1)
})
})
}))
}()
})
	})
	return semigroupEffectFn4
}

var semigroupEffectFn3 gopurs_runtime.Value
var once_semigroupEffectFn3 sync.Once
func Get_semigroupEffectFn3() gopurs_runtime.Value {
	once_semigroupEffectFn3.Do(func() {
		semigroupEffectFn3 = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_1 gopurs_runtime.Value, f2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func3(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_6_0 := gopurs_runtime.UncurriedApp3(f1_1, a_3, b_4, c_5)
_ = a_prime_6_0
a_prime_7_1 := gopurs_runtime.UncurriedApp3(f2_2, a_3, b_4, c_5)
_ = a_prime_7_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), a_prime_6_0, a_prime_7_1)
})
})
}))
}()
})
	})
	return semigroupEffectFn3
}

var semigroupEffectFn2 gopurs_runtime.Value
var once_semigroupEffectFn2 sync.Once
func Get_semigroupEffectFn2() gopurs_runtime.Value {
	once_semigroupEffectFn2.Do(func() {
		semigroupEffectFn2 = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_1 gopurs_runtime.Value, f2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_5_0 := gopurs_runtime.UncurriedApp2(f1_1, a_3, b_4)
_ = a_prime_5_0
a_prime_6_1 := gopurs_runtime.UncurriedApp2(f2_2, a_3, b_4)
_ = a_prime_6_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), a_prime_5_0, a_prime_6_1)
})
})
}))
}()
})
	})
	return semigroupEffectFn2
}

var semigroupEffectFn10 gopurs_runtime.Value
var once_semigroupEffectFn10 sync.Once
func Get_semigroupEffectFn10() gopurs_runtime.Value {
	once_semigroupEffectFn10.Do(func() {
		semigroupEffectFn10 = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_1 gopurs_runtime.Value, f2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_13_0 := gopurs_runtime.UncurriedApp(f1_1, a_3, b_4, c_5, d_6, e_7, f_8, g_9, h_10, i_11, j_12)
_ = a_prime_13_0
a_prime_14_1 := gopurs_runtime.UncurriedApp(f2_2, a_3, b_4, c_5, d_6, e_7, f_8, g_9, h_10, i_11, j_12)
_ = a_prime_14_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), a_prime_13_0, a_prime_14_1)
})
})
})
})
})
})
})
})
})
})
})
}))
}()
})
	})
	return semigroupEffectFn10
}

var semigroupEffectFn1 gopurs_runtime.Value
var once_semigroupEffectFn1 sync.Once
func Get_semigroupEffectFn1() gopurs_runtime.Value {
	once_semigroupEffectFn1.Do(func() {
		semigroupEffectFn1 = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_1 gopurs_runtime.Value, f2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_4_0 := gopurs_runtime.UncurriedApp(f1_1, a_3)
_ = a_prime_4_0
a_prime_5_1 := gopurs_runtime.UncurriedApp(f2_2, a_3)
_ = a_prime_5_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), a_prime_4_0, a_prime_5_1)
})
})
}))
}()
})
	})
	return semigroupEffectFn1
}

var monoidEffectFn9 gopurs_runtime.Value
var once_monoidEffectFn9 sync.Once
func Get_monoidEffectFn9() gopurs_runtime.Value {
	once_monoidEffectFn9.Do(func() {
		monoidEffectFn9 = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.RecordGet(dictMonoid_0_loop, "mempty")
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupEffectFn91_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_3 gopurs_runtime.Value, f2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_14_3 := gopurs_runtime.UncurriedApp(f1_3, a_5, b_6, c_7, d_8, e_9, f_10, g_11, h_12, i_13)
_ = a_prime_14_3
a_prime_15_4 := gopurs_runtime.UncurriedApp(f2_4, a_5, b_6, c_7, d_8, e_9, f_10, g_11, h_12, i_13)
_ = a_prime_15_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), a_prime_14_3, a_prime_15_4)
})
})
})
})
})
})
})
})
})
})
}))
_ = semigroupEffectFn91_3_2
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v5_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v6_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v7_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v8_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
})
})
})
})
})
})
})
})
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffectFn91_3_2
}))
}()
})
	})
	return monoidEffectFn9
}

var monoidEffectFn8 gopurs_runtime.Value
var once_monoidEffectFn8 sync.Once
func Get_monoidEffectFn8() gopurs_runtime.Value {
	once_monoidEffectFn8.Do(func() {
		monoidEffectFn8 = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.RecordGet(dictMonoid_0_loop, "mempty")
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupEffectFn81_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_3 gopurs_runtime.Value, f2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_13_3 := gopurs_runtime.UncurriedApp(f1_3, a_5, b_6, c_7, d_8, e_9, f_10, g_11, h_12)
_ = a_prime_13_3
a_prime_14_4 := gopurs_runtime.UncurriedApp(f2_4, a_5, b_6, c_7, d_8, e_9, f_10, g_11, h_12)
_ = a_prime_14_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), a_prime_13_3, a_prime_14_4)
})
})
})
})
})
})
})
})
})
}))
_ = semigroupEffectFn81_3_2
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v5_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v6_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v7_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
})
})
})
})
})
})
})
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffectFn81_3_2
}))
}()
})
	})
	return monoidEffectFn8
}

var monoidEffectFn7 gopurs_runtime.Value
var once_monoidEffectFn7 sync.Once
func Get_monoidEffectFn7() gopurs_runtime.Value {
	once_monoidEffectFn7.Do(func() {
		monoidEffectFn7 = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.RecordGet(dictMonoid_0_loop, "mempty")
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupEffectFn71_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_3 gopurs_runtime.Value, f2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_12_3 := gopurs_runtime.UncurriedApp(f1_3, a_5, b_6, c_7, d_8, e_9, f_10, g_11)
_ = a_prime_12_3
a_prime_13_4 := gopurs_runtime.UncurriedApp(f2_4, a_5, b_6, c_7, d_8, e_9, f_10, g_11)
_ = a_prime_13_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), a_prime_12_3, a_prime_13_4)
})
})
})
})
})
})
})
})
}))
_ = semigroupEffectFn71_3_2
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v5_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v6_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
})
})
})
})
})
})
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffectFn71_3_2
}))
}()
})
	})
	return monoidEffectFn7
}

var monoidEffectFn6 gopurs_runtime.Value
var once_monoidEffectFn6 sync.Once
func Get_monoidEffectFn6() gopurs_runtime.Value {
	once_monoidEffectFn6.Do(func() {
		monoidEffectFn6 = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.RecordGet(dictMonoid_0_loop, "mempty")
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupEffectFn61_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_3 gopurs_runtime.Value, f2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_11_3 := gopurs_runtime.UncurriedApp(f1_3, a_5, b_6, c_7, d_8, e_9, f_10)
_ = a_prime_11_3
a_prime_12_4 := gopurs_runtime.UncurriedApp(f2_4, a_5, b_6, c_7, d_8, e_9, f_10)
_ = a_prime_12_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), a_prime_11_3, a_prime_12_4)
})
})
})
})
})
})
})
}))
_ = semigroupEffectFn61_3_2
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v5_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
})
})
})
})
})
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffectFn61_3_2
}))
}()
})
	})
	return monoidEffectFn6
}

var monoidEffectFn5 gopurs_runtime.Value
var once_monoidEffectFn5 sync.Once
func Get_monoidEffectFn5() gopurs_runtime.Value {
	once_monoidEffectFn5.Do(func() {
		monoidEffectFn5 = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.RecordGet(dictMonoid_0_loop, "mempty")
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupEffectFn51_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_3 gopurs_runtime.Value, f2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func5(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value, c_7 gopurs_runtime.Value, d_8 gopurs_runtime.Value, e_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_10_3 := gopurs_runtime.UncurriedApp5(f1_3, a_5, b_6, c_7, d_8, e_9)
_ = a_prime_10_3
a_prime_11_4 := gopurs_runtime.UncurriedApp5(f2_4, a_5, b_6, c_7, d_8, e_9)
_ = a_prime_11_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), a_prime_10_3, a_prime_11_4)
})
})
}))
_ = semigroupEffectFn51_3_2
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func5(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value, v3_7 gopurs_runtime.Value, v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffectFn51_3_2
}))
}()
})
	})
	return monoidEffectFn5
}

var monoidEffectFn4 gopurs_runtime.Value
var once_monoidEffectFn4 sync.Once
func Get_monoidEffectFn4() gopurs_runtime.Value {
	once_monoidEffectFn4.Do(func() {
		monoidEffectFn4 = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.RecordGet(dictMonoid_0_loop, "mempty")
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupEffectFn41_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_3 gopurs_runtime.Value, f2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func4(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value, c_7 gopurs_runtime.Value, d_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_9_3 := gopurs_runtime.UncurriedApp4(f1_3, a_5, b_6, c_7, d_8)
_ = a_prime_9_3
a_prime_10_4 := gopurs_runtime.UncurriedApp4(f2_4, a_5, b_6, c_7, d_8)
_ = a_prime_10_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), a_prime_9_3, a_prime_10_4)
})
})
}))
_ = semigroupEffectFn41_3_2
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func4(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value, v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffectFn41_3_2
}))
}()
})
	})
	return monoidEffectFn4
}

var monoidEffectFn3 gopurs_runtime.Value
var once_monoidEffectFn3 sync.Once
func Get_monoidEffectFn3() gopurs_runtime.Value {
	once_monoidEffectFn3.Do(func() {
		monoidEffectFn3 = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.RecordGet(dictMonoid_0_loop, "mempty")
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupEffectFn31_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_3 gopurs_runtime.Value, f2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func3(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value, c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_8_3 := gopurs_runtime.UncurriedApp3(f1_3, a_5, b_6, c_7)
_ = a_prime_8_3
a_prime_9_4 := gopurs_runtime.UncurriedApp3(f2_4, a_5, b_6, c_7)
_ = a_prime_9_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), a_prime_8_3, a_prime_9_4)
})
})
}))
_ = semigroupEffectFn31_3_2
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffectFn31_3_2
}))
}()
})
	})
	return monoidEffectFn3
}

var monoidEffectFn2 gopurs_runtime.Value
var once_monoidEffectFn2 sync.Once
func Get_monoidEffectFn2() gopurs_runtime.Value {
	once_monoidEffectFn2.Do(func() {
		monoidEffectFn2 = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.RecordGet(dictMonoid_0_loop, "mempty")
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupEffectFn21_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_3 gopurs_runtime.Value, f2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_7_3 := gopurs_runtime.UncurriedApp2(f1_3, a_5, b_6)
_ = a_prime_7_3
a_prime_8_4 := gopurs_runtime.UncurriedApp2(f2_4, a_5, b_6)
_ = a_prime_8_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), a_prime_7_3, a_prime_8_4)
})
})
}))
_ = semigroupEffectFn21_3_2
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffectFn21_3_2
}))
}()
})
	})
	return monoidEffectFn2
}

var monoidEffectFn10 gopurs_runtime.Value
var once_monoidEffectFn10 sync.Once
func Get_monoidEffectFn10() gopurs_runtime.Value {
	once_monoidEffectFn10.Do(func() {
		monoidEffectFn10 = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.RecordGet(dictMonoid_0_loop, "mempty")
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupEffectFn101_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_3 gopurs_runtime.Value, f2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(d_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_15_3 := gopurs_runtime.UncurriedApp(f1_3, a_5, b_6, c_7, d_8, e_9, f_10, g_11, h_12, i_13, j_14)
_ = a_prime_15_3
a_prime_16_4 := gopurs_runtime.UncurriedApp(f2_4, a_5, b_6, c_7, d_8, e_9, f_10, g_11, h_12, i_13, j_14)
_ = a_prime_16_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), a_prime_15_3, a_prime_16_4)
})
})
})
})
})
})
})
})
})
})
})
}))
_ = semigroupEffectFn101_3_2
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v5_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v6_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v7_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v8_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v9_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
})
})
})
})
})
})
})
})
})
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffectFn101_3_2
}))
}()
})
	})
	return monoidEffectFn10
}

var monoidEffectFn1 gopurs_runtime.Value
var once_monoidEffectFn1 sync.Once
func Get_monoidEffectFn1() gopurs_runtime.Value {
	once_monoidEffectFn1.Do(func() {
		monoidEffectFn1 = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.RecordGet(dictMonoid_0_loop, "mempty")
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0_loop, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupEffectFn11_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(f1_3 gopurs_runtime.Value, f2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_6_3 := gopurs_runtime.UncurriedApp(f1_3, a_5)
_ = a_prime_6_3
a_prime_7_4 := gopurs_runtime.UncurriedApp(f2_4, a_5)
_ = a_prime_7_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), a_prime_6_3, a_prime_7_4)
})
})
}))
_ = semigroupEffectFn11_3_2
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
})
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEffectFn11_3_2
}))
}()
})
	})
	return monoidEffectFn1
}



func Get_mkEffectFn1() gopurs_runtime.Value {
	return _Gopurs_MkEffectFn1
}

func Get_mkEffectFn10() gopurs_runtime.Value {
	return _Gopurs_MkEffectFn10
}

func Get_mkEffectFn2() gopurs_runtime.Value {
	return _Gopurs_MkEffectFn2
}

func Get_mkEffectFn3() gopurs_runtime.Value {
	return _Gopurs_MkEffectFn3
}

func Get_mkEffectFn4() gopurs_runtime.Value {
	return _Gopurs_MkEffectFn4
}

func Get_mkEffectFn5() gopurs_runtime.Value {
	return _Gopurs_MkEffectFn5
}

func Get_mkEffectFn6() gopurs_runtime.Value {
	return _Gopurs_MkEffectFn6
}

func Get_mkEffectFn7() gopurs_runtime.Value {
	return _Gopurs_MkEffectFn7
}

func Get_mkEffectFn8() gopurs_runtime.Value {
	return _Gopurs_MkEffectFn8
}

func Get_mkEffectFn9() gopurs_runtime.Value {
	return _Gopurs_MkEffectFn9
}

func Get_runEffectFn1() gopurs_runtime.Value {
	return _Gopurs_RunEffectFn1
}

func Get_runEffectFn10() gopurs_runtime.Value {
	return _Gopurs_RunEffectFn10
}

func Get_runEffectFn2() gopurs_runtime.Value {
	return _Gopurs_RunEffectFn2
}

func Get_runEffectFn3() gopurs_runtime.Value {
	return _Gopurs_RunEffectFn3
}

func Get_runEffectFn4() gopurs_runtime.Value {
	return _Gopurs_RunEffectFn4
}

func Get_runEffectFn5() gopurs_runtime.Value {
	return _Gopurs_RunEffectFn5
}

func Get_runEffectFn6() gopurs_runtime.Value {
	return _Gopurs_RunEffectFn6
}

func Get_runEffectFn7() gopurs_runtime.Value {
	return _Gopurs_RunEffectFn7
}

func Get_runEffectFn8() gopurs_runtime.Value {
	return _Gopurs_RunEffectFn8
}

func Get_runEffectFn9() gopurs_runtime.Value {
	return _Gopurs_RunEffectFn9
}
