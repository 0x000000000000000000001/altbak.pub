package Data_String_Regex_Flags

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Plus "gopurs/output/Control.Plus"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_eqArray gopurs_runtime.Value
var once_eqArray sync.Once
func Get_eqArray() gopurs_runtime.Value {
	once_eqArray.Do(func() {
		cache_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&pkg_Data_Eq.Constructor_Eq[[]string]{1, gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqString(), "eq"))})}
	})
	return cache_eqArray
}

var cache_RegexFlags gopurs_runtime.Value
var once_RegexFlags sync.Once
func Get_RegexFlags() gopurs_runtime.Value {
	once_RegexFlags.Do(func() {
		cache_RegexFlags = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_RegexFlags(x_0_box)
})
	})
	return cache_RegexFlags
}

var cache_unicode gopurs_runtime.Value
var once_unicode sync.Once
func Get_unicode() gopurs_runtime.Value {
	once_unicode.Do(func() {
		cache_unicode = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true)})
	})
	return cache_unicode
}

var cache_sticky gopurs_runtime.Value
var once_sticky sync.Once
func Get_sticky() gopurs_runtime.Value {
	once_sticky.Do(func() {
		cache_sticky = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false)})
	})
	return cache_sticky
}

var cache_showRegexFlags gopurs_runtime.Value
var once_showRegexFlags sync.Once
func Get_showRegexFlags() gopurs_runtime.Value {
	once_showRegexFlags.Do(func() {
		cache_showRegexFlags = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
usedFlags_1_0 := gopurs_runtime.Apply2(Get_append__2285093048(), gopurs_runtime.Array([]gopurs_runtime.Value{}), gopurs_runtime.Apply2(Get_append__2285093048(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("global")
}), gopurs_runtime.Apply2(pkg_Control_Alternative.Get_guard(), pkg_Control_Alternative.Get_alternativeArray(), gopurs_runtime.RecordGet(v_0, "global"))), gopurs_runtime.Apply2(Get_append__2285093048(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("ignoreCase")
}), gopurs_runtime.Apply2(pkg_Control_Alternative.Get_guard(), pkg_Control_Alternative.Get_alternativeArray(), gopurs_runtime.RecordGet(v_0, "ignoreCase"))), gopurs_runtime.Apply2(Get_append__2285093048(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("multiline")
}), gopurs_runtime.Apply2(pkg_Control_Alternative.Get_guard(), pkg_Control_Alternative.Get_alternativeArray(), gopurs_runtime.RecordGet(v_0, "multiline"))), gopurs_runtime.Apply2(Get_append__2285093048(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("dotAll")
}), gopurs_runtime.Apply2(pkg_Control_Alternative.Get_guard(), pkg_Control_Alternative.Get_alternativeArray(), gopurs_runtime.RecordGet(v_0, "dotAll"))), gopurs_runtime.Apply2(Get_append__2285093048(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("sticky")
}), gopurs_runtime.Apply2(pkg_Control_Alternative.Get_guard(), pkg_Control_Alternative.Get_alternativeArray(), gopurs_runtime.RecordGet(v_0, "sticky"))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unicode")
}), gopurs_runtime.Apply2(pkg_Control_Alternative.Get_guard(), pkg_Control_Alternative.Get_alternativeArray(), gopurs_runtime.RecordGet(v_0, "unicode")))))))))
_ = usedFlags_1_0
var __t1 string
{
if (gopurs_runtime.Apply2(Get_eq__3977211983(), usedFlags_1_0, gopurs_runtime.Array([]gopurs_runtime.Value{})).IntVal) != (0) {
__t1 = "noFlags"
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("("), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply2(pkg_Data_String_Common.Get_joinWith(), gopurs_runtime.Str(" <> "), usedFlags_1_0), gopurs_runtime.Str(")"))).StrVal()
}
end_branch_1:
return gopurs_runtime.Str(__t1)
}))
	})
	return cache_showRegexFlags
}

