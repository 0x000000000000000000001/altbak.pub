package Data_Either

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Left gopurs_runtime.Value
var once_Left sync.Once
func Get_Left() gopurs_runtime.Value {
	once_Left.Do(func() {
		cache_Left = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_Left
}

var cache_Right gopurs_runtime.Value
var once_Right sync.Once
func Get_Right() gopurs_runtime.Value {
	once_Right.Do(func() {
		cache_Right = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_Right
}

var cache_showEither gopurs_runtime.Value
var once_showEither sync.Once
func Get_showEither() gopurs_runtime.Value {
	once_showEither.Do(func() {
		cache_showEither = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showEither(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showEither
}

var cache_note_prime gopurs_runtime.Value
var once_note_prime sync.Once
func Get_note_prime() gopurs_runtime.Value {
	once_note_prime.Do(func() {
		cache_note_prime = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_note_prime(f_0_box)
})
	})
	return cache_note_prime
}

var cache_note gopurs_runtime.Value
var once_note sync.Once
func Get_note() gopurs_runtime.Value {
	once_note.Do(func() {
		cache_note = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_note(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_note
}

var cache_genericEither gopurs_runtime.Value
var once_genericEither sync.Once
func Get_genericEither() gopurs_runtime.Value {
	once_genericEither.Do(func() {
		cache_genericEither = gopurs_runtime.RecordDict2("from", "to", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0})}
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 492034566) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0})}
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
	return cache_genericEither
}

var cache_functorEither gopurs_runtime.Value
var once_functorEither sync.Once
func Get_functorEither() gopurs_runtime.Value {
	once_functorEither.Do(func() {
		cache_functorEither = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
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
	return cache_functorEither
}

var cache_functorEither__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__13820179 gopurs_runtime.Value
var once_functorEither__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__13820179 sync.Once
func Get_functorEither__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__13820179() gopurs_runtime.Value {
	once_functorEither__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__13820179.Do(func() {
		cache_functorEither__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__13820179 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
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
})})}
	})
	return cache_functorEither__ptrData_Functor_Constructor_Functor_gopurs_runtime_Value__13820179
}

var cache_functorEither__gopurs_runtime_Value_1771778897 gopurs_runtime.Value
var once_functorEither__gopurs_runtime_Value_1771778897 sync.Once
func Get_functorEither__gopurs_runtime_Value_1771778897() gopurs_runtime.Value {
	once_functorEither__gopurs_runtime_Value_1771778897.Do(func() {
		cache_functorEither__gopurs_runtime_Value_1771778897 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
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
	return cache_functorEither__gopurs_runtime_Value_1771778897
}

var cache_invariantEither gopurs_runtime.Value
var once_invariantEither sync.Once
func Get_invariantEither() gopurs_runtime.Value {
	once_invariantEither.Do(func() {
		cache_invariantEither = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorEither(), "map"), f_0)
})
}))
	})
	return cache_invariantEither
}

var cache_fromRight_prime gopurs_runtime.Value
var once_fromRight_prime sync.Once
func Get_fromRight_prime() gopurs_runtime.Value {
	once_fromRight_prime.Do(func() {
		cache_fromRight_prime = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromRight_prime(v_0_box, v1_1_box)
})
	})
	return cache_fromRight_prime
}

var cache_fromRight gopurs_runtime.Value
var once_fromRight sync.Once
func Get_fromRight() gopurs_runtime.Value {
	once_fromRight.Do(func() {
		cache_fromRight = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromRight(v_0_box, v1_1_box)
})
	})
	return cache_fromRight
}

var cache_fromLeft_prime gopurs_runtime.Value
var once_fromLeft_prime sync.Once
func Get_fromLeft_prime() gopurs_runtime.Value {
	once_fromLeft_prime.Do(func() {
		cache_fromLeft_prime = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromLeft_prime(v_0_box, v1_1_box)
})
	})
	return cache_fromLeft_prime
}

var cache_fromLeft gopurs_runtime.Value
var once_fromLeft sync.Once
func Get_fromLeft() gopurs_runtime.Value {
	once_fromLeft.Do(func() {
		cache_fromLeft = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromLeft(v_0_box, v1_1_box)
})
	})
	return cache_fromLeft
}

var cache_extendEither gopurs_runtime.Value
var once_extendEither sync.Once
func Get_extendEither() gopurs_runtime.Value {
	once_extendEither.Do(func() {
		cache_extendEither = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, v1_1)})}
}
end_branch_0:
return __t0
})
}))
	})
	return cache_extendEither
}

