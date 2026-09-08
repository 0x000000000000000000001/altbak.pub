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
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
})
	})
	return cache_Data_Either_Left
}

var cache_Data_Either_Left__3336285006 gopurs_runtime.Value
var once_Data_Either_Left__3336285006 sync.Once
func Get_Data_Either_Left__3336285006() gopurs_runtime.Value {
	once_Data_Either_Left__3336285006.Do(func() {
		cache_Data_Either_Left__3336285006 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__3336285006(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](__eta_norm_0_0_box))
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
})
	})
	return cache_Data_Either_Left__3336285006
}

var cache_Data_Either_Left__247931836 gopurs_runtime.Value
var once_Data_Either_Left__247931836 sync.Once
func Get_Data_Either_Left__247931836() gopurs_runtime.Value {
	once_Data_Either_Left__247931836.Do(func() {
		cache_Data_Either_Left__247931836 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Left__247931836(__eta_norm_0_0_box)
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
})
	})
	return cache_Data_Either_Left__247931836
}

var cache_Data_Either_Right gopurs_runtime.Value
var once_Data_Either_Right sync.Once
func Get_Data_Either_Right() gopurs_runtime.Value {
	once_Data_Either_Right.Do(func() {
		cache_Data_Either_Right = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0}))}
})
	})
	return cache_Data_Either_Right
}

var cache_Data_Either_Right__3779110556 gopurs_runtime.Value
var once_Data_Either_Right__3779110556 sync.Once
func Get_Data_Either_Right__3779110556() gopurs_runtime.Value {
	once_Data_Either_Right__3779110556.Do(func() {
		cache_Data_Either_Right__3779110556 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Right__3779110556(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](__eta_norm_0_0_box))
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
})
	})
	return cache_Data_Either_Right__3779110556
}

var cache_Data_Either_Right__1333784983 gopurs_runtime.Value
var once_Data_Either_Right__1333784983 sync.Once
func Get_Data_Either_Right__1333784983() gopurs_runtime.Value {
	once_Data_Either_Right__1333784983.Do(func() {
		cache_Data_Either_Right__1333784983 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Either_Right__1333784983(__eta_norm_0_unused_0_box)
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
})
	})
	return cache_Data_Either_Right__1333784983
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

var cache_Data_Either_note_prime_ gopurs_runtime.Value
var once_Data_Either_note_prime_ sync.Once
func Get_Data_Either_note_prime_() gopurs_runtime.Value {
	once_Data_Either_note_prime_.Do(func() {
		cache_Data_Either_note_prime_ = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Either_note_prime_(f_0_box)
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
})
	})
	return cache_Data_Either_note_prime_
}

var cache_Data_Either_note gopurs_runtime.Value
var once_Data_Either_note sync.Once
func Get_Data_Either_note() gopurs_runtime.Value {
	once_Data_Either_note.Do(func() {
		cache_Data_Either_note = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Either_note(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_1_box))
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
})
	})
	return cache_Data_Either_note
}

var cache_Data_Either_genericEither gopurs_runtime.Value
var once_Data_Either_genericEither sync.Once
func Get_Data_Either_genericEither() gopurs_runtime.Value {
	once_Data_Either_genericEither.Do(func() {
		cache_Data_Either_genericEither = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0}))}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer((&Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0}))}
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
__t1 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_1
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 492034566) {
__t1 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})}))}
	})
	return cache_Data_Either_genericEither
}

var cache_Data_Either_functorEither gopurs_runtime.Value
var once_Data_Either_functorEither sync.Once
func Get_Data_Either_functorEither() gopurs_runtime.Value {
	once_Data_Either_functorEither.Do(func() {
		cache_Data_Either_functorEither = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})}))}
	})
	return cache_Data_Either_functorEither
}

var cache_Data_Either_invariantEither gopurs_runtime.Value
var once_Data_Either_invariantEither sync.Once
func Get_Data_Either_invariantEither() gopurs_runtime.Value {
	once_Data_Either_invariantEither.Do(func() {
		cache_Data_Either_invariantEither = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_2.Type == 9 && m_2.IntVal == 3711209382) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
if (m_2.Type == 9 && m_2.IntVal == 2465973597) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})}))}
	})
	return cache_Data_Either_invariantEither
}

var cache_Data_Either_fromRight_prime_ gopurs_runtime.Value
var once_Data_Either_fromRight_prime_ sync.Once
func Get_Data_Either_fromRight_prime_() gopurs_runtime.Value {
	once_Data_Either_fromRight_prime_.Do(func() {
		cache_Data_Either_fromRight_prime_ = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_fromRight_prime_(v_0_box, v1_1_box)
})
	})
	return cache_Data_Either_fromRight_prime_
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

