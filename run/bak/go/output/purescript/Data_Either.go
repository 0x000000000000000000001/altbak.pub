package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Either_Left gopurs_runtime.Value
var once_Data_Either_Left sync.Once
func Get_Data_Either_Left() gopurs_runtime.Value {
	once_Data_Either_Left.Do(func() {
		cache_Data_Either_Left = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_Data_Either_Left
}

var cache_Data_Either_Right gopurs_runtime.Value
var once_Data_Either_Right sync.Once
func Get_Data_Either_Right() gopurs_runtime.Value {
	once_Data_Either_Right.Do(func() {
		cache_Data_Either_Right = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_Data_Either_Right
}

var cache_Data_Either_showEither gopurs_runtime.Value
var once_Data_Either_showEither sync.Once
func Get_Data_Either_showEither() gopurs_runtime.Value {
	once_Data_Either_showEither.Do(func() {
		cache_Data_Either_showEither = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_showEither(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_Either_showEither
}

var cache_Data_Either_note_prime gopurs_runtime.Value
var once_Data_Either_note_prime sync.Once
func Get_Data_Either_note_prime() gopurs_runtime.Value {
	once_Data_Either_note_prime.Do(func() {
		cache_Data_Either_note_prime = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_note_prime(f_0_box)
})
	})
	return cache_Data_Either_note_prime
}

var cache_Data_Either_note gopurs_runtime.Value
var once_Data_Either_note sync.Once
func Get_Data_Either_note() gopurs_runtime.Value {
	once_Data_Either_note.Do(func() {
		cache_Data_Either_note = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_note(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_Data_Either_note
}

var cache_Data_Either_genericEither gopurs_runtime.Value
var once_Data_Either_genericEither sync.Once
func Get_Data_Either_genericEither() gopurs_runtime.Value {
	once_Data_Either_genericEither.Do(func() {
		cache_Data_Either_genericEither = gopurs_runtime.RecordDict2("from", "to", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0})}
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
if (x_0.Type == 9 && x_0.IntVal == 3478632216) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 492034566) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0})}
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
	return cache_Data_Either_genericEither
}

var cache_Data_Either_functorEither gopurs_runtime.Value
var once_Data_Either_functorEither sync.Once
func Get_Data_Either_functorEither() gopurs_runtime.Value {
	once_Data_Either_functorEither.Do(func() {
		cache_Data_Either_functorEither = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_Data_Either_functorEither
}

var cache_Data_Either_invariantEither gopurs_runtime.Value
var once_Data_Either_invariantEither sync.Once
func Get_Data_Either_invariantEither() gopurs_runtime.Value {
	once_Data_Either_invariantEither.Do(func() {
		cache_Data_Either_invariantEither = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_0)
})
}))
	})
	return cache_Data_Either_invariantEither
}

var cache_Data_Either_fromRight_prime gopurs_runtime.Value
var once_Data_Either_fromRight_prime sync.Once
func Get_Data_Either_fromRight_prime() gopurs_runtime.Value {
	once_Data_Either_fromRight_prime.Do(func() {
		cache_Data_Either_fromRight_prime = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_fromRight_prime(v_0_box, v1_1_box)
})
	})
	return cache_Data_Either_fromRight_prime
}

var cache_Data_Either_fromRight gopurs_runtime.Value
var once_Data_Either_fromRight sync.Once
func Get_Data_Either_fromRight() gopurs_runtime.Value {
	once_Data_Either_fromRight.Do(func() {
		cache_Data_Either_fromRight = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_fromRight(v_0_box, v1_1_box)
})
	})
	return cache_Data_Either_fromRight
}

var cache_Data_Either_fromLeft_prime gopurs_runtime.Value
var once_Data_Either_fromLeft_prime sync.Once
func Get_Data_Either_fromLeft_prime() gopurs_runtime.Value {
	once_Data_Either_fromLeft_prime.Do(func() {
		cache_Data_Either_fromLeft_prime = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_fromLeft_prime(v_0_box, v1_1_box)
})
	})
	return cache_Data_Either_fromLeft_prime
}

var cache_Data_Either_fromLeft gopurs_runtime.Value
var once_Data_Either_fromLeft sync.Once
func Get_Data_Either_fromLeft() gopurs_runtime.Value {
	once_Data_Either_fromLeft.Do(func() {
		cache_Data_Either_fromLeft = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_fromLeft(v_0_box, v1_1_box)
})
	})
	return cache_Data_Either_fromLeft
}