var cache_semigroupRegexFlags gopurs_runtime.Value
var once_semigroupRegexFlags sync.Once
func Get_semigroupRegexFlags() gopurs_runtime.Value {
	once_semigroupRegexFlags.Do(func() {
		cache_semigroupRegexFlags = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_disj__3676519832(), gopurs_runtime.RecordGet(v_0, "dotAll"), gopurs_runtime.RecordGet(v1_1, "dotAll")).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_disj__3676519832(), gopurs_runtime.RecordGet(v_0, "global"), gopurs_runtime.RecordGet(v1_1, "global")).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_disj__3676519832(), gopurs_runtime.RecordGet(v_0, "ignoreCase"), gopurs_runtime.RecordGet(v1_1, "ignoreCase")).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_disj__3676519832(), gopurs_runtime.RecordGet(v_0, "multiline"), gopurs_runtime.RecordGet(v1_1, "multiline")).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_disj__3676519832(), gopurs_runtime.RecordGet(v_0, "sticky"), gopurs_runtime.RecordGet(v1_1, "sticky")).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_disj__3676519832(), gopurs_runtime.RecordGet(v_0, "unicode"), gopurs_runtime.RecordGet(v1_1, "unicode")).IntVal) != (0))})
})
}))
	})
	return cache_semigroupRegexFlags
}

var cache_noFlags gopurs_runtime.Value
var once_noFlags sync.Once
func Get_noFlags() gopurs_runtime.Value {
	once_noFlags.Do(func() {
		cache_noFlags = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_noFlags
}

var cache_newtypeRegexFlags gopurs_runtime.Value
var once_newtypeRegexFlags sync.Once
func Get_newtypeRegexFlags() gopurs_runtime.Value {
	once_newtypeRegexFlags.Do(func() {
		cache_newtypeRegexFlags = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeRegexFlags
}

var cache_multiline gopurs_runtime.Value
var once_multiline sync.Once
func Get_multiline() gopurs_runtime.Value {
	once_multiline.Do(func() {
		cache_multiline = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_multiline
}

var cache_monoidRegexFlags gopurs_runtime.Value
var once_monoidRegexFlags sync.Once
func Get_monoidRegexFlags() gopurs_runtime.Value {
	once_monoidRegexFlags.Do(func() {
		cache_monoidRegexFlags = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupRegexFlags()
}), Get_noFlags())
	})
	return cache_monoidRegexFlags
}

var cache_ignoreCase gopurs_runtime.Value
var once_ignoreCase sync.Once
func Get_ignoreCase() gopurs_runtime.Value {
	once_ignoreCase.Do(func() {
		cache_ignoreCase = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_ignoreCase
}

var cache_global gopurs_runtime.Value
var once_global sync.Once
func Get_global() gopurs_runtime.Value {
	once_global.Do(func() {
		cache_global = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_global
}

var cache_eqRegexFlags gopurs_runtime.Value
var once_eqRegexFlags sync.Once
func Get_eqRegexFlags() gopurs_runtime.Value {
	once_eqRegexFlags.Do(func() {
		cache_eqRegexFlags = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.RecordGet(ra_0, "dotAll"), gopurs_runtime.RecordGet(rb_1, "dotAll")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.RecordGet(ra_0, "global"), gopurs_runtime.RecordGet(rb_1, "global")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.RecordGet(ra_0, "ignoreCase"), gopurs_runtime.RecordGet(rb_1, "ignoreCase")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.RecordGet(ra_0, "multiline"), gopurs_runtime.RecordGet(rb_1, "multiline")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.RecordGet(ra_0, "sticky"), gopurs_runtime.RecordGet(rb_1, "sticky")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.RecordGet(ra_0, "unicode"), gopurs_runtime.RecordGet(rb_1, "unicode")), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqRowNil(), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_0, rb_1))))))).IntVal) != (0))
})
}))
	})
	return cache_eqRegexFlags
}

var cache_dotAll gopurs_runtime.Value
var once_dotAll sync.Once
func Get_dotAll() gopurs_runtime.Value {
	once_dotAll.Do(func() {
		cache_dotAll = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_dotAll
}

var cache_altArray__2010533188 gopurs_runtime.Value
var once_altArray__2010533188 sync.Once
func Get_altArray__2010533188() gopurs_runtime.Value {
	once_altArray__2010533188.Do(func() {
		cache_altArray__2010533188 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"))
	})
	return cache_altArray__2010533188
}

var cache_alternativeArray__2415002109 gopurs_runtime.Value
var once_alternativeArray__2415002109 sync.Once
func Get_alternativeArray__2415002109() gopurs_runtime.Value {
	once_alternativeArray__2415002109.Do(func() {
		cache_alternativeArray__2415002109 = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Plus.Get_plusArray()
}))
	})
	return cache_alternativeArray__2415002109
}

