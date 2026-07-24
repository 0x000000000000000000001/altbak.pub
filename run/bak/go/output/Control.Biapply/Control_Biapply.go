package Control_Biapply

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Function "gopurs/output/Data.Function"
)

var biapplyTuple gopurs_runtime.Value
var once_biapplyTuple sync.Once
func Get_biapplyTuple() gopurs_runtime.Value {
	once_biapplyTuple.Do(func() {
		biapplyTuple = gopurs_runtime.RecordDict2("biapply", "Bifunctor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[0]), gopurs_runtime.Apply((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[1]))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorTuple()
}))
	})
	return biapplyTuple
}

var biapply gopurs_runtime.Value
var once_biapply sync.Once
func Get_biapply() gopurs_runtime.Value {
	once_biapply.Do(func() {
		biapply = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "biapply")
})
	})
	return biapply
}

var biapplyFirst gopurs_runtime.Value
var once_biapplyFirst sync.Once
func Get_biapplyFirst() gopurs_runtime.Value {
	once_biapplyFirst.Do(func() {
		biapplyFirst = gopurs_runtime.Func3(func(dictBiapply_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{}), "bimap"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity")
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity")
}), a_1), b_2)
})
	})
	return biapplyFirst
}

var biapplySecond gopurs_runtime.Value
var once_biapplySecond sync.Once
func Get_biapplySecond() gopurs_runtime.Value {
	once_biapplySecond.Do(func() {
		biapplySecond = gopurs_runtime.Func3(func(dictBiapply_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{}), "bimap"), pkg_Data_Function.Get_const_(), pkg_Data_Function.Get_const_(), a_1), b_2)
})
	})
	return biapplySecond
}

var bilift2 gopurs_runtime.Value
var once_bilift2 sync.Once
func Get_bilift2() gopurs_runtime.Value {
	once_bilift2.Do(func() {
		bilift2 = gopurs_runtime.Func5(func(dictBiapply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{}), "bimap"), f_1, g_2, a_3), b_4)
})
	})
	return bilift2
}

var bilift3 gopurs_runtime.Value
var once_bilift3 sync.Once
func Get_bilift3() gopurs_runtime.Value {
	once_bilift3.Do(func() {
		bilift3 = gopurs_runtime.Func5(func(dictBiapply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBiapply_0, "biapply"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBiapply_0, "Bifunctor0"), gopurs_runtime.Value{}), "bimap"), f_1, g_2, a_3), b_4), c_5)
})
})
	})
	return bilift3
}