var cache_Data_Either_fromLeft_prime_ gopurs_runtime.Value
var once_Data_Either_fromLeft_prime_ sync.Once
func Get_Data_Either_fromLeft_prime_() gopurs_runtime.Value {
	once_Data_Either_fromLeft_prime_.Do(func() {
		cache_Data_Either_fromLeft_prime_ = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_fromLeft_prime_(v_0_box, v1_1_box)
})
	})
	return cache_Data_Either_fromLeft_prime_
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
		cache_Data_Either_extendEither = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Either_functorEither()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(v_0, v1_1), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
}
end_branch_0:
return __t0
})}))}
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

var cache_Data_Either_either__1841469567 gopurs_runtime.Value
var once_Data_Either_either__1841469567 sync.Once
func Get_Data_Either_either__1841469567() gopurs_runtime.Value {
	once_Data_Either_either__1841469567.Do(func() {
		cache_Data_Either_either__1841469567 = gopurs_runtime.Func3(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Either_either__1841469567(v_unused_0_box, v1_unused_1_box, v2_2_box)
})
	})
	return cache_Data_Either_either__1841469567
}

var cache_Data_Either_hush gopurs_runtime.Value
var once_Data_Either_hush sync.Once
func Get_Data_Either_hush() gopurs_runtime.Value {
	once_Data_Either_hush.Do(func() {
		cache_Data_Either_hush = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Either_hush(v2_0_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
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
return func() gopurs_runtime.Value {
				_v := Call_Data_Either_blush(v2_0_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Either_blush
}

var cache_Data_Either_applyEither gopurs_runtime.Value
var once_Data_Either_applyEither sync.Once
func Get_Data_Either_applyEither() gopurs_runtime.Value {
	once_Data_Either_applyEither.Do(func() {
		cache_Data_Either_applyEither = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Either_functorEither()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 3711209382) {
__t1 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_1
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 2465973597) {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2465973597) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply((*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})}))}
	})
	return cache_Data_Either_applyEither
}

var cache_Data_Either_bindEither gopurs_runtime.Value
var once_Data_Either_bindEither sync.Once
func Get_Data_Either_bindEither() gopurs_runtime.Value {
	once_Data_Either_bindEither.Do(func() {
		cache_Data_Either_bindEither = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Either_applyEither()))}
}), gopurs_runtime.Func(func(v2_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
// TAST (Let): __local_var_1_0 shape=Other bindingType=(TypeVar a)
__local_var_1_0 := (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0
_ = __local_var_1_0
__t2 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{__local_var_1_0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
})
goto end_branch_2
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
// TAST (Let): __local_var_1_1 shape=Other bindingType=(TypeVar b)
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
})}))}
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
		cache_Data_Either_applicativeEither = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Either_applyEither()))}
}), Get_Data_Either_Right()}))}
	})
	return cache_Data_Either_applicativeEither
}

var cache_Data_Either_monadEither gopurs_runtime.Value
var once_Data_Either_monadEither sync.Once
func Get_Data_Either_monadEither() gopurs_runtime.Value {
	once_Data_Either_monadEither.Do(func() {
		cache_Data_Either_monadEither = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Either_applicativeEither()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Either_bindEither()))}
})}))}
	})
	return cache_Data_Either_monadEither
}

var cache_Data_Either_altEither gopurs_runtime.Value
var once_Data_Either_altEither sync.Once
func Get_Data_Either_altEither() gopurs_runtime.Value {
	once_Data_Either_altEither.Do(func() {
		cache_Data_Either_altEither = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Either_functorEither()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
})}))}
	})
	return cache_Data_Either_altEither
}

type Constructor_Data_Either_Left[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
}


type Constructor_Data_Either_Right[T_a any, T_b any] struct {
	Rc uint32
	V0 T_b
}


func Call_Data_Either_Left__3336285006(__eta_norm_0_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
Left__3336285006:
for {
if false { continue Left__3336285006 }
var __eta_norm_0_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_Either_3123684004_1293498952(__eta_norm_0_0))}, gopurs_runtime.Value{}, false}
}
}

func Call_Data_Either_Left__247931836(__eta_norm_0_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
Left__247931836:
for {
if false { continue Left__247931836 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{__eta_norm_0_0, gopurs_runtime.Value{}, false}
}
}

func Call_Data_Either_Right__3779110556(__eta_norm_0_0_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
Right__3779110556:
for {
if false { continue Right__3779110556 }
var __eta_norm_0_0 *Constructor_Data_Map_Internal_Node[uint32, float64] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Either_261879545_2487766124(__eta_norm_0_0))}, true}
}
}

func Call_Data_Either_Right__1333784983(__eta_norm_0_unused_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
Right__1333784983:
for {
if false { continue Right__1333784983 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, Get_Data_Unit_unit(), true}
}
}

func Call_Data_Either_showEither(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 string
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t0 = (("(Left ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t0 = (("(Right ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
__t0 = func() string { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0)
})}))}
}

