package AppX

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Test_RBTree "gopurs/output/Test.RBTree"
)

var main gopurs_runtime.Value
var once_main sync.Once
func Get_main() gopurs_runtime.Value {
	once_main.Do(func() {
		main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
_dollar__unused_0_0 := gopurs_runtime.Apply(pkg_Test_RBTree.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_0_0
return gopurs_runtime.Apply(pkg_Test_RBTree.Get_act(), gopurs_runtime.Value{})
})
	})
	return main
}




