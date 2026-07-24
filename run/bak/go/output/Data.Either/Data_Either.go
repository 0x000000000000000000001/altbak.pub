package Data_Either

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var Left gopurs_runtime.Value
var once_Left sync.Once
func Get_Left() gopurs_runtime.Value {
	once_Left.Do(func() {
		Left = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Left", value0)
})
	})
	return Left
}

var Right gopurs_runtime.Value
var once_Right sync.Once
func Get_Right() gopurs_runtime.Value {
	once_Right.Do(func() {
		Right = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Right", value0)
})
	})
	return Right
}

var showEither gopurs_runtime.Value
var once_showEither sync.Once
func Get_showEither() gopurs_runtime.Value {
	once_showEither.Do(func() {
		showEither = gopurs_runtime.Func2(Call_showEither)
	})
	return showEither
}

var note_prime gopurs_runtime.Value
var once_note_prime sync.Once
func Get_note_prime() gopurs_runtime.Value {
	once_note_prime.Do(func() {
		note_prime = gopurs_runtime.Func2(Call_note_prime)
	})
	return note_prime
}

var note gopurs_runtime.Value
var once_note sync.Once
func Get_note() gopurs_runtime.Value {
	once_note.Do(func() {
		note = gopurs_runtime.Func2(Call_note)
	})
	return note
}

var genericEither gopurs_runtime.Value
var once_genericEither sync.Once
func Get_genericEither() gopurs_runtime.Value {
	once_genericEither.Do(func() {
		genericEither = gopurs_runtime.RecordDict2("to", "from", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_0.StrVal == "Inl").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(x_0.StrVal == "Inr").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_0.StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Inl", (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(x_0.StrVal == "Right").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Inr", (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
	})
	return genericEither
}

var functorEither gopurs_runtime.Value
var once_functorEither sync.Once
func Get_functorEither() gopurs_runtime.Value {
	once_functorEither.Do(func() {
		functorEither = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(m_1.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(m_1.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_1.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(m_1.UnsafePtr)[0]))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return functorEither
}

var invariantEither gopurs_runtime.Value
var once_invariantEither sync.Once
func Get_invariantEither() gopurs_runtime.Value {
	once_invariantEither.Do(func() {
		invariantEither = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(m_2.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(m_2.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(m_2.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(m_2.UnsafePtr)[0]))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return invariantEither
}

var fromRight_prime gopurs_runtime.Value
var once_fromRight_prime sync.Once
func Get_fromRight_prime() gopurs_runtime.Value {
	once_fromRight_prime.Do(func() {
		fromRight_prime = gopurs_runtime.Func2(Call_fromRight_prime)
	})
	return fromRight_prime
}

var fromRight gopurs_runtime.Value
var once_fromRight sync.Once
func Get_fromRight() gopurs_runtime.Value {
	once_fromRight.Do(func() {
		fromRight = gopurs_runtime.Func2(Call_fromRight)
	})
	return fromRight
}

var fromLeft_prime gopurs_runtime.Value
var once_fromLeft_prime sync.Once
func Get_fromLeft_prime() gopurs_runtime.Value {
	once_fromLeft_prime.Do(func() {
		fromLeft_prime = gopurs_runtime.Func2(Call_fromLeft_prime)
	})
	return fromLeft_prime
}

var fromLeft gopurs_runtime.Value
var once_fromLeft sync.Once
func Get_fromLeft() gopurs_runtime.Value {
	once_fromLeft.Do(func() {
		fromLeft = gopurs_runtime.Func2(Call_fromLeft)
	})
	return fromLeft
}

var extendEither gopurs_runtime.Value
var once_extendEither sync.Once
func Get_extendEither() gopurs_runtime.Value {
	once_extendEither.Do(func() {
		extendEither = gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(v_0, v1_1))
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
}))
	})
	return extendEither
}

var eqEither gopurs_runtime.Value
var once_eqEither sync.Once
func Get_eqEither() gopurs_runtime.Value {
	once_eqEither.Do(func() {
		eqEither = gopurs_runtime.Func2(Call_eqEither)
	})
	return eqEither
}

var ordEither gopurs_runtime.Value
var once_ordEither sync.Once
func Get_ordEither() gopurs_runtime.Value {
	once_ordEither.Do(func() {
		ordEither = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0_loop, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_3_1
eqEither2_4_2 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_4.StrVal == "Left").IntVal != 0 {
__t3 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_5.StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0]).IntVal != 0)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_4.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(y_5.StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "eq"), (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0]).IntVal != 0)
}
end_branch_3:
return __t3
}))
_ = eqEither2_4_2
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_5.StrVal == "Left").IntVal != 0 {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(y_6.StrVal == "Left").IntVal != 0 {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0_loop, "compare"), (*[1024]gopurs_runtime.Value)(x_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0])
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Constructor0("LT")
}
end_branch_5:
__t4 = __t5
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(y_6.StrVal == "Left").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("GT")
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(x_5.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(y_6.StrVal == "Right").IntVal != 0 {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*[1024]gopurs_runtime.Value)(x_5.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_6.UnsafePtr)[0])
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqEither2_4_2
}))
})
}()
})
	})
	return ordEither
}