func Call_Data_Either_note_prime_(f_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
				_v := gopurs_runtime.Apply2(Get_Data_Maybe_maybe_prime_(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, x_1)}))}
}), Get_Data_Either_Right())
				if _v.Type == 9 && _v.IntVal == 3234899973 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: gopurs_runtime.Value{}, V1: (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V2: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: gopurs_runtime.Value{}, V2: false}
			}()
}

func Call_Data_Either_note(a_0_loop gopurs_runtime.Value, v2_1_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (v2_1 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0}))}
goto end_branch_0
} else {

}
}
{
if (v2_1 != nil) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, (v2_1).V0, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
				_v := __t0
				if _v.Type == 9 && _v.IntVal == 3234899973 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: gopurs_runtime.Value{}, V1: (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V2: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: gopurs_runtime.Value{}, V2: false}
			}()
}

func Call_Data_Either_fromRight_prime_(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Either_fromLeft_prime_(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 bool
{
if (x_2.Type == 9 && x_2.IntVal == 3711209382) {
__t0 = ((y_3.Type == 9 && y_3.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
__t0 = ((x_2.Type == 9 && x_2.IntVal == 2465973597)) && (((y_3.Type == 9 && y_3.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_0:
return gopurs_runtime.Bool(__t0)
})}))}
}

func Call_Data_Either_ordEither(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqEither1__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
eqEither1__193435443_1_0 := gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 bool
{
if (x_3.Type == 9 && x_3.IntVal == 3711209382) {
__t2 = ((y_4.Type == 9 && y_4.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0))
goto end_branch_2
} else {

}
}
{
__t2 = ((x_3.Type == 9 && x_3.IntVal == 2465973597)) && (((y_4.Type == 9 && y_4.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
})}))}
})
_ = eqEither1__193435443_1_0
return gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqEither2_3_3 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Either","Either"] [(TypeVar a), (TypeVar b)])])
eqEither2_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqEither1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})))
_ = eqEither2_3_3
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqEither2_3_3)}
}), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 uint32
{
if (x_4.Type == 9 && x_4.IntVal == 3711209382) {
var __t4 uint32
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t4 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t5 = 380165415
goto end_branch_5
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2465973597)) && ((y_5.Type == 9 && y_5.IntVal == 2465973597)) {
__t5 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
goto end_branch_5
} else {

}
}
{
__t5 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
})}))}
})
}

func Call_Data_Either_eq1Either(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictEq1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 bool
{
if (x_2.Type == 9 && x_2.IntVal == 3711209382) {
__t0 = ((y_3.Type == 9 && y_3.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
__t0 = ((x_2.Type == 9 && x_2.IntVal == 2465973597)) && (((y_3.Type == 9 && y_3.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_1, "eq"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_0:
return gopurs_runtime.Bool(__t0)
})}))}
}

func Call_Data_Either_ord1Either(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): eqEither1__193435443_1_1 shape=Let(Abs(LitRecord)) bindingType=Any
eqEither1__193435443_1_1 := gopurs_runtime.Func(func(dictEq1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
{
if (x_3.Type == 9 && x_3.IntVal == 3711209382) {
__t3 = ((y_4.Type == 9 && y_4.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "eq"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0))
goto end_branch_3
} else {

}
}
{
__t3 = ((x_3.Type == 9 && x_3.IntVal == 2465973597)) && (((y_4.Type == 9 && y_4.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_2, "eq"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})}))}
})
_ = eqEither1__193435443_1_1
// TAST (Let): ordEither1__193435443_1_0 shape=Let(Abs(Let(LitRecord))) bindingType=Any
ordEither1__193435443_1_0 := gopurs_runtime.Func(func(dictOrd1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqEither2_3_4 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Either","Either"] [(TypeVar a), (TypeVar b)])])
eqEither2_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqEither1__193435443_1_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_2, "Eq0"), gopurs_runtime.Value{})))
_ = eqEither2_3_4
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqEither2_3_4)}
}), gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 uint32
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
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if (y_5.Type == 9 && y_5.IntVal == 3711209382) {
__t6 = 380165415
goto end_branch_6
} else {

}
}
{
if ((x_4.Type == 9 && x_4.IntVal == 2465973597)) && ((y_5.Type == 9 && y_5.IntVal == 2465973597)) {
__t6 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_2, "compare"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal)
goto end_branch_6
} else {

}
}
{
__t6 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t6), UnsafePtr: nil}
})}))}
})
_ = ordEither1__193435443_1_0
// TAST (Let): __local_var_2_8 shape=App(Other) bindingType=Any
__local_var_2_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_8
// TAST (Let): eq1Either1_2_7 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Eq1"] [(ADT ["Data","Either","Either"] [(TypeVar a)])])
eq1Either1_2_7 := (&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictEq1_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 bool
{
if (x_4.Type == 9 && x_4.IntVal == 3711209382) {
__t9 = ((y_5.Type == 9 && y_5.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_8, "eq"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0))
goto end_branch_9
} else {

}
}
{
__t9 = ((x_4.Type == 9 && x_4.IntVal == 2465973597)) && (((y_5.Type == 9 && y_5.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_9:
return gopurs_runtime.Bool(__t9)
})})
_ = eq1Either1_2_7
return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1Either1_2_7)}
}), gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordEither1__193435443_1_0, dictOrd1_3), "compare")
})}))}
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

