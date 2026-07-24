package Control_Monad_List_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var Yield gopurs_runtime.Value
var once_Yield sync.Once
func Get_Yield() gopurs_runtime.Value {
	once_Yield.Do(func() {
		Yield = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Yield", value0, value1)
})
})
	})
	return Yield
}

var Skip gopurs_runtime.Value
var once_Skip sync.Once
func Get_Skip() gopurs_runtime.Value {
	once_Skip.Do(func() {
		Skip = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Skip", value0)
})
	})
	return Skip
}

var Done gopurs_runtime.Value
var once_Done sync.Once
func Get_Done() gopurs_runtime.Value {
	once_Done.Do(func() {
		Done = gopurs_runtime.Constructor0("Done")
	})
	return Done
}

var ListT gopurs_runtime.Value
var once_ListT sync.Once
func Get_ListT() gopurs_runtime.Value {
	once_ListT.Do(func() {
		ListT = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return ListT
}

var wrapLazy gopurs_runtime.Value
var once_wrapLazy sync.Once
func Get_wrapLazy() gopurs_runtime.Value {
	once_wrapLazy.Do(func() {
		wrapLazy = gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor1("Skip", v_1))
})
	})
	return wrapLazy
}

var wrapEffect gopurs_runtime.Value
var once_wrapEffect sync.Once
func Get_wrapEffect() gopurs_runtime.Value {
	once_wrapEffect.Do(func() {
		wrapEffect = gopurs_runtime.Func2(func(dictFunctor_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
})))
}), v_1)
})
	})
	return wrapEffect
}

var unfold gopurs_runtime.Value
var once_unfold sync.Once
func Get_unfold() gopurs_runtime.Value {
	once_unfold.Do(func() {
		unfold = gopurs_runtime.Func3(Call_unfold)
	})
	return unfold
}

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
uncons:
for {
if false { continue uncons }
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Bind1"), gopurs_runtime.Value{}), "bind"), v_2, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3.StrVal == "Yield").IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0], gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]))))
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_3.StrVal == "Skip").IntVal != 0 {
__t1 = gopurs_runtime.Apply2(Get_uncons(), dictMonad_0_loop, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_3.StrVal == "Done").IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Constructor0("Nothing"))
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
}
}()
})
	})
	return uncons
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
uncons1_1_0 := gopurs_runtime.Apply(Get_uncons(), dictMonad_0)
_ = uncons1_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0].UnsafePtr)[1])
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Apply(uncons1_1_0, l_2))
})
})
	})
	return tail
}

var takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		takeWhile = gopurs_runtime.Func(func(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
takeWhile:
for {
if false { continue takeWhile }
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0_loop, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4.StrVal == "Yield").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]).IntVal != 0 {
__local_var_5_3 := gopurs_runtime.Apply2(Get_takeWhile(), dictApplicative_0_loop, f_2)
_ = __local_var_5_3
__t2 = gopurs_runtime.Constructor2("Yield", (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]))
})))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("Done")
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Skip").IntVal != 0 {
__local_var_5_4 := (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]
_ = __local_var_5_4
__local_var_6_5 := gopurs_runtime.Apply2(Get_takeWhile(), dictApplicative_0_loop, f_2)
_ = __local_var_6_5
__t1 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_5_4))
})))
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Done").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Done")
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v_3)
})
}
}()
})
	})
	return takeWhile
}

var scanl gopurs_runtime.Value
var once_scanl sync.Once
func Get_scanl() gopurs_runtime.Value {
	once_scanl.Do(func() {
		scanl = gopurs_runtime.Func4(func(dictMonad_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_unfold(), dictMonad_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_0 := (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]
_ = __local_var_5_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_6.StrVal == "Yield").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(f_1, __local_var_5_0, (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0]), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1])), __local_var_5_0))
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_6.StrVal == "Skip").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor2("Tuple", __local_var_5_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0])), __local_var_5_0))
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_6.StrVal == "Done").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1])
}), gopurs_runtime.Constructor2("Tuple", b_2, l_3))
})
	})
	return scanl
}

var prepend_prime gopurs_runtime.Value
var once_prepend_prime sync.Once
func Get_prepend_prime() gopurs_runtime.Value {
	once_prepend_prime.Do(func() {
		prepend_prime = gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, h_1 gopurs_runtime.Value, t_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor2("Yield", h_1, t_2))
})
	})
	return prepend_prime
}