var eq1Either gopurs_runtime.Value
var once_eq1Either sync.Once
func Get_eq1Either() gopurs_runtime.Value {
	once_eq1Either.Do(func() {
		eq1Either = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_2.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_3.StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0_loop, "eq"), (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0]).IntVal != 0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_2.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(y_3.StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0]).IntVal != 0)
}
end_branch_0:
return __t0
}))
}()
})
	})
	return eq1Either
}

var ord1Either gopurs_runtime.Value
var once_ord1Either sync.Once
func Get_ord1Either() gopurs_runtime.Value {
	once_ord1Either.Do(func() {
		ord1Either = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordEither1_1_0 := gopurs_runtime.Apply(Get_ordEither(), dictOrd_0_loop)
_ = ordEither1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0_loop, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_1
eq1Either1_3_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq1_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_4.StrVal == "Left").IntVal != 0 {
__t3 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_5.StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "eq"), (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0]).IntVal != 0)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_4.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(y_5.StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0]).IntVal != 0)
}
end_branch_3:
return __t3
}))
_ = eq1Either1_3_2
return gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordEither1_1_0, dictOrd1_4), "compare")
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Either1_3_2
}))
}()
})
	})
	return ord1Either
}

var either gopurs_runtime.Value
var once_either sync.Once
func Get_either() gopurs_runtime.Value {
	once_either.Do(func() {
		either = gopurs_runtime.Func3(Call_either)
	})
	return either
}

var hush gopurs_runtime.Value
var once_hush sync.Once
func Get_hush() gopurs_runtime.Value {
	once_hush.Do(func() {
		hush = gopurs_runtime.Func(func(v2_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_0_loop.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_0_loop.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v2_0_loop.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
	})
	return hush
}

var isLeft gopurs_runtime.Value
var once_isLeft sync.Once
func Get_isLeft() gopurs_runtime.Value {
	once_isLeft.Do(func() {
		isLeft = gopurs_runtime.Func(func(v2_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_0_loop.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_0_loop.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
	})
	return isLeft
}

var isRight gopurs_runtime.Value
var once_isRight sync.Once
func Get_isRight() gopurs_runtime.Value {
	once_isRight.Do(func() {
		isRight = gopurs_runtime.Func(func(v2_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_0_loop.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_0_loop.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
	})
	return isRight
}

var choose gopurs_runtime.Value
var once_choose sync.Once
func Get_choose() gopurs_runtime.Value {
	once_choose.Do(func() {
		choose = gopurs_runtime.Func(func(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0_loop, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0_loop, "alt"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), Get_Left(), a_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), Get_Right(), b_3))
})
}()
})
	})
	return choose
}

var boundedEither gopurs_runtime.Value
var once_boundedEither sync.Once
func Get_boundedEither() gopurs_runtime.Value {
	once_boundedEither.Do(func() {
		boundedEither = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
bottom_1_0 := gopurs_runtime.RecordGet(dictBounded_0_loop, "bottom")
_ = bottom_1_0
ordEither1_2_1 := gopurs_runtime.Apply(Get_ordEither(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0_loop, "Ord0"), gopurs_runtime.Value{}))
_ = ordEither1_2_1
return gopurs_runtime.Func(func(dictBounded1_3 gopurs_runtime.Value) gopurs_runtime.Value {
ordEither2_4_2 := gopurs_runtime.Apply(ordEither1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded1_3, "Ord0"), gopurs_runtime.Value{}))
_ = ordEither2_4_2
return gopurs_runtime.RecordDict3("top", "bottom", "Ord0", gopurs_runtime.Constructor1("Right", gopurs_runtime.RecordGet(dictBounded1_3, "top")), gopurs_runtime.Constructor1("Left", bottom_1_0), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return ordEither2_4_2
}))
})
}()
})
	})
	return boundedEither
}

var blush gopurs_runtime.Value
var once_blush sync.Once
func Get_blush() gopurs_runtime.Value {
	once_blush.Do(func() {
		blush = gopurs_runtime.Func(func(v2_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_0_loop.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v2_0_loop.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_0_loop.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
	})
	return blush
}

var applyEither gopurs_runtime.Value
var once_applyEither sync.Once
func Get_applyEither() gopurs_runtime.Value {
	once_applyEither.Do(func() {
		applyEither = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Right").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1.StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_1.StrVal == "Right").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
}))
	})
	return applyEither
}