var cache_eqEither gopurs_runtime.Value
var once_eqEither sync.Once
func Get_eqEither() gopurs_runtime.Value {
	once_eqEither.Do(func() {
		cache_eqEither = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, dictEq1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqEither(dictEq_0_box, dictEq1_1_box)
})
	})
	return cache_eqEither
}

var cache_ordEither gopurs_runtime.Value
var once_ordEither sync.Once
func Get_ordEither() gopurs_runtime.Value {
	once_ordEither.Do(func() {
		cache_ordEither = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordEither(dictOrd_0_box)
})
	})
	return cache_ordEither
}

var cache_eq1Either gopurs_runtime.Value
var once_eq1Either sync.Once
func Get_eq1Either() gopurs_runtime.Value {
	once_eq1Either.Do(func() {
		cache_eq1Either = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1Either(dictEq_0_box)
})
	})
	return cache_eq1Either
}

var cache_ord1Either gopurs_runtime.Value
var once_ord1Either sync.Once
func Get_ord1Either() gopurs_runtime.Value {
	once_ord1Either.Do(func() {
		cache_ord1Either = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1Either(dictOrd_0_box)
})
	})
	return cache_ord1Either
}

var cache_either gopurs_runtime.Value
var once_either sync.Once
func Get_either() gopurs_runtime.Value {
	once_either.Do(func() {
		cache_either = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_either
}

var cache_either__gopurs_runtime_Value_2158544585 gopurs_runtime.Value
var once_either__gopurs_runtime_Value_2158544585 sync.Once
func Get_either__gopurs_runtime_Value_2158544585() gopurs_runtime.Value {
	once_either__gopurs_runtime_Value_2158544585.Do(func() {
		cache_either__gopurs_runtime_Value_2158544585 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_either__gopurs_runtime_Value_2158544585(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_either__gopurs_runtime_Value_2158544585
}

var cache_hush gopurs_runtime.Value
var once_hush sync.Once
func Get_hush() gopurs_runtime.Value {
	once_hush.Do(func() {
		cache_hush = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_hush(v2_0_box))}
})
	})
	return cache_hush
}

var cache_isLeft gopurs_runtime.Value
var once_isLeft sync.Once
func Get_isLeft() gopurs_runtime.Value {
	once_isLeft.Do(func() {
		cache_isLeft = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isLeft(v2_0_box))
})
	})
	return cache_isLeft
}

var cache_isRight gopurs_runtime.Value
var once_isRight sync.Once
func Get_isRight() gopurs_runtime.Value {
	once_isRight.Do(func() {
		cache_isRight = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isRight(v2_0_box))
})
	})
	return cache_isRight
}

var cache_choose gopurs_runtime.Value
var once_choose sync.Once
func Get_choose() gopurs_runtime.Value {
	once_choose.Do(func() {
		cache_choose = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_choose(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dictAlt_0_box))
})
	})
	return cache_choose
}

var cache_boundedEither gopurs_runtime.Value
var once_boundedEither sync.Once
func Get_boundedEither() gopurs_runtime.Value {
	once_boundedEither.Do(func() {
		cache_boundedEither = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedEither(dictBounded_0_box)
})
	})
	return cache_boundedEither
}

var cache_blush gopurs_runtime.Value
var once_blush sync.Once
func Get_blush() gopurs_runtime.Value {
	once_blush.Do(func() {
		cache_blush = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_blush(v2_0_box))}
})
	})
	return cache_blush
}

var cache_applyEither gopurs_runtime.Value
var once_applyEither sync.Once
func Get_applyEither() gopurs_runtime.Value {
	once_applyEither.Do(func() {
		cache_applyEither = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorEither(), "map"), (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)
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
	return cache_applyEither
}

var cache_applyEither__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__2246489028 gopurs_runtime.Value
var once_applyEither__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__2246489028 sync.Once
func Get_applyEither__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__2246489028() gopurs_runtime.Value {
	once_applyEither__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__2246489028.Do(func() {
		cache_applyEither__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__2246489028 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorEither(), "map"), (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)
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
})})}
	})
	return cache_applyEither__ptrControl_Apply_Constructor_Apply_gopurs_runtime_Value__2246489028
}

