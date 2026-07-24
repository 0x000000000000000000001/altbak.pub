package Data_Maybe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var Nothing gopurs_runtime.Value
var once_Nothing sync.Once
func Get_Nothing() gopurs_runtime.Value {
	once_Nothing.Do(func() {
		Nothing = gopurs_runtime.Constructor0("Nothing")
	})
	return Nothing
}

var Just gopurs_runtime.Value
var once_Just sync.Once
func Get_Just() gopurs_runtime.Value {
	once_Just.Do(func() {
		Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Just", value0)
})
	})
	return Just
}

var showMaybe gopurs_runtime.Value
var once_showMaybe sync.Once
func Get_showMaybe() gopurs_runtime.Value {
	once_showMaybe.Do(func() {
		showMaybe = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Str("(Just " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0]).StrVal + ")")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_1.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Str("Nothing")
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
	return showMaybe
}

var semigroupMaybe gopurs_runtime.Value
var once_semigroupMaybe sync.Once
func Get_semigroupMaybe() gopurs_runtime.Value {
	once_semigroupMaybe.Do(func() {
		semigroupMaybe = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1.StrVal == "Nothing").IntVal != 0 {
__t0 = v1_2
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v1_2.StrVal == "Nothing").IntVal != 0 {
__t0 = v_1
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_1.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(v1_2.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[0]))
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
	return semigroupMaybe
}

var optional gopurs_runtime.Value
var once_optional sync.Once
func Get_optional() gopurs_runtime.Value {
	once_optional.Do(func() {
		optional = gopurs_runtime.Func3(func(dictAlt_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_optional(dictAlt_0_box, dictApplicative_1_box, a_2_box)
})
	})
	return optional
}

var monoidMaybe gopurs_runtime.Value
var once_monoidMaybe sync.Once
func Get_monoidMaybe() gopurs_runtime.Value {
	once_monoidMaybe.Do(func() {
		monoidMaybe = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
semigroupMaybe1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1.StrVal == "Nothing").IntVal != 0 {
__t1 = v1_2
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_2.StrVal == "Nothing").IntVal != 0 {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_1.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(v1_2.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[0]))
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
_ = semigroupMaybe1_1_0
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMaybe1_1_0
}))
}()
})
	})
	return monoidMaybe
}

var maybe_prime gopurs_runtime.Value
var once_maybe_prime sync.Once
func Get_maybe_prime() gopurs_runtime.Value {
	once_maybe_prime.Do(func() {
		maybe_prime = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe_prime(v_0_box, v1_1_box, v2_2_box)
})
	})
	return maybe_prime
}

var maybe gopurs_runtime.Value
var once_maybe sync.Once
func Get_maybe() gopurs_runtime.Value {
	once_maybe.Do(func() {
		maybe = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe(v_0_box, v1_1_box, v2_2_box)
})
	})
	return maybe
}

var isNothing gopurs_runtime.Value
var once_isNothing sync.Once
func Get_isNothing() gopurs_runtime.Value {
	once_isNothing.Do(func() {
		isNothing = gopurs_runtime.Func(func(v2_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_0.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_0.StrVal == "Just").IntVal != 0 {
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
	return isNothing
}

var isJust gopurs_runtime.Value
var once_isJust sync.Once
func Get_isJust() gopurs_runtime.Value {
	once_isJust.Do(func() {
		isJust = gopurs_runtime.Func(func(v2_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_0 gopurs_runtime.Value = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_0.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_0.StrVal == "Just").IntVal != 0 {
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
	return isJust
}

var genericMaybe gopurs_runtime.Value
var once_genericMaybe sync.Once
func Get_genericMaybe() gopurs_runtime.Value {
	once_genericMaybe.Do(func() {
		genericMaybe = gopurs_runtime.RecordDict2("to", "from", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_0.StrVal == "Inl").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(x_0.StrVal == "Inr").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(x_0.UnsafePtr)[0])
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
if gopurs_runtime.Bool(x_0.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Inl", gopurs_runtime.Constructor0("NoArguments"))
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(x_0.StrVal == "Just").IntVal != 0 {
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
	return genericMaybe
}

var functorMaybe gopurs_runtime.Value
var once_functorMaybe sync.Once
func Get_functorMaybe() gopurs_runtime.Value {
	once_functorMaybe.Do(func() {
		functorMaybe = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(v_0, (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[0]))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}))
	})
	return functorMaybe
}

var invariantMaybe gopurs_runtime.Value
var once_invariantMaybe sync.Once
func Get_invariantMaybe() gopurs_runtime.Value {
	once_invariantMaybe.Do(func() {
		invariantMaybe = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_2.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[0]))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}))
	})
	return invariantMaybe
}

var fromMaybe_prime gopurs_runtime.Value
var once_fromMaybe_prime sync.Once
func Get_fromMaybe_prime() gopurs_runtime.Value {
	once_fromMaybe_prime.Do(func() {
		fromMaybe_prime = gopurs_runtime.Func(func(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply2(Get_maybe_prime(), a_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}()
})
	})
	return fromMaybe_prime
}

var fromMaybe gopurs_runtime.Value
var once_fromMaybe sync.Once
func Get_fromMaybe() gopurs_runtime.Value {
	once_fromMaybe.Do(func() {
		fromMaybe = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe(a_0_box, v2_1_box)
})
	})
	return fromMaybe
}

var fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		fromJust = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust(_dollar__unused_0_box, v_1_box)
})
	})
	return fromJust
}

