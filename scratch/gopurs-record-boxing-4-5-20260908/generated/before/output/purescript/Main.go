package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_test gopurs_runtime.Value
var once_Main_test sync.Once
func Get_Main_test() gopurs_runtime.Value {
	once_Main_test.Do(func() {
		cache_Main_test = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993((&Constructor_Data_List_Types_Cons[int64]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Main_3704040722_849153993((*Constructor_Data_List_Types_Cons[int64])(nil)))}.IntVal, (*Constructor_Data_List_Types_Cons[int64])(nil)})))}
	})
	return cache_Main_test
}

func Rebox_Main_3704040722_849153993(in *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = Rebox_Main_3704040722_849153993(in.V1)
	return out
}