var cache_applyEither__gopurs_runtime_Value_3806012498 gopurs_runtime.Value
var once_applyEither__gopurs_runtime_Value_3806012498 sync.Once
func Get_applyEither__gopurs_runtime_Value_3806012498() gopurs_runtime.Value {
	once_applyEither__gopurs_runtime_Value_3806012498.Do(func() {
		cache_applyEither__gopurs_runtime_Value_3806012498 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorEither(), "map"), (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)
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
	return cache_applyEither__gopurs_runtime_Value_3806012498
}

var cache_bindEither gopurs_runtime.Value
var once_bindEither sync.Once
func Get_bindEither() gopurs_runtime.Value {
	once_bindEither.Do(func() {
		cache_bindEither = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEither()
}), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__local_var_1_1 := (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
_ = __local_var_1_1
__t0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_1_1})}
})
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__local_var_1_2 := (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
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
}))
	})
	return cache_bindEither
}

var cache_bindEither__gopurs_runtime_Value_3337174823 gopurs_runtime.Value
var once_bindEither__gopurs_runtime_Value_3337174823 sync.Once
func Get_bindEither__gopurs_runtime_Value_3337174823() gopurs_runtime.Value {
	once_bindEither__gopurs_runtime_Value_3337174823.Do(func() {
		cache_bindEither__gopurs_runtime_Value_3337174823 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEither()
}), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__local_var_1_1 := (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
_ = __local_var_1_1
__t0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_1_1})}
})
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__local_var_1_2 := (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
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
}))
	})
	return cache_bindEither__gopurs_runtime_Value_3337174823
}

var cache_semigroupEither gopurs_runtime.Value
var once_semigroupEither sync.Once
func Get_semigroupEither() gopurs_runtime.Value {
	once_semigroupEither.Do(func() {
		cache_semigroupEither = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupEither(dictSemigroup_0_box)
})
	})
	return cache_semigroupEither
}

var cache_applicativeEither gopurs_runtime.Value
var once_applicativeEither sync.Once
func Get_applicativeEither() gopurs_runtime.Value {
	once_applicativeEither.Do(func() {
		cache_applicativeEither = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEither()
}), Get_Right())
	})
	return cache_applicativeEither
}

var cache_applicativeEither__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__4081990212 gopurs_runtime.Value
var once_applicativeEither__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__4081990212 sync.Once
func Get_applicativeEither__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__4081990212() gopurs_runtime.Value {
	once_applicativeEither__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__4081990212.Do(func() {
		cache_applicativeEither__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__4081990212 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEither()
}), Get_Right()})}
	})
	return cache_applicativeEither__ptrControl_Applicative_Constructor_Applicative_gopurs_runtime_Value__4081990212
}

var cache_applicativeEither__gopurs_runtime_Value_2440223464 gopurs_runtime.Value
var once_applicativeEither__gopurs_runtime_Value_2440223464 sync.Once
func Get_applicativeEither__gopurs_runtime_Value_2440223464() gopurs_runtime.Value {
	once_applicativeEither__gopurs_runtime_Value_2440223464.Do(func() {
		cache_applicativeEither__gopurs_runtime_Value_2440223464 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyEither()
}), Get_Right())
	})
	return cache_applicativeEither__gopurs_runtime_Value_2440223464
}

var cache_monadEither gopurs_runtime.Value
var once_monadEither sync.Once
func Get_monadEither() gopurs_runtime.Value {
	once_monadEither.Do(func() {
		cache_monadEither = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindEither()
}))
	})
	return cache_monadEither
}

var cache_monadEither__gopurs_runtime_Value_2975460307 gopurs_runtime.Value
var once_monadEither__gopurs_runtime_Value_2975460307 sync.Once
func Get_monadEither__gopurs_runtime_Value_2975460307() gopurs_runtime.Value {
	once_monadEither__gopurs_runtime_Value_2975460307.Do(func() {
		cache_monadEither__gopurs_runtime_Value_2975460307 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindEither()
}))
	})
	return cache_monadEither__gopurs_runtime_Value_2975460307
}

var cache_altEither gopurs_runtime.Value
var once_altEither sync.Once
func Get_altEither() gopurs_runtime.Value {
	once_altEither.Do(func() {
		cache_altEither = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorEither()
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
	return cache_altEither
}

type Constructor_Left[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
}


type Constructor_Right[T_a any, T_b any] struct {
	Rc uint32
	V0 T_b
}


func Call_showEither(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Left "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Str(")"))).StrVal())
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Right "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Str(")"))).StrVal())
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

func Call_note_prime(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply2(pkg_Data_Maybe.Get_maybe_prime(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, x_1)})}
}), Get_Right())
}

