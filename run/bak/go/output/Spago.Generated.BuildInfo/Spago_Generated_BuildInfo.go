package Spago_Generated_BuildInfo

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_spagoVersion gopurs_runtime.Value
var once_spagoVersion sync.Once
func Get_spagoVersion() gopurs_runtime.Value {
	once_spagoVersion.Do(func() {
		cache_spagoVersion = gopurs_runtime.Str("1.0.4")
	})
	return cache_spagoVersion
}

var cache_pursVersion gopurs_runtime.Value
var once_pursVersion sync.Once
func Get_pursVersion() gopurs_runtime.Value {
	once_pursVersion.Do(func() {
		cache_pursVersion = gopurs_runtime.Str("0.15.16")
	})
	return cache_pursVersion
}

var cache_packages gopurs_runtime.Value
var once_packages sync.Once
func Get_packages() gopurs_runtime.Value {
	once_packages.Do(func() {
		cache_packages = gopurs_runtime.RecordDict1("ps-php-test", gopurs_runtime.Str("0.0.0"))
	})
	return cache_packages
}




