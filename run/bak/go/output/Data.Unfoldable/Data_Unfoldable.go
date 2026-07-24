package Data_Unfoldable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		fromJust = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.StrVal == "Just")).IntVal != 0 {
__t0 = (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]
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
	})
	return fromJust
}

var unfoldr gopurs_runtime.Value
var once_unfoldr sync.Once
func Get_unfoldr() gopurs_runtime.Value {
	once_unfoldr.Do(func() {
		unfoldr = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "unfoldr")
})
	})
	return unfoldr
}

var unfoldableMaybe gopurs_runtime.Value
var once_unfoldableMaybe sync.Once
func Get_unfoldableMaybe() gopurs_runtime.Value {
	once_unfoldableMaybe.Do(func() {
		unfoldableMaybe = gopurs_runtime.RecordDict2("unfoldr", "Unfoldable10", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, b_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_2_0.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[0].UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Maybe()
}))
	})
	return unfoldableMaybe
}

var unfoldableArray gopurs_runtime.Value
var once_unfoldableArray sync.Once
func Get_unfoldableArray() gopurs_runtime.Value {
	once_unfoldableArray.Do(func() {
		unfoldableArray = gopurs_runtime.RecordDict2("unfoldr", "Unfoldable10", gopurs_runtime.Apply4(Get_unfoldrArrayImpl(), pkg_Data_Maybe.Get_isNothing(), Get_fromJust(), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Array()
}))
	})
	return unfoldableArray
}

var replicate gopurs_runtime.Value
var once_replicate sync.Once
func Get_replicate() gopurs_runtime.Value {
	once_replicate.Do(func() {
		replicate = gopurs_runtime.Func3(func(dictUnfoldable_0 gopurs_runtime.Value, n_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(i_3.IntVal <= gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", v_2, gopurs_runtime.Int(i_3.IntVal - gopurs_runtime.Int(1).IntVal)))
}
end_branch_0:
return __t0
}), n_1)
})
	})
	return replicate
}

var replicateA gopurs_runtime.Value
var once_replicateA sync.Once
func Get_replicateA() gopurs_runtime.Value {
	once_replicateA.Do(func() {
		replicateA = gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, dictUnfoldable_1 gopurs_runtime.Value, dictTraversable_2 gopurs_runtime.Value) gopurs_runtime.Value {
sequence_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_2, "sequence"), dictApplicative_0)
_ = sequence_3_0
return gopurs_runtime.Func2(func(n_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence_3_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_1, "unfoldr"), gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(i_6.IntVal <= gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", m_5, gopurs_runtime.Int(i_6.IntVal - gopurs_runtime.Int(1).IntVal)))
}
end_branch_1:
return __t1
}), n_4))
})
})
	})
	return replicateA
}

var none gopurs_runtime.Value
var once_none sync.Once
func Get_none() gopurs_runtime.Value {
	once_none.Do(func() {
		none = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
}), pkg_Data_Unit.Get_unit())
})
	})
	return none
}

var fromMaybe gopurs_runtime.Value
var once_fromMaybe sync.Once
func Get_fromMaybe() gopurs_runtime.Value {
	once_fromMaybe.Do(func() {
		fromMaybe = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(b_1.StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(b_1.UnsafePtr)[0], gopurs_runtime.Constructor0("Nothing")))
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
	})
	return fromMaybe
}

func Get_unfoldrArrayImpl() gopurs_runtime.Value {
	return _Gopurs_UnfoldrArrayImpl
}