var prepend gopurs_runtime.Value
var once_prepend sync.Once
func Get_prepend() gopurs_runtime.Value {
	once_prepend.Do(func() {
		prepend = gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, h_1 gopurs_runtime.Value, t_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor2("Yield", h_1, gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return t_2
}))))
})
	})
	return prepend
}

var nil gopurs_runtime.Value
var once_nil sync.Once
func Get_nil() gopurs_runtime.Value {
	once_nil.Do(func() {
		nil = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor0("Done"))
})
	})
	return nil
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor0("Done"))
_ = nil1_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor2("Yield", a_2, gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
}))))
})
})
	})
	return singleton
}

var take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		take = gopurs_runtime.Func(func(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
take:
for {
if false { continue take }
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0_loop, "pure"), gopurs_runtime.Constructor0("Done"))
_ = nil1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0_loop, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if v_3.IntVal == 0 {
__t7 = nil1_1_0
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_5.StrVal == "Yield").IntVal != 0 {
__local_var_6_3 := (*[1024]gopurs_runtime.Value)(v2_5.UnsafePtr)[1]
_ = __local_var_6_3
__local_var_7_4 := gopurs_runtime.Apply2(Get_take(), dictApplicative_0_loop, gopurs_runtime.Int(v_3.IntVal - 1))
_ = __local_var_7_4
__t2 = gopurs_runtime.Constructor2("Yield", (*[1024]gopurs_runtime.Value)(v2_5.UnsafePtr)[0], gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_6_3))
})))
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v2_5.StrVal == "Skip").IntVal != 0 {
__local_var_6_5 := (*[1024]gopurs_runtime.Value)(v2_5.UnsafePtr)[0]
_ = __local_var_6_5
__local_var_7_6 := gopurs_runtime.Apply2(Get_take(), dictApplicative_0_loop, v_3)
_ = __local_var_7_6
__t2 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_6, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_6_5))
})))
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v2_5.StrVal == "Done").IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("Done")
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), v1_4)
}
end_branch_7:
return __t7
})
}
}()
})
	})
	return take
}

var zipWith_prime gopurs_runtime.Value
var once_zipWith_prime sync.Once
func Get_zipWith_prime() gopurs_runtime.Value {
	once_zipWith_prime.Do(func() {
		zipWith_prime = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
zipWith_prime:
for {
if false { continue zipWith_prime }
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Constructor0("Done"))
_ = nil1_2_1
Bind1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_3_2
Functor0_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_3_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = Functor0_4_3
uncons1_5_4 := gopurs_runtime.Apply(Get_uncons(), dictMonad_0_loop)
_ = uncons1_5_4
return gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, fa_7 gopurs_runtime.Value, fb_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_4_3, "map"), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return x_9
})))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_3_2, "bind"), gopurs_runtime.Apply(uncons1_5_4, fa_7), gopurs_runtime.Func(func(ua_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_3_2, "bind"), gopurs_runtime.Apply(uncons1_5_4, fb_8), gopurs_runtime.Func(func(ub_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(ub_10.StrVal == "Nothing").IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), nil1_2_1)
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(ua_9.StrVal == "Nothing").IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), nil1_2_1)
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(ua_9.StrVal == "Just").IntVal != 0 && gopurs_runtime.Bool(ub_10.StrVal == "Just").IntVal != 0 {
__local_var_11_6 := (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(ua_9.UnsafePtr)[0].UnsafePtr)[1]
_ = __local_var_11_6
__local_var_12_7 := (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(ub_10.UnsafePtr)[0].UnsafePtr)[1]
_ = __local_var_12_7
__local_var_13_8 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_zipWith_prime(), dictMonad_0_loop, f_6, __local_var_11_6, __local_var_12_7)
}))
_ = __local_var_13_8
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_4_3, "map"), gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Constructor2("Yield", a_14, __local_var_13_8))
}), gopurs_runtime.Apply2(f_6, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(ua_9.UnsafePtr)[0].UnsafePtr)[0], (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(ub_10.UnsafePtr)[0].UnsafePtr)[0]))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
})))
})
}
}()
})
	})
	return zipWith_prime
}

var zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		zipWith = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
zipWith_prime1_1_0 := gopurs_runtime.Apply(Get_zipWith_prime(), dictMonad_0)
_ = zipWith_prime1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(zipWith_prime1_1_0, gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply2(f_2, a_3, b_4))
}))
})
})
	})
	return zipWith
}

var newtypeListT gopurs_runtime.Value
var once_newtypeListT sync.Once
func Get_newtypeListT() gopurs_runtime.Value {
	once_newtypeListT.Do(func() {
		newtypeListT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeListT
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func3(Call_mapMaybe)
	})
	return mapMaybe
}

var iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		iterate = gopurs_runtime.Func3(func(dictMonad_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_unfold(), dictMonad_0, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(f_1, x_3), x_3)))
}), a_2)
})
	})
	return iterate
}

var repeat gopurs_runtime.Value
var once_repeat sync.Once
func Get_repeat() gopurs_runtime.Value {
	once_repeat.Do(func() {
		repeat = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_iterate(), dictMonad_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return repeat
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
uncons1_1_0 := gopurs_runtime.Apply(Get_uncons(), dictMonad_0)
_ = uncons1_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Apply(uncons1_1_0, l_2))
})
})
	})
	return head
}

var functorListT gopurs_runtime.Value
var once_functorListT sync.Once
func Get_functorListT() gopurs_runtime.Value {
	once_functorListT.Do(func() {
		functorListT = gopurs_runtime.Func(func(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
functorListT:
for {
if false { continue functorListT }
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0_loop, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Yield").IntVal != 0 {
__local_var_4_1 := (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]
_ = __local_var_4_1
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_functorListT(), dictFunctor_0_loop), "map"), f_1)
_ = __local_var_5_2
__t0 = gopurs_runtime.Constructor2("Yield", gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_4_1))
})))
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Skip").IntVal != 0 {
__local_var_4_3 := (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]
_ = __local_var_4_3
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_functorListT(), dictFunctor_0_loop), "map"), f_1)
_ = __local_var_5_4
__t0 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_4_3))
})))
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Done").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Done")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), v_2)
}))
}
}()
})
	})
	return functorListT
}

var fromEffect gopurs_runtime.Value
var once_fromEffect sync.Once
func Get_fromEffect() gopurs_runtime.Value {
	once_fromEffect.Do(func() {
		fromEffect = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor0("Done"))
_ = nil1_1_0
return gopurs_runtime.Func(func(fa_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
}))
_ = __local_var_3_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Yield", a_4, __local_var_3_1)
}), fa_2)
})
})
	})
	return fromEffect
}

var monadTransListT gopurs_runtime.Value
var once_monadTransListT sync.Once
func Get_monadTransListT() gopurs_runtime.Value {
	once_monadTransListT.Do(func() {
		monadTransListT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_fromEffect(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
}))
	})
	return monadTransListT
}

var foldlRec_prime gopurs_runtime.Value
var once_foldlRec_prime sync.Once
func Get_foldlRec_prime() gopurs_runtime.Value {
	once_foldlRec_prime.Do(func() {
		foldlRec_prime = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_2
uncons1_4_3 := gopurs_runtime.Apply(Get_uncons(), Monad0_1_0)
_ = uncons1_4_3
return gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply(uncons1_4_3, gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_10.StrVal == "Nothing").IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Constructor1("Done", __local_var_9_4))
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v_10.StrVal == "Just").IntVal != 0 {
__local_var_11_6 := (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_10.UnsafePtr)[0].UnsafePtr)[1]
_ = __local_var_11_6
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply2(f_5, __local_var_9_4, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_10.UnsafePtr)[0].UnsafePtr)[0]), gopurs_runtime.Func(func(b_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Constructor1("Loop", gopurs_runtime.RecordDict2("a", "b", b_prime_12, __local_var_11_6)))
}))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
}), gopurs_runtime.RecordDict2("a", "b", a_6, b_7))
})
})
	})
	return foldlRec_prime
}

var runListTRec gopurs_runtime.Value
var once_runListTRec sync.Once
func Get_runListTRec() gopurs_runtime.Value {
	once_runListTRec.Do(func() {
		runListTRec = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_foldlRec_prime(), dictMonadRec_0, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
}), pkg_Data_Unit.Get_unit())
})
	})
	return runListTRec
}