func Call_Data_Either_either__1841469567(v_unused_0_loop gopurs_runtime.Value, v1_unused_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
either__1841469567:
for {
if false { continue either__1841469567 }
var v_unused_0 gopurs_runtime.Value = v_unused_0_loop
_ = v_unused_0
var v1_unused_1 gopurs_runtime.Value = v1_unused_1_loop
_ = v1_unused_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str((*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0.StrVal()))
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0
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
}

func Call_Data_Either_hush(v2_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
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
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlt_0.V0), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictAlt_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Either_Left(), a_2), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Either_Right(), b_3))
})
}

func Call_Data_Either_boundedEither(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): eqEither1__193435443_2_2 shape=Let(Abs(LitRecord)) bindingType=Any
eqEither1__193435443_2_2 := gopurs_runtime.Func(func(dictEq1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 bool
{
if (x_4.Type == 9 && x_4.IntVal == 3711209382) {
__t4 = ((y_5.Type == 9 && y_5.IntVal == 3711209382)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "eq"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0))
goto end_branch_4
} else {

}
}
{
__t4 = ((x_4.Type == 9 && x_4.IntVal == 2465973597)) && (((y_5.Type == 9 && y_5.IntVal == 2465973597)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq1_3, "eq"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0).IntVal) != (0)))
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
})}))}
})
_ = eqEither1__193435443_2_2
// TAST (Let): ordEither1__193435443_1_0 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
ordEither1__193435443_1_0 := gopurs_runtime.Func(func(dictOrd1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqEither2_4_5 shape=App(Other) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Either","Either"] [(TypeVar a), (TypeVar b)])])
eqEither2_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(eqEither1__193435443_2_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_3, "Eq0"), gopurs_runtime.Value{})))
_ = eqEither2_4_5
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqEither2_4_5)}
}), gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 uint32
{
if (x_5.Type == 9 && x_5.IntVal == 3711209382) {
var __t6 uint32
{
if (y_6.Type == 9 && y_6.IntVal == 3711209382) {
__t6 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "compare"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0).IntVal)
goto end_branch_6
} else {

}
}
{
__t6 = 1527465420
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
if (y_6.Type == 9 && y_6.IntVal == 3711209382) {
__t7 = 380165415
goto end_branch_7
} else {

}
}
{
if ((x_5.Type == 9 && x_5.IntVal == 2465973597)) && ((y_6.Type == 9 && y_6.IntVal == 2465973597)) {
__t7 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd1_3, "compare"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0).IntVal)
goto end_branch_7
} else {

}
}
{
__t7 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t7), UnsafePtr: nil}
})}))}
})
_ = ordEither1__193435443_1_0
return gopurs_runtime.Func(func(dictBounded1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordEither2_3_8 shape=App(Other) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Either","Either"] [(TypeVar a), (TypeVar b)])])
ordEither2_3_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(ordEither1__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded1_2, "Ord0"), gopurs_runtime.Value{})))
_ = ordEither2_3_8
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordEither2_3_8)}
}), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.RecordGet(dictBounded_0, "bottom"), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}(), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.RecordGet(dictBounded1_2, "top"), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()}))}
})
}

func Call_Data_Either_blush(v2_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v2_0.Type == 9 && v2_0.IntVal == 3711209382) {
__t0 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_0.UnsafePtr).V0})
goto end_branch_0
} else {

}
}
{
if (v2_0.Type == 9 && v2_0.IntVal == 2465973597) {
__t0 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Either_semigroupEither(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_1.Type == 9 && x_1.IntVal == 3711209382) {
__t1 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_1
} else {

}
}
{
if (x_1.Type == 9 && x_1.IntVal == 2465973597) {
var __t0 gopurs_runtime.Value
{
if (y_2.Type == 9 && y_2.IntVal == 3711209382) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
if (y_2.Type == 9 && y_2.IntVal == 2465973597) {
__t0 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})}))}
}

func Rebox_Data_Either_261879545_2487766124(in *Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
		out.V3 = gopurs_runtime.Float(in.V3)
		out.V4 = Rebox_Data_Either_261879545_2487766124(in.V4)
		out.V5 = Rebox_Data_Either_261879545_2487766124(in.V5)
	return out
}

func Rebox_Data_Either_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}