var bindEither gopurs_runtime.Value
var once_bindEither sync.Once
func Get_bindEither() gopurs_runtime.Value {
	once_bindEither.Do(func() {
		bindEither = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_0.StrVal == "Left").IntVal != 0 {
__local_var_1_1 := (*[1024]gopurs_runtime.Value)(v2_0.UnsafePtr)[0]
_ = __local_var_1_1
__t0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Left", __local_var_1_1)
})
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_0.StrVal == "Right").IntVal != 0 {
__local_var_1_2 := (*[1024]gopurs_runtime.Value)(v2_0.UnsafePtr)[0]
_ = __local_var_1_2
__t0 = gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, __local_var_1_2)
})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEither()
}))
	})
	return bindEither
}

var semigroupEither gopurs_runtime.Value
var once_semigroupEither sync.Once
func Get_semigroupEither() gopurs_runtime.Value {
	once_semigroupEither.Do(func() {
		semigroupEither = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_1.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(x_1.StrVal == "Right").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(y_2.StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(y_2.StrVal == "Right").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0_loop, "append"), (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}()
})
	})
	return semigroupEither
}

var applicativeEither gopurs_runtime.Value
var once_applicativeEither sync.Once
func Get_applicativeEither() gopurs_runtime.Value {
	once_applicativeEither.Do(func() {
		applicativeEither = gopurs_runtime.RecordDict2("pure", "Apply0", Get_Right(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEither()
}))
	})
	return applicativeEither
}

var monadEither gopurs_runtime.Value
var once_monadEither sync.Once
func Get_monadEither() gopurs_runtime.Value {
	once_monadEither.Do(func() {
		monadEither = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindEither()
}))
	})
	return monadEither
}

var altEither gopurs_runtime.Value
var once_altEither sync.Once
func Get_altEither() gopurs_runtime.Value {
	once_altEither.Do(func() {
		altEither = gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Left").IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
__t0 = v_0
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
}))
	})
	return altEither
}

func Call_showEither(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Str("(Left " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]).StrVal + ")")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Str("(Right " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1_loop, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]).StrVal + ")")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}

func Call_note_prime(f_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v2_1 gopurs_runtime.Value = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_1_loop.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", gopurs_runtime.Apply(f_0_loop, pkg_Data_Unit.Get_unit()))
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_1_loop.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)(v2_1_loop.UnsafePtr)[0])
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

func Call_note(a_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 gopurs_runtime.Value = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_1_loop.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", a_0_loop)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_1_loop.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)(v2_1_loop.UnsafePtr)[0])
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

func Call_fromRight_prime(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1_loop.StrVal == "Right").IntVal != 0 {
__t0 = (*[1024]gopurs_runtime.Value)(v1_1_loop.UnsafePtr)[0]
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(v_0_loop, pkg_Data_Unit.Get_unit())
}
end_branch_0:
return __t0
}

func Call_fromRight(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1_loop.StrVal == "Right").IntVal != 0 {
__t0 = (*[1024]gopurs_runtime.Value)(v1_1_loop.UnsafePtr)[0]
goto end_branch_0
} else {

}
}
{
__t0 = v_0_loop
}
end_branch_0:
return __t0
}

func Call_fromLeft_prime(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1_loop.StrVal == "Left").IntVal != 0 {
__t0 = (*[1024]gopurs_runtime.Value)(v1_1_loop.UnsafePtr)[0]
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(v_0_loop, pkg_Data_Unit.Get_unit())
}
end_branch_0:
return __t0
}

func Call_fromLeft(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1_loop.StrVal == "Left").IntVal != 0 {
__t0 = (*[1024]gopurs_runtime.Value)(v1_1_loop.UnsafePtr)[0]
goto end_branch_0
} else {

}
}
{
__t0 = v_0_loop
}
end_branch_0:
return __t0
}

func Call_eqEither(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_2.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(y_3.StrVal == "Left").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0_loop, "eq"), (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0]).IntVal != 0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_2.StrVal == "Right").IntVal != 0 && gopurs_runtime.Bool(y_3.StrVal == "Right").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1_loop, "eq"), (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0]).IntVal != 0)
}
end_branch_0:
return __t0
}))
}

func Call_either(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_2_loop.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply(v_0_loop, (*[1024]gopurs_runtime.Value)(v2_2_loop.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_2_loop.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Apply(v1_1_loop, (*[1024]gopurs_runtime.Value)(v2_2_loop.UnsafePtr)[0])
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