var cache_Data_Either_extendEither gopurs_runtime.Value
var once_Data_Either_extendEither sync.Once
func Get_Data_Either_extendEither() gopurs_runtime.Value {
	once_Data_Either_extendEither.Do(func() {
		cache_Data_Either_extendEither = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_functorEither()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1)})}
}
end_branch_0:
return __t0
})
}))
	})
	return cache_Data_Either_extendEither
}

var cache_Data_Either_eqEither gopurs_runtime.Value
var once_Data_Either_eqEither sync.Once
func Get_Data_Either_eqEither() gopurs_runtime.Value {
	once_Data_Either_eqEither.Do(func() {
		cache_Data_Either_eqEither = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_eqEither(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_Data_Either_eqEither
}

var cache_Data_Either_ordEither gopurs_runtime.Value
var once_Data_Either_ordEither sync.Once
func Get_Data_Either_ordEither() gopurs_runtime.Value {
	once_Data_Either_ordEither.Do(func() {
		cache_Data_Either_ordEither = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_ordEither(dictOrd_0_box)
})
	})
	return cache_Data_Either_ordEither
}

var cache_Data_Either_eq1Either gopurs_runtime.Value
var once_Data_Either_eq1Either sync.Once
func Get_Data_Either_eq1Either() gopurs_runtime.Value {
	once_Data_Either_eq1Either.Do(func() {
		cache_Data_Either_eq1Either = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_eq1Either(dictEq_0_box)
})
	})
	return cache_Data_Either_eq1Either
}

var cache_Data_Either_ord1Either gopurs_runtime.Value
var once_Data_Either_ord1Either sync.Once
func Get_Data_Either_ord1Either() gopurs_runtime.Value {
	once_Data_Either_ord1Either.Do(func() {
		cache_Data_Either_ord1Either = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_ord1Either(dictOrd_0_box)
})
	})
	return cache_Data_Either_ord1Either
}

var cache_Data_Either_either gopurs_runtime.Value
var once_Data_Either_either sync.Once
func Get_Data_Either_either() gopurs_runtime.Value {
	once_Data_Either_either.Do(func() {
		cache_Data_Either_either = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_either(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Either_either
}

var cache_Data_Either_hush gopurs_runtime.Value
var once_Data_Either_hush sync.Once
func Get_Data_Either_hush() gopurs_runtime.Value {
	once_Data_Either_hush.Do(func() {
		cache_Data_Either_hush = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Either_hush(v2_0_box))}
})
	})
	return cache_Data_Either_hush
}

var cache_Data_Either_isLeft gopurs_runtime.Value
var once_Data_Either_isLeft sync.Once
func Get_Data_Either_isLeft() gopurs_runtime.Value {
	once_Data_Either_isLeft.Do(func() {
		cache_Data_Either_isLeft = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Either_isLeft(v2_0_box))
})
	})
	return cache_Data_Either_isLeft
}

var cache_Data_Either_isRight gopurs_runtime.Value
var once_Data_Either_isRight sync.Once
func Get_Data_Either_isRight() gopurs_runtime.Value {
	once_Data_Either_isRight.Do(func() {
		cache_Data_Either_isRight = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Either_isRight(v2_0_box))
})
	})
	return cache_Data_Either_isRight
}

var cache_Data_Either_choose gopurs_runtime.Value
var once_Data_Either_choose sync.Once
func Get_Data_Either_choose() gopurs_runtime.Value {
	once_Data_Either_choose.Do(func() {
		cache_Data_Either_choose = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_choose(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](dictAlt_0_box))
})
	})
	return cache_Data_Either_choose
}

var cache_Data_Either_boundedEither gopurs_runtime.Value
var once_Data_Either_boundedEither sync.Once
func Get_Data_Either_boundedEither() gopurs_runtime.Value {
	once_Data_Either_boundedEither.Do(func() {
		cache_Data_Either_boundedEither = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_boundedEither(dictBounded_0_box)
})
	})
	return cache_Data_Either_boundedEither
}

var cache_Data_Either_blush gopurs_runtime.Value
var once_Data_Either_blush sync.Once
func Get_Data_Either_blush() gopurs_runtime.Value {
	once_Data_Either_blush.Do(func() {
		cache_Data_Either_blush = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Either_blush(v2_0_box))}
})
	})
	return cache_Data_Either_blush
}