var cache_guard__2168855335 gopurs_runtime.Value
var once_guard__2168855335 sync.Once
func Get_guard__2168855335() gopurs_runtime.Value {
	once_guard__2168855335.Do(func() {
		cache_guard__2168855335 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_guard__2168855335(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_guard__2168855335
}

var cache_applicativeArray__1604836744 gopurs_runtime.Value
var once_applicativeArray__1604836744 sync.Once
func Get_applicativeArray__1604836744() gopurs_runtime.Value {
	once_applicativeArray__1604836744.Do(func() {
		cache_applicativeArray__1604836744 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
}))
	})
	return cache_applicativeArray__1604836744
}

var cache_pure__2935994064 gopurs_runtime.Value
var once_pure__2935994064 sync.Once
func Get_pure__2935994064() gopurs_runtime.Value {
	once_pure__2935994064.Do(func() {
		cache_pure__2935994064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__2935994064(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__2935994064
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

var cache_applyArray__2998472828 gopurs_runtime.Value
var once_applyArray__2998472828 sync.Once
func Get_applyArray__2998472828() gopurs_runtime.Value {
	once_applyArray__2998472828.Do(func() {
		cache_applyArray__2998472828 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), pkg_Control_Apply.Get_arrayApply())
	})
	return cache_applyArray__2998472828
}

var cache_plusArray__4260531026 gopurs_runtime.Value
var once_plusArray__4260531026 sync.Once
func Get_plusArray__4260531026() gopurs_runtime.Value {
	once_plusArray__4260531026.Do(func() {
		cache_plusArray__4260531026 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Alt.Get_altArray()
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
	})
	return cache_plusArray__4260531026
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

var cache_eq__3977211983 gopurs_runtime.Value
var once_eq__3977211983 sync.Once
func Get_eq__3977211983() gopurs_runtime.Value {
	once_eq__3977211983.Do(func() {
		cache_eq__3977211983 = gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[[]string]](Get_eqArray()).V0
	})
	return cache_eq__3977211983
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_functorArray__2747750794 gopurs_runtime.Value
var once_functorArray__2747750794 sync.Once
func Get_functorArray__2747750794() gopurs_runtime.Value {
	once_functorArray__2747750794.Do(func() {
		cache_functorArray__2747750794 = gopurs_runtime.RecordDict1("map", pkg_Data_Functor.Get_arrayMap())
	})
	return cache_functorArray__2747750794
}

var cache_functorArray__361387505 gopurs_runtime.Value
var once_functorArray__361387505 sync.Once
func Get_functorArray__361387505() gopurs_runtime.Value {
	once_functorArray__361387505.Do(func() {
		cache_functorArray__361387505 = gopurs_runtime.RecordDict1("map", pkg_Data_Functor.Get_arrayMap())
	})
	return cache_functorArray__361387505
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_voidLeft__171362140 gopurs_runtime.Value
var once_voidLeft__171362140 sync.Once
func Get_voidLeft__171362140() gopurs_runtime.Value {
	once_voidLeft__171362140.Do(func() {
		cache_voidLeft__171362140 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidLeft__171362140(f_0_box, x_1_box)
})
	})
	return cache_voidLeft__171362140
}

var cache_voidLeft__32301756 gopurs_runtime.Value
var once_voidLeft__32301756 sync.Once
func Get_voidLeft__32301756() gopurs_runtime.Value {
	once_voidLeft__32301756.Do(func() {
		cache_voidLeft__32301756 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidLeft__32301756(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, x_2_box)
})
	})
	return cache_voidLeft__32301756
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj")
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not")
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
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

var cache_append__2285093048 gopurs_runtime.Value
var once_append__2285093048 sync.Once
func Get_append__2285093048() gopurs_runtime.Value {
	once_append__2285093048.Do(func() {
		cache_append__2285093048 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append")
	})
	return cache_append__2285093048
}

var cache_semigroupArray__4207347319 gopurs_runtime.Value
var once_semigroupArray__4207347319 sync.Once
func Get_semigroupArray__4207347319() gopurs_runtime.Value {
	once_semigroupArray__4207347319.Do(func() {
		cache_semigroupArray__4207347319 = gopurs_runtime.RecordDict1("append", pkg_Data_Semigroup.Get_concatArray())
	})
	return cache_semigroupArray__4207347319
}

func Call_RegexFlags(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_guard__2168855335(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
empty_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "empty")
_ = empty_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.IntVal) != (0) {
__t2 = gopurs_runtime.Apply(Applicative0_1_0.V1, pkg_Data_Unit.Get_unit())
goto end_branch_2
} else {

}
}
{
__t2 = empty_2_1
}
end_branch_2:
return __t2
})
}

func Call_pure__2935994064(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_voidLeft__171362140(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}), f_0)
}

func Call_voidLeft__32301756(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), f_1)
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


