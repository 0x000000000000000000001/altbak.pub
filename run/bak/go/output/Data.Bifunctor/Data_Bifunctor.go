package Data_Bifunctor

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var bimap gopurs_runtime.Value
var once_bimap sync.Once
func Get_bimap() gopurs_runtime.Value {
	once_bimap.Do(func() {
		bimap = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "bimap")
})
	})
	return bimap
}

var bivoid gopurs_runtime.Value
var once_bivoid sync.Once
func Get_bivoid() gopurs_runtime.Value {
	once_bivoid.Do(func() {
		bivoid = gopurs_runtime.Func(func(dictBifunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
})
	})
	return bivoid
}

var lmap gopurs_runtime.Value
var once_lmap sync.Once
func Get_lmap() gopurs_runtime.Value {
	once_lmap.Do(func() {
		lmap = gopurs_runtime.Func2(func(dictBifunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), f_1, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return lmap
}

var rmap gopurs_runtime.Value
var once_rmap sync.Once
func Get_rmap() gopurs_runtime.Value {
	once_rmap.Do(func() {
		rmap = gopurs_runtime.Func(func(dictBifunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBifunctor_0, "bimap"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return rmap
}

var bifunctorTuple gopurs_runtime.Value
var once_bifunctorTuple sync.Once
func Get_bifunctorTuple() gopurs_runtime.Value {
	once_bifunctorTuple.Do(func() {
		bifunctorTuple = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_2, "value0")), gopurs_runtime.Apply(g_1, gopurs_runtime.RecordGet(v_2, "value1")))
}))
	})
	return bifunctorTuple
}

var bifunctorEither gopurs_runtime.Value
var once_bifunctorEither sync.Once
func Get_bifunctorEither() gopurs_runtime.Value {
	once_bifunctorEither.Do(func() {
		bifunctorEither = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Left"), gopurs_runtime.Apply(v_0, gopurs_runtime.RecordGet(v2_2, "value0")))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Right"), gopurs_runtime.Apply(v1_1, gopurs_runtime.RecordGet(v2_2, "value0")))
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
	return bifunctorEither
}

var bifunctorConst gopurs_runtime.Value
var once_bifunctorConst sync.Once
func Get_bifunctorConst() gopurs_runtime.Value {
	once_bifunctorConst.Do(func() {
		bifunctorConst = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
}))
	})
	return bifunctorConst
}


