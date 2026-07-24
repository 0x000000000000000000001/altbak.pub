package Data_Maybe_First

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
)

var First gopurs_runtime.Value
var once_First sync.Once
func Get_First() gopurs_runtime.Value {
	once_First.Do(func() {
		First = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return First
}

var showFirst gopurs_runtime.Value
var once_showFirst sync.Once
func Get_showFirst() gopurs_runtime.Value {
	once_showFirst.Do(func() {
		showFirst = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1.StrVal == "Just").IntVal != 0 {
__t0 = gopurs_runtime.Str("First ((Just " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0]).StrVal + "))")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_1.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Str("First (Nothing)")
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
	})
	return showFirst
}

var semigroupFirst gopurs_runtime.Value
var once_semigroupFirst sync.Once
func Get_semigroupFirst() gopurs_runtime.Value {
	once_semigroupFirst.Do(func() {
		semigroupFirst = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Just").IntVal != 0 {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
__t0 = v1_1
}
end_branch_0:
return __t0
}))
	})
	return semigroupFirst
}

var ordFirst gopurs_runtime.Value
var once_ordFirst sync.Once
func Get_ordFirst() gopurs_runtime.Value {
	once_ordFirst.Do(func() {
		ordFirst = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})
	})
	return ordFirst
}

var ord1First gopurs_runtime.Value
var once_ord1First sync.Once
func Get_ord1First() gopurs_runtime.Value {
	once_ord1First.Do(func() {
		ord1First = pkg_Data_Maybe.Get_ord1Maybe()
	})
	return ord1First
}

var newtypeFirst gopurs_runtime.Value
var once_newtypeFirst sync.Once
func Get_newtypeFirst() gopurs_runtime.Value {
	once_newtypeFirst.Do(func() {
		newtypeFirst = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeFirst
}

var monoidFirst gopurs_runtime.Value
var once_monoidFirst sync.Once
func Get_monoidFirst() gopurs_runtime.Value {
	once_monoidFirst.Do(func() {
		monoidFirst = gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupFirst()
}))
	})
	return monoidFirst
}

var monadFirst gopurs_runtime.Value
var once_monadFirst sync.Once
func Get_monadFirst() gopurs_runtime.Value {
	once_monadFirst.Do(func() {
		monadFirst = pkg_Data_Maybe.Get_monadMaybe()
	})
	return monadFirst
}

var invariantFirst gopurs_runtime.Value
var once_invariantFirst sync.Once
func Get_invariantFirst() gopurs_runtime.Value {
	once_invariantFirst.Do(func() {
		invariantFirst = pkg_Data_Maybe.Get_invariantMaybe()
	})
	return invariantFirst
}

var functorFirst gopurs_runtime.Value
var once_functorFirst sync.Once
func Get_functorFirst() gopurs_runtime.Value {
	once_functorFirst.Do(func() {
		functorFirst = pkg_Data_Maybe.Get_functorMaybe()
	})
	return functorFirst
}

var extendFirst gopurs_runtime.Value
var once_extendFirst sync.Once
func Get_extendFirst() gopurs_runtime.Value {
	once_extendFirst.Do(func() {
		extendFirst = pkg_Data_Maybe.Get_extendMaybe()
	})
	return extendFirst
}

var eqFirst gopurs_runtime.Value
var once_eqFirst sync.Once
func Get_eqFirst() gopurs_runtime.Value {
	once_eqFirst.Do(func() {
		eqFirst = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})
	})
	return eqFirst
}

var eq1First gopurs_runtime.Value
var once_eq1First sync.Once
func Get_eq1First() gopurs_runtime.Value {
	once_eq1First.Do(func() {
		eq1First = pkg_Data_Maybe.Get_eq1Maybe()
	})
	return eq1First
}

var boundedFirst gopurs_runtime.Value
var once_boundedFirst sync.Once
func Get_boundedFirst() gopurs_runtime.Value {
	once_boundedFirst.Do(func() {
		boundedFirst = gopurs_runtime.Func(func(dictBounded_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Maybe.Get_boundedMaybe(), dictBounded_0)
})
	})
	return boundedFirst
}

var bindFirst gopurs_runtime.Value
var once_bindFirst sync.Once
func Get_bindFirst() gopurs_runtime.Value {
	once_bindFirst.Do(func() {
		bindFirst = pkg_Data_Maybe.Get_bindMaybe()
	})
	return bindFirst
}

var applyFirst gopurs_runtime.Value
var once_applyFirst sync.Once
func Get_applyFirst() gopurs_runtime.Value {
	once_applyFirst.Do(func() {
		applyFirst = pkg_Data_Maybe.Get_applyMaybe()
	})
	return applyFirst
}

var applicativeFirst gopurs_runtime.Value
var once_applicativeFirst sync.Once
func Get_applicativeFirst() gopurs_runtime.Value {
	once_applicativeFirst.Do(func() {
		applicativeFirst = pkg_Data_Maybe.Get_applicativeMaybe()
	})
	return applicativeFirst
}

var altFirst gopurs_runtime.Value
var once_altFirst sync.Once
func Get_altFirst() gopurs_runtime.Value {
	once_altFirst.Do(func() {
		altFirst = gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.RecordGet(Get_semigroupFirst(), "append"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}))
	})
	return altFirst
}

var plusFirst gopurs_runtime.Value
var once_plusFirst sync.Once
func Get_plusFirst() gopurs_runtime.Value {
	once_plusFirst.Do(func() {
		plusFirst = gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Constructor0("Nothing"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altFirst()
}))
	})
	return plusFirst
}

var alternativeFirst gopurs_runtime.Value
var once_alternativeFirst sync.Once
func Get_alternativeFirst() gopurs_runtime.Value {
	once_alternativeFirst.Do(func() {
		alternativeFirst = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusFirst()
}))
	})
	return alternativeFirst
}