func Call_note(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0})}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0})}
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
if (v1_1.Type == 9 && v1_1.IntVal == 2465973597) {
__t0 = (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
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
if (v1_1.Type == 9 && v1_1.IntVal == 2465973597) {
__t0 = (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0
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

func Call_fromLeft_prime(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
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
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0
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

func Call_eqEither(dictEq_0_loop gopurs_runtime.Value, dictEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var dictEq1_1 gopurs_runtime.Value = dictEq1_1_loop
_ = dictEq1_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_2.Type == 9 && x_2.IntVal == 3711209382) {
var __t1 gopurs_runtime.Value
{
if (y_3.Type == 9 && y_3.IntVal == 3711209382) {
__t1 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if ((x_2.Type == 9 && x_2.IntVal == 2465973597)) && ((y_3.Type == 9 && y_3.IntVal == 2465973597)) {
__t0 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(false)
}
end_branch_0:
return gopurs_runtime.Bool((__t0.IntVal) != (0))
})
}))
}

func Call_ordEither(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_3_2
eqEither2_3_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (x_4.Type == 9 && x_4.IntVal == 3711209382) {
var __t4 gopurs_runtime.Value
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t4 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2465973597)) && ((y_5.Type == 9 && y_5.IntVal == 2465973597)) {
__t3 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "eq"), (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(false)
}
end_branch_3:
return gopurs_runtime.Bool((__t3.IntVal) != (0))
})
}))
_ = eqEither2_3_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqEither2_3_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (x_4.Type == 9 && x_4.IntVal == 3711209382) {
var __t6 gopurs_runtime.Value
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t6 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(nil)}
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(nil)}
goto end_branch_5
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2465973597)) && ((y_5.Type == 9 && y_5.IntVal == 2465973597)) {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
})
}))
})
}

func Call_eq1Either(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_2.Type == 9 && x_2.IntVal == 3711209382) {
var __t1 gopurs_runtime.Value
{
if (y_3.Type == 9 && y_3.IntVal == 3711209382) {
__t1 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if ((x_2.Type == 9 && x_2.IntVal == 2465973597)) && ((y_3.Type == 9 && y_3.IntVal == 2465973597)) {
__t0 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(false)
}
end_branch_0:
return gopurs_runtime.Bool((__t0.IntVal) != (0))
})
})
}))
}

func Call_ord1Either(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
ordEither1_1_0 := Call_ordEither(dictOrd_0)
_ = ordEither1_1_0
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_2
eq1Either1_2_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (x_4.Type == 9 && x_4.IntVal == 3711209382) {
var __t4 gopurs_runtime.Value
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t4 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "eq"), (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2465973597)) && ((y_5.Type == 9 && y_5.IntVal == 2465973597)) {
__t3 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(false)
}
end_branch_3:
return gopurs_runtime.Bool((__t3.IntVal) != (0))
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

func Call_either(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_either__gopurs_runtime_Value_2158544585(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_hush(v2_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_isLeft(v2_0_loop gopurs_runtime.Value) bool {
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

func Call_isRight(v2_0_loop gopurs_runtime.Value) bool {
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

func Call_choose(dictAlt_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlt_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dictAlt_0_loop
_ = dictAlt_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlt_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictAlt_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, Get_Left(), a_2), gopurs_runtime.Apply2(Functor0_1_0.V0, Get_Right(), b_3))
})
})
}

func Call_boundedEither(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
ordEither1_1_0 := Call_ordEither(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{}))
_ = ordEither1_1_0
return gopurs_runtime.Func(func(dictBounded1_2 gopurs_runtime.Value) gopurs_runtime.Value {
ordEither2_3_1 := gopurs_runtime.Apply(ordEither1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded1_2, "Ord0"), gopurs_runtime.Value{}))
_ = ordEither2_3_1
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ordEither2_3_1
}), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictBounded_0, "bottom")})}, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictBounded1_2, "top")})})
})
}

func Call_blush(v2_0_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_semigroupEither(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
append_1_0 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = append_1_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyEither(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorEither(), "map"), append_1_0, x_2), y_3)
})
}))
}


