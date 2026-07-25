package Data_Bifunctor

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_bimap gopurs_runtime.Value
var once_bimap sync.Once
func Get_bimap() gopurs_runtime.Value {
	once_bimap.Do(func() {
		cache_bimap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap(dict_0_box)
})
	})
	return cache_bimap
}

var cache_bivoid gopurs_runtime.Value
var once_bivoid sync.Once
func Get_bivoid() gopurs_runtime.Value {
	once_bivoid.Do(func() {
		cache_bivoid = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bivoid(dictBifunctor_0_box)
})
	})
	return cache_bivoid
}

var cache_lmap gopurs_runtime.Value
var once_lmap sync.Once
func Get_lmap() gopurs_runtime.Value {
	once_lmap.Do(func() {
		cache_lmap = gopurs_runtime.Func2(func(dictBifunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lmap(dictBifunctor_0_box, f_1_box)
})
	})
	return cache_lmap
}

var cache_rmap gopurs_runtime.Value
var once_rmap sync.Once
func Get_rmap() gopurs_runtime.Value {
	once_rmap.Do(func() {
		cache_rmap = gopurs_runtime.Func(func(dictBifunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rmap(dictBifunctor_0_box)
})
	})
	return cache_rmap
}

var cache_bifunctorTuple gopurs_runtime.Value
var once_bifunctorTuple sync.Once
func Get_bifunctorTuple() gopurs_runtime.Value {
	once_bifunctorTuple.Do(func() {
		cache_bifunctorTuple = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_2.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)})}
}))
	})
	return cache_bifunctorTuple
}

var cache_bifunctorEither gopurs_runtime.Value
var once_bifunctorEither sync.Once
func Get_bifunctorEither() gopurs_runtime.Value {
	once_bifunctorEither.Do(func() {
		cache_bifunctorEither = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Data_Data_Either_Left)(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Data_Data_Either_Right)(v2_2.UnsafePtr).V0)})}
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
	return cache_bifunctorEither
}

var cache_bifunctorConst gopurs_runtime.Value
var once_bifunctorConst sync.Once
func Get_bifunctorConst() gopurs_runtime.Value {
	once_bifunctorConst.Do(func() {
		cache_bifunctorConst = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
}))
	})
	return cache_bifunctorConst
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_bimap(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return ((*gopurs_runtime.RecordData1)(dict_0.UnsafePtr)).V0
}

func Call_bivoid(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictBifunctor_0.UnsafePtr)).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
}

func Call_lmap(dictBifunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictBifunctor_0.UnsafePtr)).V0, f_1, Get_identity())
}

func Call_rmap(dictBifunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifunctor_0 gopurs_runtime.Value = dictBifunctor_0_loop
_ = dictBifunctor_0
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictBifunctor_0.UnsafePtr)).V0, Get_identity())
}