var extendMaybe gopurs_runtime.Value
var once_extendMaybe sync.Once
func Get_extendMaybe() gopurs_runtime.Value {
	once_extendMaybe.Do(func() {
		extendMaybe = gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(v_0, v1_1))
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMaybe()
}))
	})
	return extendMaybe
}

var eqMaybe gopurs_runtime.Value
var once_eqMaybe sync.Once
func Get_eqMaybe() gopurs_runtime.Value {
	once_eqMaybe.Do(func() {
		eqMaybe = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_1.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_2.StrVal == "Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_1.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_2.StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0]).IntVal != 0)
}
end_branch_0:
return __t0
}))
}()
})
	})
	return eqMaybe
}

var ordMaybe gopurs_runtime.Value
var once_ordMaybe sync.Once
func Get_ordMaybe() gopurs_runtime.Value {
	once_ordMaybe.Do(func() {
		ordMaybe = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
eqMaybe1_2_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_2.StrVal == "Nothing").IntVal != 0 {
__t2 = gopurs_runtime.Bool(y_3.StrVal == "Nothing")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_2.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_3.StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_3.UnsafePtr)[0]).IntVal != 0)
}
end_branch_2:
return __t2
}))
_ = eqMaybe1_2_1
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_3.StrVal == "Nothing").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(y_4.StrVal == "Nothing").IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("EQ")
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Constructor0("LT")
}
end_branch_4:
__t3 = __t4
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(y_4.StrVal == "Nothing").IntVal != 0 {
__t3 = gopurs_runtime.Constructor0("GT")
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(x_3.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_4.StrVal == "Just").IntVal != 0 {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*[1024]gopurs_runtime.Value)(x_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_2_1
}))
}()
})
	})
	return ordMaybe
}

var eq1Maybe gopurs_runtime.Value
var once_eq1Maybe sync.Once
func Get_eq1Maybe() gopurs_runtime.Value {
	once_eq1Maybe.Do(func() {
		eq1Maybe = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_1.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Bool(y_2.StrVal == "Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_1.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_2.StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0]).IntVal != 0)
}
end_branch_0:
return __t0
}))
	})
	return eq1Maybe
}

var ord1Maybe gopurs_runtime.Value
var once_ord1Maybe sync.Once
func Get_ord1Maybe() gopurs_runtime.Value {
	once_ord1Maybe.Do(func() {
		ord1Maybe = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_1.StrVal == "Nothing").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(y_2.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("EQ")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("LT")
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(y_2.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("GT")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(x_1.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_2.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0])
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
return Get_eq1Maybe()
}))
	})
	return ord1Maybe
}

var boundedMaybe gopurs_runtime.Value
var once_boundedMaybe sync.Once
func Get_boundedMaybe() gopurs_runtime.Value {
	once_boundedMaybe.Do(func() {
		boundedMaybe = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_1
eqMaybe1_3_3 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_3.StrVal == "Nothing").IntVal != 0 {
__t4 = gopurs_runtime.Bool(y_4.StrVal == "Nothing")
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_3.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_4.StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "eq"), (*[1024]gopurs_runtime.Value)(x_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0]).IntVal != 0)
}
end_branch_4:
return __t4
}))
_ = eqMaybe1_3_3
ordMaybe1_3_2 := gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_4.StrVal == "Nothing").IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if gopurs_runtime.Bool(y_5.StrVal == "Nothing").IntVal != 0 {
__t6 = gopurs_runtime.Constructor0("EQ")
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Constructor0("LT")
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(y_5.StrVal == "Nothing").IntVal != 0 {
__t5 = gopurs_runtime.Constructor0("GT")
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(x_4.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_5.StrVal == "Just").IntVal != 0 {
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "compare"), (*[1024]gopurs_runtime.Value)(x_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_5.UnsafePtr)[0])
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqMaybe1_3_3
}))
_ = ordMaybe1_3_2
return gopurs_runtime.RecordDict3("top", "bottom", "Ord0", gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordGet(dictBounded_0, "top")), gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ordMaybe1_3_2
}))
}()
})
	})
	return boundedMaybe
}