var foldlRec gopurs_runtime.Value
var once_foldlRec sync.Once
func Get_foldlRec() gopurs_runtime.Value {
	once_foldlRec.Do(func() {
		foldlRec = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_1
uncons1_3_2 := gopurs_runtime.Apply(Get_uncons(), Monad0_1_0)
_ = uncons1_3_2
return gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(o_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_3 := gopurs_runtime.RecordGet(o_7, "a")
_ = __local_var_8_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(uncons1_3_2, gopurs_runtime.RecordGet(o_7, "b")), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_9.StrVal == "Nothing").IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Constructor1("Done", __local_var_8_3))
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v_9.StrVal == "Just").IntVal != 0 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Constructor1("Loop", gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Apply2(f_4, __local_var_8_3, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_9.UnsafePtr)[0].UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_9.UnsafePtr)[0].UnsafePtr)[1])))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}))
}), gopurs_runtime.RecordDict2("a", "b", a_5, b_6))
})
})
	})
	return foldlRec
}

var foldl_prime gopurs_runtime.Value
var once_foldl_prime sync.Once
func Get_foldl_prime() gopurs_runtime.Value {
	once_foldl_prime.Do(func() {
		foldl_prime = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_1_0
uncons1_2_1 := gopurs_runtime.Apply(Get_uncons(), dictMonad_0)
_ = uncons1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2 gopurs_runtime.Value
_ = loop_4_2
loop_4_2 = gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply(uncons1_2_1, l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_7.StrVal == "Nothing").IntVal != 0 {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_5)
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v_7.StrVal == "Just").IntVal != 0 {
__local_var_8_4 := (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0].UnsafePtr)[1]
_ = __local_var_8_4
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply2(f_3, b_5, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0].UnsafePtr)[0]), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(loop_4_2, a_9, __local_var_8_4)
}))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
})
return loop_4_2
})
})
	})
	return foldl_prime
}

var runListT gopurs_runtime.Value
var once_runListT sync.Once
func Get_runListT() gopurs_runtime.Value {
	once_runListT.Do(func() {
		runListT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_foldl_prime(), dictMonad_0, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
}), pkg_Data_Unit.Get_unit())
})
	})
	return runListT
}

var foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		foldl = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
uncons1_1_0 := gopurs_runtime.Apply(Get_uncons(), dictMonad_0)
_ = uncons1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_3_1 gopurs_runtime.Value
_ = loop_3_1
loop_3_1 = gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(uncons1_1_0, l_5), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_6.StrVal == "Nothing").IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_4)
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_6.StrVal == "Just").IntVal != 0 {
__t2 = gopurs_runtime.Apply2(loop_3_1, gopurs_runtime.Apply2(f_2, b_4, (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0].UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0].UnsafePtr)[1])
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
return loop_3_1
})
})
	})
	return foldl
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func3(Call_filter)
	})
	return filter
}

var dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		dropWhile = gopurs_runtime.Func(func(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
dropWhile:
for {
if false { continue dropWhile }
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0_loop, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4.StrVal == "Yield").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Apply(f_2, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]).IntVal != 0 {
__local_var_5_3 := gopurs_runtime.Apply2(Get_dropWhile(), dictApplicative_0_loop, f_2)
_ = __local_var_5_3
__t2 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]))
})))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor2("Yield", (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1])
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Skip").IntVal != 0 {
__local_var_5_4 := (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]
_ = __local_var_5_4
__local_var_6_5 := gopurs_runtime.Apply2(Get_dropWhile(), dictApplicative_0_loop, f_2)
_ = __local_var_6_5
__t1 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_5_4))
})))
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Done").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Done")
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v_3)
})
}
}()
})
	})
	return dropWhile
}

var drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		drop = gopurs_runtime.Func(func(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
drop:
for {
if false { continue drop }
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0_loop, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if v_2.IntVal == 0 {
__t6 = v1_3
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_4.StrVal == "Yield").IntVal != 0 {
__local_var_5_2 := (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[1]
_ = __local_var_5_2
__local_var_6_3 := gopurs_runtime.Apply2(Get_drop(), dictApplicative_0_loop, gopurs_runtime.Int(v_2.IntVal - 1))
_ = __local_var_6_3
__t1 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_3, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_5_2))
})))
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v2_4.StrVal == "Skip").IntVal != 0 {
__local_var_5_4 := (*[1024]gopurs_runtime.Value)(v2_4.UnsafePtr)[0]
_ = __local_var_5_4
__local_var_6_5 := gopurs_runtime.Apply2(Get_drop(), dictApplicative_0_loop, v_2)
_ = __local_var_6_5
__t1 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_5_4))
})))
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v2_4.StrVal == "Done").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Done")
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v1_3)
}
end_branch_6:
return __t6
})
}
}()
})
	})
	return drop
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, lh_1 gopurs_runtime.Value, t_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor2("Yield", gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), lh_1), t_2))
})
	})
	return cons
}

var unfoldable1ListT gopurs_runtime.Value
var once_unfoldable1ListT sync.Once
func Get_unfoldable1ListT() gopurs_runtime.Value {
	once_unfoldable1ListT.Do(func() {
		unfoldable1ListT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
singleton1_2_1 := gopurs_runtime.Apply(Get_singleton(), Applicative0_1_0)
_ = singleton1_2_1
return gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_2 gopurs_runtime.Value
_ = go__5_2
go__5_2 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1].StrVal == "Nothing").IntVal != 0 {
__t3 = gopurs_runtime.Apply(singleton1_2_1, (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0])
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1].StrVal == "Just").IntVal != 0 {
__local_var_7_4 := (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0]
_ = __local_var_7_4
__local_var_8_5 := (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1].UnsafePtr)[0]
_ = __local_var_8_5
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Constructor2("Yield", gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_4
}))), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__5_2, gopurs_runtime.Apply(f_3, __local_var_8_5))
}))))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
return gopurs_runtime.Apply(go__5_2, gopurs_runtime.Apply(f_3, b_4))
}))
})
	})
	return unfoldable1ListT
}

var unfoldableListT gopurs_runtime.Value
var once_unfoldableListT sync.Once
func Get_unfoldableListT() gopurs_runtime.Value {
	once_unfoldableListT.Do(func() {
		unfoldableListT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Constructor0("Done"))
_ = nil1_2_1
unfoldable1ListT1_3_2 := gopurs_runtime.Apply(Get_unfoldable1ListT(), dictMonad_0)
_ = unfoldable1ListT1_3_2
return gopurs_runtime.RecordDict2("unfoldr", "Unfoldable10", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__6_3 gopurs_runtime.Value
_ = go__6_3
go__6_3 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_7.StrVal == "Nothing").IntVal != 0 {
__t4 = nil1_2_1
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool(v_7.StrVal == "Just").IntVal != 0 {
__local_var_8_5 := (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0].UnsafePtr)[0]
_ = __local_var_8_5
__local_var_9_6 := (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0].UnsafePtr)[1]
_ = __local_var_9_6
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Constructor2("Yield", gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_8_5
}))), gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__6_3, gopurs_runtime.Apply(f_4, __local_var_9_6))
}))))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
return gopurs_runtime.Apply(go__6_3, gopurs_runtime.Apply(f_4, b_5))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return unfoldable1ListT1_3_2
}))
})
	})
	return unfoldableListT
}

var semigroupListT gopurs_runtime.Value
var once_semigroupListT sync.Once
func Get_semigroupListT() gopurs_runtime.Value {
	once_semigroupListT.Do(func() {
		semigroupListT = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_concat(), dictApplicative_0))
})
	})
	return semigroupListT
}

var concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		concat = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4.StrVal == "Yield").IntVal != 0 {
__local_var_5_2 := (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]
_ = __local_var_5_2
__t1 = gopurs_runtime.Constructor2("Yield", (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_concat(), dictApplicative_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_5_2), y_3)
})))
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Skip").IntVal != 0 {
__local_var_5_3 := (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0]
_ = __local_var_5_3
__t1 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_concat(), dictApplicative_0, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_5_3), y_3)
})))
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Done").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return y_3
})))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), x_2)
})
})
	})
	return concat
}

var monoidListT gopurs_runtime.Value
var once_monoidListT sync.Once
func Get_monoidListT() gopurs_runtime.Value {
	once_monoidListT.Do(func() {
		monoidListT = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupListT1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(Get_concat(), dictApplicative_0))
_ = semigroupListT1_1_0
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Constructor0("Done")), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupListT1_1_0
}))
})
	})
	return monoidListT
}

var catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		catMaybes = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_mapMaybe(), dictFunctor_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return catMaybes
}

var monadListT gopurs_runtime.Value
var once_monadListT sync.Once
func Get_monadListT() gopurs_runtime.Value {
	once_monadListT.Do(func() {
		monadListT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeListT(), dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindListT(), dictMonad_0)
}))
})
	})
	return monadListT
}

var bindListT gopurs_runtime.Value
var once_bindListT sync.Once
func Get_bindListT() gopurs_runtime.Value {
	once_bindListT.Do(func() {
		bindListT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
append_1_0 := gopurs_runtime.Apply(Get_concat(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = append_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(fa_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_5.StrVal == "Yield").IntVal != 0 {
__local_var_6_3 := (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0]
_ = __local_var_6_3
__local_var_7_4 := (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1]
_ = __local_var_7_4
__t2 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(append_1_0, gopurs_runtime.Apply(f_4, __local_var_6_3), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_bindListT(), dictMonad_0), "bind"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_7_4), f_4))
})))
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_5.StrVal == "Skip").IntVal != 0 {
__local_var_6_5 := (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0]
_ = __local_var_6_5
__t2 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_bindListT(), dictMonad_0), "bind"), gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_6_5), f_4)
})))
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v_5.StrVal == "Done").IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("Done")
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), fa_3)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyListT(), dictMonad_0)
}))
})
	})
	return bindListT
}

var applyListT gopurs_runtime.Value
var once_applyListT sync.Once
func Get_applyListT() gopurs_runtime.Value {
	once_applyListT.Do(func() {
		applyListT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
functorListT1_1_0 := gopurs_runtime.Apply(Get_functorListT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = functorListT1_1_0
__local_var_2_1 := gopurs_runtime.Apply(Get_bindListT(), dictMonad_0)
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_applicativeListT(), dictMonad_0), "pure"), gopurs_runtime.Apply(f_prime_5, a_prime_6))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorListT1_1_0
}))
})
	})
	return applyListT
}

var applicativeListT gopurs_runtime.Value
var once_applicativeListT sync.Once
func Get_applicativeListT() gopurs_runtime.Value {
	once_applicativeListT.Do(func() {
		applicativeListT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Apply(Get_singleton(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyListT(), dictMonad_0)
}))
})
	})
	return applicativeListT
}

var monadEffectListT gopurs_runtime.Value
var once_monadEffectListT sync.Once
func Get_monadEffectListT() gopurs_runtime.Value {
	once_monadEffectListT.Do(func() {
		monadEffectListT = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadListT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Apply(Get_singleton(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyListT(), Monad0_1_0)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindListT(), Monad0_1_0)
}))
_ = monadListT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(Get_fromEffect(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("liftEffect", "Monad0", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_2_1
}))
})
	})
	return monadEffectListT
}

var monadSTListT gopurs_runtime.Value
var once_monadSTListT sync.Once
func Get_monadSTListT() gopurs_runtime.Value {
	once_monadSTListT.Do(func() {
		monadSTListT = gopurs_runtime.Func(func(dictMonadST_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadListT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Apply(Get_singleton(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyListT(), Monad0_1_0)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindListT(), Monad0_1_0)
}))
_ = monadListT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(Get_fromEffect(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("liftST", "Monad0", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_2_1
}))
})
	})
	return monadSTListT
}

var altListT gopurs_runtime.Value
var once_altListT sync.Once
func Get_altListT() gopurs_runtime.Value {
	once_altListT.Do(func() {
		altListT = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
functorListT1_1_0 := gopurs_runtime.Apply(Get_functorListT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = functorListT1_1_0
return gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Apply(Get_concat(), dictApplicative_0), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorListT1_1_0
}))
})
	})
	return altListT
}

var plusListT gopurs_runtime.Value
var once_plusListT sync.Once
func Get_plusListT() gopurs_runtime.Value {
	once_plusListT.Do(func() {
		plusListT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
altListT1_2_1 := gopurs_runtime.Apply(Get_altListT(), Applicative0_1_0)
_ = altListT1_2_1
return gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Constructor0("Done")), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return altListT1_2_1
}))
})
	})
	return plusListT
}