var cache_Data_Either_applyEither gopurs_runtime.Value
var once_Data_Either_applyEither sync.Once
func Get_Data_Either_applyEither() gopurs_runtime.Value {
	once_Data_Either_applyEither.Do(func() {
		cache_Data_Either_applyEither = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_functorEither()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_Data_Either_applyEither
}

var cache_Data_Either_bindEither gopurs_runtime.Value
var once_Data_Either_bindEither sync.Once
func Get_Data_Either_bindEither() gopurs_runtime.Value {
	once_Data_Either_bindEither.Do(func() {
		cache_Data_Either_bindEither = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_applyEither()
}), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
_ = __local_var_1_0
__t2 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_1_0})}
})
goto end_branch_2
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
_ = __local_var_1_1
__t2 = gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, __local_var_1_1)
})
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
	})
	return cache_Data_Either_bindEither
}

var cache_Data_Either_semigroupEither gopurs_runtime.Value
var once_Data_Either_semigroupEither sync.Once
func Get_Data_Either_semigroupEither() gopurs_runtime.Value {
	once_Data_Either_semigroupEither.Do(func() {
		cache_Data_Either_semigroupEither = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_semigroupEither(dictSemigroup_0_box)
})
	})
	return cache_Data_Either_semigroupEither
}

var cache_Data_Either_applicativeEither gopurs_runtime.Value
var once_Data_Either_applicativeEither sync.Once
func Get_Data_Either_applicativeEither() gopurs_runtime.Value {
	once_Data_Either_applicativeEither.Do(func() {
		cache_Data_Either_applicativeEither = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_applyEither()
}), Get_Data_Either_Right())
	})
	return cache_Data_Either_applicativeEither
}

var cache_Data_Either_monadEither gopurs_runtime.Value
var once_Data_Either_monadEither sync.Once
func Get_Data_Either_monadEither() gopurs_runtime.Value {
	once_Data_Either_monadEither.Do(func() {
		cache_Data_Either_monadEither = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_applicativeEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_bindEither()
}))
	})
	return cache_Data_Either_monadEither
}

var cache_Data_Either_altEither gopurs_runtime.Value
var once_Data_Either_altEither sync.Once
func Get_Data_Either_altEither() gopurs_runtime.Value {
	once_Data_Either_altEither.Do(func() {
		cache_Data_Either_altEither = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_functorEither()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
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
})
}))
	})
	return cache_Data_Either_altEither
}

var cache_Data_Either_applicativeEither__4081990212 gopurs_runtime.Value
var once_Data_Either_applicativeEither__4081990212 sync.Once
func Get_Data_Either_applicativeEither__4081990212() gopurs_runtime.Value {
	once_Data_Either_applicativeEither__4081990212.Do(func() {
		cache_Data_Either_applicativeEither__4081990212 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_applyEither()
}), Get_Data_Either_Right())
	})
	return cache_Data_Either_applicativeEither__4081990212
}

var cache_Data_Either_applicativeEither__2440223464 gopurs_runtime.Value
var once_Data_Either_applicativeEither__2440223464 sync.Once
func Get_Data_Either_applicativeEither__2440223464() gopurs_runtime.Value {
	once_Data_Either_applicativeEither__2440223464.Do(func() {
		cache_Data_Either_applicativeEither__2440223464 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_applyEither()
}), Get_Data_Either_Right())
	})
	return cache_Data_Either_applicativeEither__2440223464
}

var cache_Data_Either_applyEither__2246489028 gopurs_runtime.Value
var once_Data_Either_applyEither__2246489028 sync.Once
func Get_Data_Either_applyEither__2246489028() gopurs_runtime.Value {
	once_Data_Either_applyEither__2246489028.Do(func() {
		cache_Data_Either_applyEither__2246489028 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_functorEither()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_Data_Either_applyEither__2246489028
}

var cache_Data_Either_applyEither__3806012498 gopurs_runtime.Value
var once_Data_Either_applyEither__3806012498 sync.Once
func Get_Data_Either_applyEither__3806012498() gopurs_runtime.Value {
	once_Data_Either_applyEither__3806012498.Do(func() {
		cache_Data_Either_applyEither__3806012498 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_functorEither()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_Data_Either_applyEither__3806012498
}

var cache_Data_Either_bindEither__3337174823 gopurs_runtime.Value
var once_Data_Either_bindEither__3337174823 sync.Once
func Get_Data_Either_bindEither__3337174823() gopurs_runtime.Value {
	once_Data_Either_bindEither__3337174823.Do(func() {
		cache_Data_Either_bindEither__3337174823 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_applyEither()
}), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
_ = __local_var_1_0
__t2 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_1_0})}
})
goto end_branch_2
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
_ = __local_var_1_1
__t2 = gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, __local_var_1_1)
})
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
	})
	return cache_Data_Either_bindEither__3337174823
}