var applyMaybe gopurs_runtime.Value
var once_applyMaybe sync.Once
func Get_applyMaybe() gopurs_runtime.Value {
	once_applyMaybe.Do(func() {
		applyMaybe = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Just").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Nothing").IntVal != 0 {
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
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorMaybe()
}))
	})
	return applyMaybe
}

var bindMaybe gopurs_runtime.Value
var once_bindMaybe sync.Once
func Get_bindMaybe() gopurs_runtime.Value {
	once_bindMaybe.Do(func() {
		bindMaybe = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Apply(v1_1, (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Nothing").IntVal != 0 {
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
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMaybe()
}))
	})
	return bindMaybe
}

var semiringMaybe gopurs_runtime.Value
var once_semiringMaybe sync.Once
func Get_semiringMaybe() gopurs_runtime.Value {
	once_semiringMaybe.Do(func() {
		semiringMaybe = gopurs_runtime.Func(func(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.RecordDict4("zero", "one", "add", "mul", gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordGet(dictSemiring_0, "one")), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1.StrVal == "Nothing").IntVal != 0 {
__t0 = v1_2
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v1_2.StrVal == "Nothing").IntVal != 0 {
__t0 = v_1
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_1.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(v1_2.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[0]))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(x_1.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_2.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}))
}()
})
	})
	return semiringMaybe
}

var applicativeMaybe gopurs_runtime.Value
var once_applicativeMaybe sync.Once
func Get_applicativeMaybe() gopurs_runtime.Value {
	once_applicativeMaybe.Do(func() {
		applicativeMaybe = gopurs_runtime.RecordDict2("pure", "Apply0", Get_Just(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyMaybe()
}))
	})
	return applicativeMaybe
}

var monadMaybe gopurs_runtime.Value
var once_monadMaybe sync.Once
func Get_monadMaybe() gopurs_runtime.Value {
	once_monadMaybe.Do(func() {
		monadMaybe = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindMaybe()
}))
	})
	return monadMaybe
}

var altMaybe gopurs_runtime.Value
var once_altMaybe sync.Once
func Get_altMaybe() gopurs_runtime.Value {
	once_altMaybe.Do(func() {
		altMaybe = gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Nothing").IntVal != 0 {
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
return Get_functorMaybe()
}))
	})
	return altMaybe
}

var plusMaybe gopurs_runtime.Value
var once_plusMaybe sync.Once
func Get_plusMaybe() gopurs_runtime.Value {
	once_plusMaybe.Do(func() {
		plusMaybe = gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altMaybe()
}))
	})
	return plusMaybe
}

var alternativeMaybe gopurs_runtime.Value
var once_alternativeMaybe sync.Once
func Get_alternativeMaybe() gopurs_runtime.Value {
	once_alternativeMaybe.Do(func() {
		alternativeMaybe = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusMaybe()
}))
	})
	return alternativeMaybe
}

func Call_optional(dictAlt_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{}), "map"), Get_Just(), a_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_1, "pure"), gopurs_runtime.Constructor0("Nothing")))
}

func Call_maybe_prime(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_2.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_2.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Apply(v1_1, (*[1024]gopurs_runtime.Value)(v2_2.UnsafePtr)[0])
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

func Call_maybe(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_2.StrVal == "Nothing").IntVal != 0 {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_2.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Apply(v1_1, (*[1024]gopurs_runtime.Value)(v2_2.UnsafePtr)[0])
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

func Call_fromMaybe(a_0_loop gopurs_runtime.Value, v2_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 gopurs_runtime.Value = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_1.StrVal == "Nothing").IntVal != 0 {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v2_1.StrVal == "Just").IntVal != 0 {
__t0 = (*[1024]gopurs_runtime.Value)(v2_1.UnsafePtr)[0]
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

func Call_fromJust(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1.StrVal == "Just").IntVal != 0 {
__t0 = (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0]
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