var alternativeListT gopurs_runtime.Value
var once_alternativeListT sync.Once
func Get_alternativeListT() gopurs_runtime.Value {
	once_alternativeListT.Do(func() {
		alternativeListT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeListT1_1_0 := gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Apply(Get_singleton(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyListT(), dictMonad_0)
}))
_ = applicativeListT1_1_0
plusListT1_2_1 := gopurs_runtime.Apply(Get_plusListT(), dictMonad_0)
_ = plusListT1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeListT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusListT1_2_1
}))
})
	})
	return alternativeListT
}

var monadPlusListT gopurs_runtime.Value
var once_monadPlusListT sync.Once
func Get_monadPlusListT() gopurs_runtime.Value {
	once_monadPlusListT.Do(func() {
		monadPlusListT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadListT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Apply(Get_singleton(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyListT(), dictMonad_0)
}))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindListT(), dictMonad_0)
}))
_ = monadListT1_1_0
alternativeListT1_2_1 := gopurs_runtime.Apply(Get_alternativeListT(), dictMonad_0)
_ = alternativeListT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "Alternative1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeListT1_2_1
}))
})
	})
	return monadPlusListT
}

func Call_unfold(dictMonad_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
unfold:
for {
if false { continue unfold }
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var z_2 gopurs_runtime.Value = z_2_loop
_ = z_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0_loop, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Just").IntVal != 0 {
__local_var_4_1 := (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0].UnsafePtr)[0]
_ = __local_var_4_1
__t0 = gopurs_runtime.Constructor2("Yield", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0].UnsafePtr)[1], gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_unfold(), dictMonad_0_loop, f_1_loop, __local_var_4_1)
})))
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Nothing").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Done")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Apply(f_1_loop, z_2_loop))
}
}

func Call_mapMaybe(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapMaybe:
for {
if false { continue mapMaybe }
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0_loop, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Yield").IntVal != 0 {
__local_var_4_1 := gopurs_runtime.Apply(f_1_loop, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0])
_ = __local_var_4_1
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_4_1.StrVal == "Just").IntVal != 0 {
__local_var_5_4 := gopurs_runtime.Apply2(Get_mapMaybe(), dictFunctor_0_loop, f_1_loop)
_ = __local_var_5_4
__t3 = gopurs_runtime.Constructor2("Yield", (*[1024]gopurs_runtime.Value)(__local_var_4_1.UnsafePtr)[0], gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]))
})))
goto end_branch_3
} else {

}
}
{
__local_var_5_2 := gopurs_runtime.Apply2(Get_mapMaybe(), dictFunctor_0_loop, f_1_loop)
_ = __local_var_5_2
__t3 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]))
})))
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Skip").IntVal != 0 {
__local_var_4_5 := (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]
_ = __local_var_4_5
__local_var_5_6 := gopurs_runtime.Apply2(Get_mapMaybe(), dictFunctor_0_loop, f_1_loop)
_ = __local_var_5_6
__t0 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_6, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_4_5))
})))
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Done").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Done")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), v_2_loop)
}
}

func Call_filter(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
filter:
for {
if false { continue filter }
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0_loop, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Yield").IntVal != 0 {
__local_var_4_1 := (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]
_ = __local_var_4_1
__local_var_5_2 := gopurs_runtime.Apply2(Get_filter(), dictFunctor_0_loop, f_1_loop)
_ = __local_var_5_2
s_prime_6_3 := gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_4_1))
}))
_ = s_prime_6_3
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Apply(f_1_loop, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]).IntVal != 0 {
__t4 = gopurs_runtime.Constructor2("Yield", (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0], s_prime_6_3)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Constructor1("Skip", s_prime_6_3)
}
end_branch_4:
__t0 = __t4
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Skip").IntVal != 0 {
__local_var_4_5 := (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]
_ = __local_var_4_5
__local_var_5_6 := gopurs_runtime.Apply2(Get_filter(), dictFunctor_0_loop, f_1_loop)
_ = __local_var_5_6
__t0 = gopurs_runtime.Constructor1("Skip", gopurs_runtime.Apply(pkg_Data_Lazy.Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_6, gopurs_runtime.Apply(pkg_Data_Lazy.Get_force(), __local_var_4_5))
})))
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Done").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Done")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), v_2_loop)
}
}