var cache_Data_Either_either__1999147371 gopurs_runtime.Value
var once_Data_Either_either__1999147371 sync.Once
func Get_Data_Either_either__1999147371() gopurs_runtime.Value {
	once_Data_Either_either__1999147371.Do(func() {
		cache_Data_Either_either__1999147371 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_either__1999147371(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Either_either__1999147371
}

var cache_Data_Either_either__1539695579 gopurs_runtime.Value
var once_Data_Either_either__1539695579 sync.Once
func Get_Data_Either_either__1539695579() gopurs_runtime.Value {
	once_Data_Either_either__1539695579.Do(func() {
		cache_Data_Either_either__1539695579 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Either_either__1539695579(v_0_box, v1_1_box, v2_2_box))
})
	})
	return cache_Data_Either_either__1539695579
}

var cache_Data_Either_either__2158544585 gopurs_runtime.Value
var once_Data_Either_either__2158544585 sync.Once
func Get_Data_Either_either__2158544585() gopurs_runtime.Value {
	once_Data_Either_either__2158544585.Do(func() {
		cache_Data_Either_either__2158544585 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_either__2158544585(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Either_either__2158544585
}

var cache_Data_Either_either__271265665 gopurs_runtime.Value
var once_Data_Either_either__271265665 sync.Once
func Get_Data_Either_either__271265665() gopurs_runtime.Value {
	once_Data_Either_either__271265665.Do(func() {
		cache_Data_Either_either__271265665 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_either__271265665(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Either_either__271265665
}

var cache_Data_Either_either__3836941544 gopurs_runtime.Value
var once_Data_Either_either__3836941544 sync.Once
func Get_Data_Either_either__3836941544() gopurs_runtime.Value {
	once_Data_Either_either__3836941544.Do(func() {
		cache_Data_Either_either__3836941544 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Either_either__3836941544(v_0_box, v1_1_box, v2_2_box))}
})
	})
	return cache_Data_Either_either__3836941544
}

var cache_Data_Either_either__3251677286 gopurs_runtime.Value
var once_Data_Either_either__3251677286 sync.Once
func Get_Data_Either_either__3251677286() gopurs_runtime.Value {
	once_Data_Either_either__3251677286.Do(func() {
		cache_Data_Either_either__3251677286 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_either__3251677286(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Either_either__3251677286
}

var cache_Data_Either_either__2191010510 gopurs_runtime.Value
var once_Data_Either_either__2191010510 sync.Once
func Get_Data_Either_either__2191010510() gopurs_runtime.Value {
	once_Data_Either_either__2191010510.Do(func() {
		cache_Data_Either_either__2191010510 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_either__2191010510(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Either_either__2191010510
}

var cache_Data_Either_either__671458049 gopurs_runtime.Value
var once_Data_Either_either__671458049 sync.Once
func Get_Data_Either_either__671458049() gopurs_runtime.Value {
	once_Data_Either_either__671458049.Do(func() {
		cache_Data_Either_either__671458049 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_either__671458049(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Either_either__671458049
}

var cache_Data_Either_functorEither__13820179 gopurs_runtime.Value
var once_Data_Either_functorEither__13820179 sync.Once
func Get_Data_Either_functorEither__13820179() gopurs_runtime.Value {
	once_Data_Either_functorEither__13820179.Do(func() {
		cache_Data_Either_functorEither__13820179 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_Data_Either_functorEither__13820179
}

var cache_Data_Either_functorEither__1771778897 gopurs_runtime.Value
var once_Data_Either_functorEither__1771778897 sync.Once
func Get_Data_Either_functorEither__1771778897() gopurs_runtime.Value {
	once_Data_Either_functorEither__1771778897.Do(func() {
		cache_Data_Either_functorEither__1771778897 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_Data_Either_functorEither__1771778897
}

var cache_Data_Either_monadEither__2975460307 gopurs_runtime.Value
var once_Data_Either_monadEither__2975460307 sync.Once
func Get_Data_Either_monadEither__2975460307() gopurs_runtime.Value {
	once_Data_Either_monadEither__2975460307.Do(func() {
		cache_Data_Either_monadEither__2975460307 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_applicativeEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_bindEither()
}))
	})
	return cache_Data_Either_monadEither__2975460307
}

type Constructor_Data_Either_Left[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
}


type Constructor_Data_Either_Right[T_a any, T_b any] struct {
	Rc uint32
	V0 T_b
}


func Call_Data_Either_showEither(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Str((("(Left ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (")"))
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Str((("(Right ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (")"))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0.StrVal())
}))
}

func Call_Data_Either_note_prime(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply2(Get_Data_Maybe_maybe_prime(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, x_1)})}
}), Get_Data_Either_Right())
}

func Call_Data_Either_note(a_0_loop gopurs_runtime.Value, v2_1_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0})}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0})}
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

func Call_Data_Either_fromRight_prime(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 2465973597) {
__t0 = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
}
end_branch_0:
return __t0
}

func Call_Data_Either_fromRight(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 2465973597) {
__t0 = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = v_0
}
end_branch_0:
return __t0
}

func Call_Data_Either_fromLeft_prime(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
}
end_branch_0:
return __t0
}

