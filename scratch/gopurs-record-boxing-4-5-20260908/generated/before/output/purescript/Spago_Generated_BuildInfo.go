package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Spago_Generated_BuildInfo_spagoVersion gopurs_runtime.Value
var once_Spago_Generated_BuildInfo_spagoVersion sync.Once
func Get_Spago_Generated_BuildInfo_spagoVersion() gopurs_runtime.Value {
	once_Spago_Generated_BuildInfo_spagoVersion.Do(func() {
		cache_Spago_Generated_BuildInfo_spagoVersion = gopurs_runtime.Str("1.0.4")
	})
	return cache_Spago_Generated_BuildInfo_spagoVersion
}

var cache_Spago_Generated_BuildInfo_pursVersion gopurs_runtime.Value
var once_Spago_Generated_BuildInfo_pursVersion sync.Once
func Get_Spago_Generated_BuildInfo_pursVersion() gopurs_runtime.Value {
	once_Spago_Generated_BuildInfo_pursVersion.Do(func() {
		cache_Spago_Generated_BuildInfo_pursVersion = gopurs_runtime.Str("0.15.16")
	})
	return cache_Spago_Generated_BuildInfo_pursVersion
}

var cache_Spago_Generated_BuildInfo_packages gopurs_runtime.Value
var once_Spago_Generated_BuildInfo_packages sync.Once
func Get_Spago_Generated_BuildInfo_packages() gopurs_runtime.Value {
	once_Spago_Generated_BuildInfo_packages.Do(func() {
		cache_Spago_Generated_BuildInfo_packages = func() gopurs_runtime.Value {
				orig := struct{
	ps_minus_go_minus_test string
}{"0.0.0"}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"ps-go-test"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.ps_minus_go_minus_test)})
				}()
	})
	return cache_Spago_Generated_BuildInfo_packages
}




