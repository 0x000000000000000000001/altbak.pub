package AppX

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_App "gopurs/output/App"
	pkg_Test_FileOps "gopurs/output/Test.FileOps"
	pkg_Test_STArray "gopurs/output/Test.STArray"
	pkg_Test_StringOps "gopurs/output/Test.StringOps"
	pkg_Test_AffOperations "gopurs/output/Test.AffOperations"
)

var cache_main gopurs_runtime.Value
var once_main sync.Once
func Get_main() gopurs_runtime.Value {
	once_main.Do(func() {
		cache_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
_dollar__unused_0_0 := gopurs_runtime.Apply(pkg_App.Get_main(), gopurs_runtime.Value{})
_ = _dollar__unused_0_0
_dollar__unused_1_1 := gopurs_runtime.Apply(pkg_Test_FileOps.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_1_1
_dollar__unused_2_2 := gopurs_runtime.Apply(pkg_Test_FileOps.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_2_2
_dollar__unused_3_3 := gopurs_runtime.Apply(pkg_Test_STArray.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_3_3
_dollar__unused_4_4 := gopurs_runtime.Apply(pkg_Test_STArray.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_4_4
_dollar__unused_5_5 := gopurs_runtime.Apply(pkg_Test_StringOps.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_5_5
_dollar__unused_6_6 := gopurs_runtime.Apply(pkg_Test_StringOps.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_6_6
_dollar__unused_7_7 := gopurs_runtime.Apply(pkg_Test_AffOperations.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_7_7
return gopurs_runtime.Apply(pkg_Test_AffOperations.Get_act(), gopurs_runtime.Value{})
})
	})
	return cache_main
}




