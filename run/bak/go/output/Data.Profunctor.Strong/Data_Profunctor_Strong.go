package Data_Profunctor_Strong

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
)

var strongFn gopurs_runtime.Value
var once_strongFn sync.Once
func Get_strongFn() gopurs_runtime.Value {
	once_strongFn.Do(func() {
		strongFn = gopurs_runtime.RecordDict3("first", "second", "Profunctor0", gopurs_runtime.Func2(func(a2b_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply(a2b_0, gopurs_runtime.ConstructorGet(v_1, 0)), gopurs_runtime.ConstructorGet(v_1, 1))
}), gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Profunctor.Get_profunctorFn()
}))
	})
	return strongFn
}

var second gopurs_runtime.Value
var once_second sync.Once
func Get_second() gopurs_runtime.Value {
	once_second.Do(func() {
		second = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "second")
})
	})
	return second
}

var first gopurs_runtime.Value
var once_first sync.Once
func Get_first() gopurs_runtime.Value {
	once_first.Do(func() {
		first = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "first")
})
	})
	return first
}

var splitStrong gopurs_runtime.Value
var once_splitStrong sync.Once
func Get_splitStrong() gopurs_runtime.Value {
	once_splitStrong.Do(func() {
		splitStrong = gopurs_runtime.Func4(func(dictSemigroupoid_0 gopurs_runtime.Value, dictStrong_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictStrong_1, "second"), r_3), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictStrong_1, "first"), l_2))
})
	})
	return splitStrong
}

var fanout gopurs_runtime.Value
var once_fanout sync.Once
func Get_fanout() gopurs_runtime.Value {
	once_fanout.Do(func() {
		fanout = gopurs_runtime.Func2(func(dictSemigroupoid_0 gopurs_runtime.Value, dictStrong_1 gopurs_runtime.Value) gopurs_runtime.Value {
lcmap_2_0 := gopurs_runtime.Apply(pkg_Data_Profunctor.Get_lcmap(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictStrong_1, "Profunctor0"), gopurs_runtime.Value{}))
_ = lcmap_2_0
return gopurs_runtime.Func2(func(l_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(lcmap_2_0, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", a_5, a_5)
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictStrong_1, "second"), r_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictStrong_1, "first"), l_3)))
})
})
	})
	return fanout
}


