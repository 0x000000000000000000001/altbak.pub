package Data_Enum_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	unsafe "unsafe"
)

var cache_foldable1NonEmpty gopurs_runtime.Value
var once_foldable1NonEmpty sync.Once
func Get_foldable1NonEmpty() gopurs_runtime.Value {
	once_foldable1NonEmpty.Do(func() {
		cache_foldable1NonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), pkg_Data_Foldable.Get_foldableArray())
	})
	return cache_foldable1NonEmpty
}

var cache_genBoundedEnum gopurs_runtime.Value
var once_genBoundedEnum sync.Once
func Get_genBoundedEnum() gopurs_runtime.Value {
	once_genBoundedEnum.Do(func() {
		cache_genBoundedEnum = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genBoundedEnum(dictMonadGen_0_box)
})
	})
	return cache_genBoundedEnum
}

func Call_genBoundedEnum(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
elements_1_0 := gopurs_runtime.Apply2(pkg_Control_Monad_Gen.Get_elements(), dictMonadGen_0, Get_foldable1NonEmpty())
_ = elements_1_0
return gopurs_runtime.Func(func(dictBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
Enum1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Enum1"), gopurs_runtime.Value{})
_ = Enum1_3_1
Bounded0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Bounded0"), gopurs_runtime.Value{})
_ = Bounded0_4_2
v_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Enum1_3_1, "succ"), gopurs_runtime.RecordGet(Bounded0_4_2, "bottom"))
_ = v_5_3
var __t4 gopurs_runtime.Value
{
if (v_5_3.Type == 9 && v_5_3.IntVal == 930809136) {
__t4 = gopurs_runtime.Apply(elements_1_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty{gopurs_runtime.RecordGet(Bounded0_4_2, "bottom"), gopurs_runtime.Apply4(pkg_Data_Enum.Get_enumFromTo(), Enum1_3_1, pkg_Data_Unfoldable1.Get_unfoldable1Array(), (*pkg_Data_Maybe.Constructor_Just)(v_5_3.UnsafePtr).V0, gopurs_runtime.RecordGet(Bounded0_4_2, "top"))})})
goto end_branch_4
} else {

}
}
{
if (v_5_3.Type == 9 && v_5_3.IntVal == 3589588149) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.RecordGet(Bounded0_4_2, "bottom"))
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
}


