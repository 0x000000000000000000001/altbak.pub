package Data_Profunctor_Choice

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	unsafe "unsafe"
)

var cache_right gopurs_runtime.Value
var once_right sync.Once
func Get_right() gopurs_runtime.Value {
	once_right.Do(func() {
		cache_right = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "right")
}()
})
	})
	return cache_right
}

var cache_left gopurs_runtime.Value
var once_left sync.Once
func Get_left() gopurs_runtime.Value {
	once_left.Do(func() {
		cache_left = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "left")
}()
})
	})
	return cache_left
}

var cache_splitChoice gopurs_runtime.Value
var once_splitChoice sync.Once
func Get_splitChoice() gopurs_runtime.Value {
	once_splitChoice.Do(func() {
		cache_splitChoice = gopurs_runtime.Func4(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictChoice_1_box gopurs_runtime.Value, l_2_box gopurs_runtime.Value, r_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_splitChoice(dictSemigroupoid_0_box, dictChoice_1_box, l_2_box, r_3_box)
})
	})
	return cache_splitChoice
}

var cache_fanin gopurs_runtime.Value
var once_fanin sync.Once
func Get_fanin() gopurs_runtime.Value {
	once_fanin.Do(func() {
		cache_fanin = gopurs_runtime.Func2(func(dictSemigroupoid_0_box gopurs_runtime.Value, dictChoice_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fanin(dictSemigroupoid_0_box, dictChoice_1_box)
})
	})
	return cache_fanin
}

var cache_choiceFn gopurs_runtime.Value
var once_choiceFn sync.Once
func Get_choiceFn() gopurs_runtime.Value {
	once_choiceFn.Do(func() {
		cache_choiceFn = gopurs_runtime.RecordDict3("left", "right", "Profunctor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Data_Data_Either_Left)(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{(*pkg_Data_Either.Data_Data_Either_Right)(v1_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordGet(pkg_Data_Either.Get_functorEither(), "map"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Profunctor.Get_profunctorFn()
}))
	})
	return cache_choiceFn
}

func Call_splitChoice(dictSemigroupoid_0_loop gopurs_runtime.Value, dictChoice_1_loop gopurs_runtime.Value, l_2_loop gopurs_runtime.Value, r_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictChoice_1 gopurs_runtime.Value = dictChoice_1_loop
_ = dictChoice_1
var l_2 gopurs_runtime.Value = l_2_loop
_ = l_2
var r_3 gopurs_runtime.Value = r_3_loop
_ = r_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictChoice_1, "right"), r_3), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictChoice_1, "left"), l_2))
}

func Call_fanin(dictSemigroupoid_0_loop gopurs_runtime.Value, dictChoice_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var dictChoice_1 gopurs_runtime.Value = dictChoice_1_loop
_ = dictChoice_1
rmap_2_0 := gopurs_runtime.Apply(pkg_Data_Profunctor.Get_rmap(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictChoice_1, "Profunctor0"), gopurs_runtime.Value{}))
_ = rmap_2_0
return gopurs_runtime.Func2(func(l_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(rmap_2_0, gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 3711209382) {
__t1 = (*pkg_Data_Either.Data_Data_Either_Left)(v2_5.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2465973597) {
__t1 = (*pkg_Data_Either.Data_Data_Either_Right)(v2_5.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictChoice_1, "right"), r_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictChoice_1, "left"), l_3)))
})
}


