package Data_Unfoldable1

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		fromJust = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.RecordGet(v_0, "value0")
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

var unfoldr1 gopurs_runtime.Value
var once_unfoldr1 sync.Once
func Get_unfoldr1() gopurs_runtime.Value {
	once_unfoldr1.Do(func() {
		unfoldr1 = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "unfoldr1")
})
	})
	return unfoldr1
}

var unfoldable1Maybe gopurs_runtime.Value
var once_unfoldable1Maybe sync.Once
func Get_unfoldable1Maybe() gopurs_runtime.Value {
	once_unfoldable1Maybe.Do(func() {
		unfoldable1Maybe = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(f_0, b_1), "value0"))
}))
	})
	return unfoldable1Maybe
}

var unfoldable1Array gopurs_runtime.Value
var once_unfoldable1Array sync.Once
func Get_unfoldable1Array() gopurs_runtime.Value {
	once_unfoldable1Array.Do(func() {
		unfoldable1Array = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Apply4(Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing(), Get_fromJust(), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return unfoldable1Array
}

var replicate1 gopurs_runtime.Value
var once_replicate1 sync.Once
func Get_replicate1() gopurs_runtime.Value {
	once_replicate1.Do(func() {
		replicate1 = gopurs_runtime.Func3(func(dictUnfoldable1_0 gopurs_runtime.Value, n_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_0, "unfoldr1"), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(i_3.IntVal <= gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), v_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), v_2, gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.Int(i_3.IntVal - gopurs_runtime.Int(1).IntVal)))
}
end_branch_0:
return __t0
}), gopurs_runtime.Int(n_1.IntVal - gopurs_runtime.Int(1).IntVal))
})
	})
	return replicate1
}

var replicate1A gopurs_runtime.Value
var once_replicate1A sync.Once
func Get_replicate1A() gopurs_runtime.Value {
	once_replicate1A.Do(func() {
		replicate1A = gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, dictUnfoldable1_1 gopurs_runtime.Value, dictTraversable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_2, "sequence1"), dictApply_0)
_ = sequence1_3_0
return gopurs_runtime.Func2(func(n_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_3_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_1, "unfoldr1"), gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(i_6.IntVal <= gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), m_5, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), m_5, gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.Int(i_6.IntVal - gopurs_runtime.Int(1).IntVal)))
}
end_branch_1:
return __t1
}), gopurs_runtime.Int(n_4.IntVal - gopurs_runtime.Int(1).IntVal)))
})
})
	})
	return replicate1A
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func2(func(dictUnfoldable1_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_0, "unfoldr1"), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(i_2.IntVal <= gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), v_1, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), v_1, gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.Int(i_2.IntVal - gopurs_runtime.Int(1).IntVal)))
}
end_branch_0:
return __t0
}), gopurs_runtime.Int(0))
})
	})
	return singleton
}

var range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		range_ = gopurs_runtime.Func3(func(dictUnfoldable1_0 gopurs_runtime.Value, start_1 gopurs_runtime.Value, end_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(end_2.IntVal >= start_1.IntVal)).IntVal != 0 {
__t1 = gopurs_runtime.Int(1)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Int(-1)
}
end_branch_1:
__local_var_3_0 := __t1
_ = __local_var_3_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_0, "unfoldr1"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
i_prime_5_2 := gopurs_runtime.Int(i_4.IntVal + __local_var_3_0.IntVal)
_ = i_prime_5_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(i_4.IntVal == end_2.IntVal)).IntVal != 0 {
__t3 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), i_prime_5_2)
}
end_branch_3:
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), i_4, __t3)
}), start_1)
})
	})
	return range_
}

var iterateN gopurs_runtime.Value
var once_iterateN sync.Once
func Get_iterateN() gopurs_runtime.Value {
	once_iterateN.Do(func() {
		iterateN = gopurs_runtime.Func4(func(dictUnfoldable1_0 gopurs_runtime.Value, n_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value, s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_0, "unfoldr1"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "value1").IntVal > gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.Apply(f_2, gopurs_runtime.RecordGet(v_4, "value0")), gopurs_runtime.Int(gopurs_runtime.RecordGet(v_4, "value1").IntVal - gopurs_runtime.Int(1).IntVal)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_0:
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.RecordGet(v_4, "value0"), __t0)
}), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), s_3, gopurs_runtime.Int(n_1.IntVal - gopurs_runtime.Int(1).IntVal)))
})
	})
	return iterateN
}

func Get_unfoldr1ArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Unfoldr1ArrayImpl
}
