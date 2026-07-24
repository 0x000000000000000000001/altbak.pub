package Data_Profunctor_Choice

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	pkg_Data_Either "gopurs/output/Data.Either"
)

var right gopurs_runtime.Value
var once_right sync.Once
func Get_right() gopurs_runtime.Value {
	once_right.Do(func() {
		right = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "right")
})
	})
	return right
}

var left gopurs_runtime.Value
var once_left sync.Once
func Get_left() gopurs_runtime.Value {
	once_left.Do(func() {
		left = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "left")
})
	})
	return left
}

var splitChoice gopurs_runtime.Value
var once_splitChoice sync.Once
func Get_splitChoice() gopurs_runtime.Value {
	once_splitChoice.Do(func() {
		splitChoice = gopurs_runtime.Func4(func(dictSemigroupoid_0 gopurs_runtime.Value, dictChoice_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictChoice_1, "right"), r_3), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictChoice_1, "left"), l_2))
})
	})
	return splitChoice
}

var fanin gopurs_runtime.Value
var once_fanin sync.Once
func Get_fanin() gopurs_runtime.Value {
	once_fanin.Do(func() {
		fanin = gopurs_runtime.Func2(func(dictSemigroupoid_0 gopurs_runtime.Value, dictChoice_1 gopurs_runtime.Value) gopurs_runtime.Value {
rmap_2_0 := gopurs_runtime.Apply(pkg_Data_Profunctor.Get_rmap(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictChoice_1, "Profunctor0"), gopurs_runtime.Value{}))
_ = rmap_2_0
return gopurs_runtime.Func2(func(l_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(rmap_2_0, gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_5.StrVal == "Left").IntVal != 0 {
__t1 = (*[1024]gopurs_runtime.Value)(v2_5.UnsafePtr)[0]
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v2_5.StrVal == "Right").IntVal != 0 {
__t1 = (*[1024]gopurs_runtime.Value)(v2_5.UnsafePtr)[0]
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
})
	})
	return fanin
}

var choiceFn gopurs_runtime.Value
var once_choiceFn sync.Once
func Get_choiceFn() gopurs_runtime.Value {
	once_choiceFn.Do(func() {
		choiceFn = gopurs_runtime.RecordDict3("left", "right", "Profunctor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", gopurs_runtime.Apply(v_0, (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[0]))
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v1_1.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)(v1_1.UnsafePtr)[0])
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
	return choiceFn
}




