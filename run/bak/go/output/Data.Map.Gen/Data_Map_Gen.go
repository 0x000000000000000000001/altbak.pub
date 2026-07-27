package Data_Map_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var cache_genMap gopurs_runtime.Value
var once_genMap sync.Once
func Get_genMap() gopurs_runtime.Value {
	once_genMap.Do(func() {
		cache_genMap = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genMap(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genMap
}

func Call_genMap(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
Bind1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_0
Apply0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_3_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_2
unfoldable1_5_3 := gopurs_runtime.Apply3(pkg_Control_Monad_Gen.Get_unfoldable(), dictMonadRec_0, dictMonadGen_1, pkg_Data_List_Types.Get_unfoldableList())
_ = unfoldable1_5_3
return gopurs_runtime.Func(func(dictOrd_6 gopurs_runtime.Value) gopurs_runtime.Value {
fromFoldable_7_4 := gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_fromFoldable(), dictOrd_6, pkg_Data_List_Types.Get_foldableList())
_ = fromFoldable_7_4
return gopurs_runtime.Func2(func(genKey_8 gopurs_runtime.Value, genValue_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(size_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_0, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(0), size_10), gopurs_runtime.Func(func(newSize_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "resize"), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return newSize_11
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "map"), fromFoldable_7_4, gopurs_runtime.Apply(unfoldable1_5_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_3_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "map"), pkg_Data_Tuple.Get_Tuple(), genKey_8), genValue_9))))
}))
}))
})
})
}
