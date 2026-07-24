package Data_Maybe_Last

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
)

var Last gopurs_runtime.Value
var once_Last sync.Once
func Get_Last() gopurs_runtime.Value {
	once_Last.Do(func() {
		Last = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return Last
}

var showLast gopurs_runtime.Value
var once_showLast sync.Once
func Get_showLast() gopurs_runtime.Value {
	once_showLast.Do(func() {
		showLast = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Str("(Last (Just " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0]).StrVal + "))")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_1.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Str("(Last Nothing)")
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
	return showLast
}

var semigroupLast gopurs_runtime.Value
var once_semigroupLast sync.Once
func Get_semigroupLast() gopurs_runtime.Value {
	once_semigroupLast.Do(func() {
		semigroupLast = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1.StrVal == "Just").IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v1_1.StrVal == "Nothing").IntVal != 0 {
__t0 = v_0
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
	return semigroupLast
}

var ordLast gopurs_runtime.Value
var once_ordLast sync.Once
func Get_ordLast() gopurs_runtime.Value {
	once_ordLast.Do(func() {
		ordLast = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0_loop, "Eq0"), gopurs_runtime.Value{})
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
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0_loop, "compare"), (*[1024]gopurs_runtime.Value)(x_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_4.UnsafePtr)[0])
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
	return ordLast
}

var ord1Last gopurs_runtime.Value
var once_ord1Last sync.Once
func Get_ord1Last() gopurs_runtime.Value {
	once_ord1Last.Do(func() {
		ord1Last = pkg_Data_Maybe.Get_ord1Maybe()
	})
	return ord1Last
}

var newtypeLast gopurs_runtime.Value
var once_newtypeLast sync.Once
func Get_newtypeLast() gopurs_runtime.Value {
	once_newtypeLast.Do(func() {
		newtypeLast = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeLast
}

var monoidLast gopurs_runtime.Value
var once_monoidLast sync.Once
func Get_monoidLast() gopurs_runtime.Value {
	once_monoidLast.Do(func() {
		monoidLast = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupLast()
}))
	})
	return monoidLast
}

var monadLast gopurs_runtime.Value
var once_monadLast sync.Once
func Get_monadLast() gopurs_runtime.Value {
	once_monadLast.Do(func() {
		monadLast = pkg_Data_Maybe.Get_monadMaybe()
	})
	return monadLast
}

var invariantLast gopurs_runtime.Value
var once_invariantLast sync.Once
func Get_invariantLast() gopurs_runtime.Value {
	once_invariantLast.Do(func() {
		invariantLast = pkg_Data_Maybe.Get_invariantMaybe()
	})
	return invariantLast
}

var functorLast gopurs_runtime.Value
var once_functorLast sync.Once
func Get_functorLast() gopurs_runtime.Value {
	once_functorLast.Do(func() {
		functorLast = pkg_Data_Maybe.Get_functorMaybe()
	})
	return functorLast
}

var extendLast gopurs_runtime.Value
var once_extendLast sync.Once
func Get_extendLast() gopurs_runtime.Value {
	once_extendLast.Do(func() {
		extendLast = pkg_Data_Maybe.Get_extendMaybe()
	})
	return extendLast
}

var eqLast gopurs_runtime.Value
var once_eqLast sync.Once
func Get_eqLast() gopurs_runtime.Value {
	once_eqLast.Do(func() {
		eqLast = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(x_1.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(y_2.StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0_loop, "eq"), (*[1024]gopurs_runtime.Value)(x_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(y_2.UnsafePtr)[0]).IntVal != 0)
}
end_branch_0:
return __t0
}))
}()
})
	})
	return eqLast
}

var eq1Last gopurs_runtime.Value
var once_eq1Last sync.Once
func Get_eq1Last() gopurs_runtime.Value {
	once_eq1Last.Do(func() {
		eq1Last = pkg_Data_Maybe.Get_eq1Maybe()
	})
	return eq1Last
}

var boundedLast gopurs_runtime.Value
var once_boundedLast sync.Once
func Get_boundedLast() gopurs_runtime.Value {
	once_boundedLast.Do(func() {
		boundedLast = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.Apply(pkg_Data_Maybe.Get_boundedMaybe(), dictBounded_0_loop)
}()
})
	})
	return boundedLast
}

var bindLast gopurs_runtime.Value
var once_bindLast sync.Once
func Get_bindLast() gopurs_runtime.Value {
	once_bindLast.Do(func() {
		bindLast = pkg_Data_Maybe.Get_bindMaybe()
	})
	return bindLast
}

var applyLast gopurs_runtime.Value
var once_applyLast sync.Once
func Get_applyLast() gopurs_runtime.Value {
	once_applyLast.Do(func() {
		applyLast = pkg_Data_Maybe.Get_applyMaybe()
	})
	return applyLast
}

var applicativeLast gopurs_runtime.Value
var once_applicativeLast sync.Once
func Get_applicativeLast() gopurs_runtime.Value {
	once_applicativeLast.Do(func() {
		applicativeLast = pkg_Data_Maybe.Get_applicativeMaybe()
	})
	return applicativeLast
}

var altLast gopurs_runtime.Value
var once_altLast sync.Once
func Get_altLast() gopurs_runtime.Value {
	once_altLast.Do(func() {
		altLast = gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.RecordGet(Get_semigroupLast(), "append"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}))
	})
	return altLast
}

var plusLast gopurs_runtime.Value
var once_plusLast sync.Once
func Get_plusLast() gopurs_runtime.Value {
	once_plusLast.Do(func() {
		plusLast = gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altLast()
}))
	})
	return plusLast
}

var alternativeLast gopurs_runtime.Value
var once_alternativeLast sync.Once
func Get_alternativeLast() gopurs_runtime.Value {
	once_alternativeLast.Do(func() {
		alternativeLast = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusLast()
}))
	})
	return alternativeLast
}