func Call_Data_Either_fromLeft(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = v_0
}
end_branch_0:
return __t0
}

func Call_Data_Either_eqEither(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (x_2.Type == 9 && x_2.IntVal == 3711209382) {
var __t0 bool
{
if (y_3.Type == 9 && y_3.IntVal == 3711209382) {
__t0 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if ((x_2.Type == 9 && x_2.IntVal == 2465973597)) && ((y_3.Type == 9 && y_3.IntVal == 2465973597)) {
__t1 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
}))
}

func Call_Data_Either_ordEither(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): eqEither2_3_1 -> gopurs_runtime.Value
eqEither2_3_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 bool
{
if (x_4.Type == 9 && x_4.IntVal == 3711209382) {
var __t3 bool
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t3 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2465973597)) && ((y_5.Type == 9 && y_5.IntVal == 2465973597)) {
__t4 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
})
}))
_ = eqEither2_3_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqEither2_3_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (x_4.Type == 9 && x_4.IntVal == 3711209382) {
var __t5 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t5 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2465973597)) && ((y_5.Type == 9 && y_5.IntVal == 2465973597)) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t6.IntVal)), UnsafePtr: nil}
})
}))
})
}

func Call_Data_Either_eq1Either(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (x_2.Type == 9 && x_2.IntVal == 3711209382) {
var __t0 bool
{
if (y_3.Type == 9 && y_3.IntVal == 3711209382) {
__t0 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if ((x_2.Type == 9 && x_2.IntVal == 2465973597)) && ((y_3.Type == 9 && y_3.IntVal == 2465973597)) {
__t1 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
})
}))
}

func Call_Data_Either_ord1Either(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): ordEither1_1_0 -> gopurs_runtime.Value
ordEither1_1_0 := Call_Data_Either_ordEither(dictOrd_0)
_ = ordEither1_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): eq1Either1_2_1 -> gopurs_runtime.Value
eq1Either1_2_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 bool
{
if (x_4.Type == 9 && x_4.IntVal == 3711209382) {
var __t3 bool
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t3 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "eq"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2465973597)) && ((y_5.Type == 9 && y_5.IntVal == 2465973597)) {
__t4 = (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
})
})
}))
_ = eq1Either1_2_1
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Either1_2_1
}), gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordEither1_1_0, dictOrd1_3), "compare")
}))
}

func Call_Data_Either_either(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Either_hush(v2_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0)
}

func Call_Data_Either_isLeft(v2_0_loop gopurs_runtime.Value) bool {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_Data_Either_isRight(v2_0_loop gopurs_runtime.Value) bool {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_Data_Either_choose(dictAlt_0_loop *Constructor_Control_Alt_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlt_0 *Constructor_Control_Alt_Alt[gopurs_runtime.Value] = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlt_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictAlt_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Either_Left(), a_2), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Either_Right(), b_3))
})
})
}

func Call_Data_Either_boundedEither(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): ordEither1_1_0 -> gopurs_runtime.Value
ordEither1_1_0 := Call_Data_Either_ordEither(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{}))
_ = ordEither1_1_0
return gopurs_runtime.Func(func(dictBounded1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordEither2_3_1 -> gopurs_runtime.Value
ordEither2_3_1 := gopurs_runtime.Apply(ordEither1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded1_2, "Ord0"), gopurs_runtime.Value{}))
_ = ordEither2_3_1
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ordEither2_3_1
}), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictBounded_0, "bottom")})}, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictBounded1_2, "top")})})
})
}

func Call_Data_Either_blush(v2_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0)
}

func Call_Data_Either_semigroupEither(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): append_1_0 -> gopurs_runtime.Value
append_1_0 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = append_1_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_applyEither(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), append_1_0, x_2), y_3)
})
}))
}

func Call_Data_Either_either__1999147371(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, gopurs_runtime.Str((*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0.StrVal()))
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Either_either__1539695579(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) bool {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Bool((gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Bool((gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_Data_Either_either__2158544585(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Either_either__271265665(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Either_either__3836941544(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0)
}

func Call_Data_Either_either__3251677286(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Either_either__2191010510(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Either_either__671458049(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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


