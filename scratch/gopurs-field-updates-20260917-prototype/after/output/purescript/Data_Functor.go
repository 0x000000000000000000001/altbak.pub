package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Functor_dollar_Dict gopurs_runtime.Value
var once_Data_Functor_Functor_dollar_Dict sync.Once
func Get_Data_Functor_Functor_dollar_Dict() gopurs_runtime.Value {
	once_Data_Functor_Functor_dollar_Dict.Do(func() {
		cache_Data_Functor_Functor_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Functor_dollar_Dict(func() struct{
	go__map gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	go__map gopurs_runtime.Value
}{}
					clone.go__map = gopurs_runtime.RecordGet(orig, "map")
					return clone
				}()))}
})
	})
	return cache_Data_Functor_Functor_dollar_Dict
}

var cache_Data_Functor_go__map gopurs_runtime.Value
var once_Data_Functor_go__map sync.Once
func Get_Data_Functor_go__map() gopurs_runtime.Value {
	once_Data_Functor_go__map.Do(func() {
		cache_Data_Functor_go__map = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Functor_go__map
}

var cache_Data_Functor_map__151010626 gopurs_runtime.Value
var once_Data_Functor_map__151010626 sync.Once
func Get_Data_Functor_map__151010626() gopurs_runtime.Value {
	once_Data_Functor_map__151010626.Do(func() {
		cache_Data_Functor_map__151010626 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__151010626(__eta_norm_1_unused_0_box, Rebox_Data_Functor_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Functor_map__151010626
}

var cache_Data_Functor_map__1192782812 gopurs_runtime.Value
var once_Data_Functor_map__1192782812 sync.Once
func Get_Data_Functor_map__1192782812() gopurs_runtime.Value {
	once_Data_Functor_map__1192782812.Do(func() {
		cache_Data_Functor_map__1192782812 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1192782812(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Functor_map__1192782812
}

var cache_Data_Functor_map__3844050306 gopurs_runtime.Value
var once_Data_Functor_map__3844050306 sync.Once
func Get_Data_Functor_map__3844050306() gopurs_runtime.Value {
	once_Data_Functor_map__3844050306.Do(func() {
		cache_Data_Functor_map__3844050306 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__3844050306(__eta_norm_1_unused_0_box, Rebox_Data_Functor_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Functor_map__3844050306
}

var cache_Data_Functor_map__3996806813 gopurs_runtime.Value
var once_Data_Functor_map__3996806813 sync.Once
func Get_Data_Functor_map__3996806813() gopurs_runtime.Value {
	once_Data_Functor_map__3996806813.Do(func() {
		cache_Data_Functor_map__3996806813 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__3996806813(__eta_norm_1_0_box, Rebox_Data_Functor_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Functor_map__3996806813
}

var cache_Data_Functor_map__1141856137 gopurs_runtime.Value
var once_Data_Functor_map__1141856137 sync.Once
func Get_Data_Functor_map__1141856137() gopurs_runtime.Value {
	once_Data_Functor_map__1141856137.Do(func() {
		cache_Data_Functor_map__1141856137 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__1141856137(__eta_norm_1_unused_0_box, Rebox_Data_Functor_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Functor_map__1141856137
}

var cache_Data_Functor_map__3907495522 gopurs_runtime.Value
var once_Data_Functor_map__3907495522 sync.Once
func Get_Data_Functor_map__3907495522() gopurs_runtime.Value {
	once_Data_Functor_map__3907495522.Do(func() {
		cache_Data_Functor_map__3907495522 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__3907495522(__eta_norm_1_unused_0_box, Rebox_Data_Functor_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Functor_map__3907495522
}

var cache_Data_Functor_map__1365758973 gopurs_runtime.Value
var once_Data_Functor_map__1365758973 sync.Once
func Get_Data_Functor_map__1365758973() gopurs_runtime.Value {
	once_Data_Functor_map__1365758973.Do(func() {
		cache_Data_Functor_map__1365758973 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__1365758973(__eta_norm_1_0_box, Rebox_Data_Functor_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Functor_map__1365758973
}

var cache_Data_Functor_map__2983705738 gopurs_runtime.Value
var once_Data_Functor_map__2983705738 sync.Once
func Get_Data_Functor_map__2983705738() gopurs_runtime.Value {
	once_Data_Functor_map__2983705738.Do(func() {
		cache_Data_Functor_map__2983705738 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2983705738(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__2983705738
}

var cache_Data_Functor_map__1969116699 gopurs_runtime.Value
var once_Data_Functor_map__1969116699 sync.Once
func Get_Data_Functor_map__1969116699() gopurs_runtime.Value {
	once_Data_Functor_map__1969116699.Do(func() {
		cache_Data_Functor_map__1969116699 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1969116699(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__1969116699
}

var cache_Data_Functor_map__3293112731 gopurs_runtime.Value
var once_Data_Functor_map__3293112731 sync.Once
func Get_Data_Functor_map__3293112731() gopurs_runtime.Value {
	once_Data_Functor_map__3293112731.Do(func() {
		cache_Data_Functor_map__3293112731 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3293112731(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__3293112731
}

var cache_Data_Functor_map__4431069 gopurs_runtime.Value
var once_Data_Functor_map__4431069 sync.Once
func Get_Data_Functor_map__4431069() gopurs_runtime.Value {
	once_Data_Functor_map__4431069.Do(func() {
		cache_Data_Functor_map__4431069 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Functor_map__4431069(__eta_norm_1_0_box, Rebox_Data_Functor_3094389156_1387998409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Functor_map__4431069
}

var cache_Data_Functor_map__2404595521 gopurs_runtime.Value
var once_Data_Functor_map__2404595521 sync.Once
func Get_Data_Functor_map__2404595521() gopurs_runtime.Value {
	once_Data_Functor_map__2404595521.Do(func() {
		cache_Data_Functor_map__2404595521 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2404595521(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__2404595521
}

var cache_Data_Functor_map__1563461719 gopurs_runtime.Value
var once_Data_Functor_map__1563461719 sync.Once
func Get_Data_Functor_map__1563461719() gopurs_runtime.Value {
	once_Data_Functor_map__1563461719.Do(func() {
		cache_Data_Functor_map__1563461719 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1563461719(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__1563461719
}

var cache_Data_Functor_map__1555135774 gopurs_runtime.Value
var once_Data_Functor_map__1555135774 sync.Once
func Get_Data_Functor_map__1555135774() gopurs_runtime.Value {
	once_Data_Functor_map__1555135774.Do(func() {
		cache_Data_Functor_map__1555135774 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1555135774(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__1555135774
}

var cache_Data_Functor_map__2588336262 gopurs_runtime.Value
var once_Data_Functor_map__2588336262 sync.Once
func Get_Data_Functor_map__2588336262() gopurs_runtime.Value {
	once_Data_Functor_map__2588336262.Do(func() {
		cache_Data_Functor_map__2588336262 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2588336262(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__2588336262
}

var cache_Data_Functor_map__1205579984 gopurs_runtime.Value
var once_Data_Functor_map__1205579984 sync.Once
func Get_Data_Functor_map__1205579984() gopurs_runtime.Value {
	once_Data_Functor_map__1205579984.Do(func() {
		cache_Data_Functor_map__1205579984 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1205579984(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__1205579984
}

var cache_Data_Functor_map__296345017 gopurs_runtime.Value
var once_Data_Functor_map__296345017 sync.Once
func Get_Data_Functor_map__296345017() gopurs_runtime.Value {
	once_Data_Functor_map__296345017.Do(func() {
		cache_Data_Functor_map__296345017 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__296345017(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__296345017
}

var cache_Data_Functor_map__2235814739 gopurs_runtime.Value
var once_Data_Functor_map__2235814739 sync.Once
func Get_Data_Functor_map__2235814739() gopurs_runtime.Value {
	once_Data_Functor_map__2235814739.Do(func() {
		cache_Data_Functor_map__2235814739 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2235814739(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__2235814739
}

var cache_Data_Functor_map__1877032517 gopurs_runtime.Value
var once_Data_Functor_map__1877032517 sync.Once
func Get_Data_Functor_map__1877032517() gopurs_runtime.Value {
	once_Data_Functor_map__1877032517.Do(func() {
		cache_Data_Functor_map__1877032517 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1877032517(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__1877032517
}

var cache_Data_Functor_map__2061311692 gopurs_runtime.Value
var once_Data_Functor_map__2061311692 sync.Once
func Get_Data_Functor_map__2061311692() gopurs_runtime.Value {
	once_Data_Functor_map__2061311692.Do(func() {
		cache_Data_Functor_map__2061311692 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2061311692(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__2061311692
}

var cache_Data_Functor_map__1572513044 gopurs_runtime.Value
var once_Data_Functor_map__1572513044 sync.Once
func Get_Data_Functor_map__1572513044() gopurs_runtime.Value {
	once_Data_Functor_map__1572513044.Do(func() {
		cache_Data_Functor_map__1572513044 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1572513044(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__1572513044
}

var cache_Data_Functor_map__4036895682 gopurs_runtime.Value
var once_Data_Functor_map__4036895682 sync.Once
func Get_Data_Functor_map__4036895682() gopurs_runtime.Value {
	once_Data_Functor_map__4036895682.Do(func() {
		cache_Data_Functor_map__4036895682 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4036895682(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__4036895682
}

var cache_Data_Functor_map__683888683 gopurs_runtime.Value
var once_Data_Functor_map__683888683 sync.Once
func Get_Data_Functor_map__683888683() gopurs_runtime.Value {
	once_Data_Functor_map__683888683.Do(func() {
		cache_Data_Functor_map__683888683 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__683888683(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__683888683
}

var cache_Data_Functor_map__1002236606 gopurs_runtime.Value
var once_Data_Functor_map__1002236606 sync.Once
func Get_Data_Functor_map__1002236606() gopurs_runtime.Value {
	once_Data_Functor_map__1002236606.Do(func() {
		cache_Data_Functor_map__1002236606 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1002236606(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__1002236606
}

var cache_Data_Functor_map__3097389096 gopurs_runtime.Value
var once_Data_Functor_map__3097389096 sync.Once
func Get_Data_Functor_map__3097389096() gopurs_runtime.Value {
	once_Data_Functor_map__3097389096.Do(func() {
		cache_Data_Functor_map__3097389096 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3097389096(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__3097389096
}

var cache_Data_Functor_map__999202337 gopurs_runtime.Value
var once_Data_Functor_map__999202337 sync.Once
func Get_Data_Functor_map__999202337() gopurs_runtime.Value {
	once_Data_Functor_map__999202337.Do(func() {
		cache_Data_Functor_map__999202337 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__999202337(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__999202337
}

var cache_Data_Functor_map__395294283 gopurs_runtime.Value
var once_Data_Functor_map__395294283 sync.Once
func Get_Data_Functor_map__395294283() gopurs_runtime.Value {
	once_Data_Functor_map__395294283.Do(func() {
		cache_Data_Functor_map__395294283 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__395294283(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__395294283
}

var cache_Data_Functor_map__3468671773 gopurs_runtime.Value
var once_Data_Functor_map__3468671773 sync.Once
func Get_Data_Functor_map__3468671773() gopurs_runtime.Value {
	once_Data_Functor_map__3468671773.Do(func() {
		cache_Data_Functor_map__3468671773 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3468671773(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__3468671773
}

var cache_Data_Functor_map__1648975668 gopurs_runtime.Value
var once_Data_Functor_map__1648975668 sync.Once
func Get_Data_Functor_map__1648975668() gopurs_runtime.Value {
	once_Data_Functor_map__1648975668.Do(func() {
		cache_Data_Functor_map__1648975668 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1648975668(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__1648975668
}

var cache_Data_Functor_map__102065405 gopurs_runtime.Value
var once_Data_Functor_map__102065405 sync.Once
func Get_Data_Functor_map__102065405() gopurs_runtime.Value {
	once_Data_Functor_map__102065405.Do(func() {
		cache_Data_Functor_map__102065405 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__102065405(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__102065405
}

var cache_Data_Functor_map__794090731 gopurs_runtime.Value
var once_Data_Functor_map__794090731 sync.Once
func Get_Data_Functor_map__794090731() gopurs_runtime.Value {
	once_Data_Functor_map__794090731.Do(func() {
		cache_Data_Functor_map__794090731 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__794090731(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__794090731
}

var cache_Data_Functor_map__484586722 gopurs_runtime.Value
var once_Data_Functor_map__484586722 sync.Once
func Get_Data_Functor_map__484586722() gopurs_runtime.Value {
	once_Data_Functor_map__484586722.Do(func() {
		cache_Data_Functor_map__484586722 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__484586722(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__484586722
}

var cache_Data_Functor_map__478722022 gopurs_runtime.Value
var once_Data_Functor_map__478722022 sync.Once
func Get_Data_Functor_map__478722022() gopurs_runtime.Value {
	once_Data_Functor_map__478722022.Do(func() {
		cache_Data_Functor_map__478722022 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__478722022(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__478722022
}

var cache_Data_Functor_map__105413808 gopurs_runtime.Value
var once_Data_Functor_map__105413808 sync.Once
func Get_Data_Functor_map__105413808() gopurs_runtime.Value {
	once_Data_Functor_map__105413808.Do(func() {
		cache_Data_Functor_map__105413808 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__105413808(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__105413808
}

var cache_Data_Functor_map__3791346649 gopurs_runtime.Value
var once_Data_Functor_map__3791346649 sync.Once
func Get_Data_Functor_map__3791346649() gopurs_runtime.Value {
	once_Data_Functor_map__3791346649.Do(func() {
		cache_Data_Functor_map__3791346649 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3791346649(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__3791346649
}

var cache_Data_Functor_map__4079021475 gopurs_runtime.Value
var once_Data_Functor_map__4079021475 sync.Once
func Get_Data_Functor_map__4079021475() gopurs_runtime.Value {
	once_Data_Functor_map__4079021475.Do(func() {
		cache_Data_Functor_map__4079021475 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4079021475(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__4079021475
}

var cache_Data_Functor_map__1223398517 gopurs_runtime.Value
var once_Data_Functor_map__1223398517 sync.Once
func Get_Data_Functor_map__1223398517() gopurs_runtime.Value {
	once_Data_Functor_map__1223398517.Do(func() {
		cache_Data_Functor_map__1223398517 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1223398517(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__1223398517
}

var cache_Data_Functor_map__2625193884 gopurs_runtime.Value
var once_Data_Functor_map__2625193884 sync.Once
func Get_Data_Functor_map__2625193884() gopurs_runtime.Value {
	once_Data_Functor_map__2625193884.Do(func() {
		cache_Data_Functor_map__2625193884 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2625193884(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__2625193884
}

var cache_Data_Functor_map__2185106885 gopurs_runtime.Value
var once_Data_Functor_map__2185106885 sync.Once
func Get_Data_Functor_map__2185106885() gopurs_runtime.Value {
	once_Data_Functor_map__2185106885.Do(func() {
		cache_Data_Functor_map__2185106885 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2185106885(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__2185106885
}

var cache_Data_Functor_map__293830419 gopurs_runtime.Value
var once_Data_Functor_map__293830419 sync.Once
func Get_Data_Functor_map__293830419() gopurs_runtime.Value {
	once_Data_Functor_map__293830419.Do(func() {
		cache_Data_Functor_map__293830419 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__293830419(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__293830419
}

var cache_Data_Functor_map__982714682 gopurs_runtime.Value
var once_Data_Functor_map__982714682 sync.Once
func Get_Data_Functor_map__982714682() gopurs_runtime.Value {
	once_Data_Functor_map__982714682.Do(func() {
		cache_Data_Functor_map__982714682 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__982714682(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__982714682
}

var cache_Data_Functor_map__4131672173 gopurs_runtime.Value
var once_Data_Functor_map__4131672173 sync.Once
func Get_Data_Functor_map__4131672173() gopurs_runtime.Value {
	once_Data_Functor_map__4131672173.Do(func() {
		cache_Data_Functor_map__4131672173 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4131672173(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__4131672173
}

var cache_Data_Functor_map__4187694843 gopurs_runtime.Value
var once_Data_Functor_map__4187694843 sync.Once
func Get_Data_Functor_map__4187694843() gopurs_runtime.Value {
	once_Data_Functor_map__4187694843.Do(func() {
		cache_Data_Functor_map__4187694843 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4187694843(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__4187694843
}

var cache_Data_Functor_map__4013754482 gopurs_runtime.Value
var once_Data_Functor_map__4013754482 sync.Once
func Get_Data_Functor_map__4013754482() gopurs_runtime.Value {
	once_Data_Functor_map__4013754482.Do(func() {
		cache_Data_Functor_map__4013754482 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4013754482(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__4013754482
}

var cache_Data_Functor_map__1723234304 gopurs_runtime.Value
var once_Data_Functor_map__1723234304 sync.Once
func Get_Data_Functor_map__1723234304() gopurs_runtime.Value {
	once_Data_Functor_map__1723234304.Do(func() {
		cache_Data_Functor_map__1723234304 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__1723234304(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__1723234304
}

var cache_Data_Functor_map__4039752726 gopurs_runtime.Value
var once_Data_Functor_map__4039752726 sync.Once
func Get_Data_Functor_map__4039752726() gopurs_runtime.Value {
	once_Data_Functor_map__4039752726.Do(func() {
		cache_Data_Functor_map__4039752726 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__4039752726(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__4039752726
}

var cache_Data_Functor_map__2414686431 gopurs_runtime.Value
var once_Data_Functor_map__2414686431 sync.Once
func Get_Data_Functor_map__2414686431() gopurs_runtime.Value {
	once_Data_Functor_map__2414686431.Do(func() {
		cache_Data_Functor_map__2414686431 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2414686431(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__2414686431
}

var cache_Data_Functor_map__3441199853 gopurs_runtime.Value
var once_Data_Functor_map__3441199853 sync.Once
func Get_Data_Functor_map__3441199853() gopurs_runtime.Value {
	once_Data_Functor_map__3441199853.Do(func() {
		cache_Data_Functor_map__3441199853 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3441199853(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__3441199853
}

var cache_Data_Functor_map__2477564603 gopurs_runtime.Value
var once_Data_Functor_map__2477564603 sync.Once
func Get_Data_Functor_map__2477564603() gopurs_runtime.Value {
	once_Data_Functor_map__2477564603.Do(func() {
		cache_Data_Functor_map__2477564603 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2477564603(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__2477564603
}

var cache_Data_Functor_map__196852562 gopurs_runtime.Value
var once_Data_Functor_map__196852562 sync.Once
func Get_Data_Functor_map__196852562() gopurs_runtime.Value {
	once_Data_Functor_map__196852562.Do(func() {
		cache_Data_Functor_map__196852562 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__196852562(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__196852562
}

var cache_Data_Functor_map__2501627275 gopurs_runtime.Value
var once_Data_Functor_map__2501627275 sync.Once
func Get_Data_Functor_map__2501627275() gopurs_runtime.Value {
	once_Data_Functor_map__2501627275.Do(func() {
		cache_Data_Functor_map__2501627275 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2501627275(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__2501627275
}

var cache_Data_Functor_map__3939840413 gopurs_runtime.Value
var once_Data_Functor_map__3939840413 sync.Once
func Get_Data_Functor_map__3939840413() gopurs_runtime.Value {
	once_Data_Functor_map__3939840413.Do(func() {
		cache_Data_Functor_map__3939840413 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__3939840413(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__3939840413
}

var cache_Data_Functor_map__2061195028 gopurs_runtime.Value
var once_Data_Functor_map__2061195028 sync.Once
func Get_Data_Functor_map__2061195028() gopurs_runtime.Value {
	once_Data_Functor_map__2061195028.Do(func() {
		cache_Data_Functor_map__2061195028 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_map__2061195028(__eta_norm_1_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Functor_map__2061195028
}

var cache_Data_Functor_mapFlipped gopurs_runtime.Value
var once_Data_Functor_mapFlipped sync.Once
func Get_Data_Functor_mapFlipped() gopurs_runtime.Value {
	once_Data_Functor_mapFlipped.Do(func() {
		cache_Data_Functor_mapFlipped = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_mapFlipped(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_Data_Functor_mapFlipped
}

var cache_Data_Functor_void gopurs_runtime.Value
var once_Data_Functor_void sync.Once
func Get_Data_Functor_void() gopurs_runtime.Value {
	once_Data_Functor_void.Do(func() {
		cache_Data_Functor_void = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box))
})
	})
	return cache_Data_Functor_void
}

var cache_Data_Functor_void__3393626650 gopurs_runtime.Value
var once_Data_Functor_void__3393626650 sync.Once
func Get_Data_Functor_void__3393626650() gopurs_runtime.Value {
	once_Data_Functor_void__3393626650.Do(func() {
		cache_Data_Functor_void__3393626650 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__3393626650(__eta_norm_0_0_box)
})
	})
	return cache_Data_Functor_void__3393626650
}

var cache_Data_Functor_void__990419905 gopurs_runtime.Value
var once_Data_Functor_void__990419905 sync.Once
func Get_Data_Functor_void__990419905() gopurs_runtime.Value {
	once_Data_Functor_void__990419905.Do(func() {
		cache_Data_Functor_void__990419905 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__990419905(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__990419905
}

var cache_Data_Functor_void__37931671 gopurs_runtime.Value
var once_Data_Functor_void__37931671 sync.Once
func Get_Data_Functor_void__37931671() gopurs_runtime.Value {
	once_Data_Functor_void__37931671.Do(func() {
		cache_Data_Functor_void__37931671 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__37931671(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__37931671
}

var cache_Data_Functor_void__2259642750 gopurs_runtime.Value
var once_Data_Functor_void__2259642750 sync.Once
func Get_Data_Functor_void__2259642750() gopurs_runtime.Value {
	once_Data_Functor_void__2259642750.Do(func() {
		cache_Data_Functor_void__2259642750 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__2259642750(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__2259642750
}

var cache_Data_Functor_void__4174605798 gopurs_runtime.Value
var once_Data_Functor_void__4174605798 sync.Once
func Get_Data_Functor_void__4174605798() gopurs_runtime.Value {
	once_Data_Functor_void__4174605798.Do(func() {
		cache_Data_Functor_void__4174605798 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__4174605798(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__4174605798
}

var cache_Data_Functor_void__1062770544 gopurs_runtime.Value
var once_Data_Functor_void__1062770544 sync.Once
func Get_Data_Functor_void__1062770544() gopurs_runtime.Value {
	once_Data_Functor_void__1062770544.Do(func() {
		cache_Data_Functor_void__1062770544 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__1062770544(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__1062770544
}

var cache_Data_Functor_void__2350790585 gopurs_runtime.Value
var once_Data_Functor_void__2350790585 sync.Once
func Get_Data_Functor_void__2350790585() gopurs_runtime.Value {
	once_Data_Functor_void__2350790585.Do(func() {
		cache_Data_Functor_void__2350790585 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__2350790585(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__2350790585
}

var cache_Data_Functor_void__220841491 gopurs_runtime.Value
var once_Data_Functor_void__220841491 sync.Once
func Get_Data_Functor_void__220841491() gopurs_runtime.Value {
	once_Data_Functor_void__220841491.Do(func() {
		cache_Data_Functor_void__220841491 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__220841491(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__220841491
}

var cache_Data_Functor_void__3758797381 gopurs_runtime.Value
var once_Data_Functor_void__3758797381 sync.Once
func Get_Data_Functor_void__3758797381() gopurs_runtime.Value {
	once_Data_Functor_void__3758797381.Do(func() {
		cache_Data_Functor_void__3758797381 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__3758797381(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__3758797381
}

var cache_Data_Functor_void__3813769196 gopurs_runtime.Value
var once_Data_Functor_void__3813769196 sync.Once
func Get_Data_Functor_void__3813769196() gopurs_runtime.Value {
	once_Data_Functor_void__3813769196.Do(func() {
		cache_Data_Functor_void__3813769196 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__3813769196(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__3813769196
}

var cache_Data_Functor_void__3599204660 gopurs_runtime.Value
var once_Data_Functor_void__3599204660 sync.Once
func Get_Data_Functor_void__3599204660() gopurs_runtime.Value {
	once_Data_Functor_void__3599204660.Do(func() {
		cache_Data_Functor_void__3599204660 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__3599204660(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__3599204660
}

var cache_Data_Functor_void__1324701602 gopurs_runtime.Value
var once_Data_Functor_void__1324701602 sync.Once
func Get_Data_Functor_void__1324701602() gopurs_runtime.Value {
	once_Data_Functor_void__1324701602.Do(func() {
		cache_Data_Functor_void__1324701602 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__1324701602(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__1324701602
}

var cache_Data_Functor_void__715030123 gopurs_runtime.Value
var once_Data_Functor_void__715030123 sync.Once
func Get_Data_Functor_void__715030123() gopurs_runtime.Value {
	once_Data_Functor_void__715030123.Do(func() {
		cache_Data_Functor_void__715030123 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__715030123(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__715030123
}

var cache_Data_Functor_void__380936478 gopurs_runtime.Value
var once_Data_Functor_void__380936478 sync.Once
func Get_Data_Functor_void__380936478() gopurs_runtime.Value {
	once_Data_Functor_void__380936478.Do(func() {
		cache_Data_Functor_void__380936478 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__380936478(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__380936478
}

var cache_Data_Functor_void__380628936 gopurs_runtime.Value
var once_Data_Functor_void__380628936 sync.Once
func Get_Data_Functor_void__380628936() gopurs_runtime.Value {
	once_Data_Functor_void__380628936.Do(func() {
		cache_Data_Functor_void__380628936 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__380628936(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__380628936
}

var cache_Data_Functor_void__2589128993 gopurs_runtime.Value
var once_Data_Functor_void__2589128993 sync.Once
func Get_Data_Functor_void__2589128993() gopurs_runtime.Value {
	once_Data_Functor_void__2589128993.Do(func() {
		cache_Data_Functor_void__2589128993 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__2589128993(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__2589128993
}

var cache_Data_Functor_void__4146435083 gopurs_runtime.Value
var once_Data_Functor_void__4146435083 sync.Once
func Get_Data_Functor_void__4146435083() gopurs_runtime.Value {
	once_Data_Functor_void__4146435083.Do(func() {
		cache_Data_Functor_void__4146435083 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__4146435083(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__4146435083
}

var cache_Data_Functor_void__1659429917 gopurs_runtime.Value
var once_Data_Functor_void__1659429917 sync.Once
func Get_Data_Functor_void__1659429917() gopurs_runtime.Value {
	once_Data_Functor_void__1659429917.Do(func() {
		cache_Data_Functor_void__1659429917 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__1659429917(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__1659429917
}

var cache_Data_Functor_void__2038934356 gopurs_runtime.Value
var once_Data_Functor_void__2038934356 sync.Once
func Get_Data_Functor_void__2038934356() gopurs_runtime.Value {
	once_Data_Functor_void__2038934356.Do(func() {
		cache_Data_Functor_void__2038934356 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__2038934356(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__2038934356
}

var cache_Data_Functor_void__234143613 gopurs_runtime.Value
var once_Data_Functor_void__234143613 sync.Once
func Get_Data_Functor_void__234143613() gopurs_runtime.Value {
	once_Data_Functor_void__234143613.Do(func() {
		cache_Data_Functor_void__234143613 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__234143613(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__234143613
}

var cache_Data_Functor_void__2094920235 gopurs_runtime.Value
var once_Data_Functor_void__2094920235 sync.Once
func Get_Data_Functor_void__2094920235() gopurs_runtime.Value {
	once_Data_Functor_void__2094920235.Do(func() {
		cache_Data_Functor_void__2094920235 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__2094920235(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__2094920235
}

var cache_Data_Functor_void__4269283266 gopurs_runtime.Value
var once_Data_Functor_void__4269283266 sync.Once
func Get_Data_Functor_void__4269283266() gopurs_runtime.Value {
	once_Data_Functor_void__4269283266.Do(func() {
		cache_Data_Functor_void__4269283266 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__4269283266(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__4269283266
}

var cache_Data_Functor_void__3317639750 gopurs_runtime.Value
var once_Data_Functor_void__3317639750 sync.Once
func Get_Data_Functor_void__3317639750() gopurs_runtime.Value {
	once_Data_Functor_void__3317639750.Do(func() {
		cache_Data_Functor_void__3317639750 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__3317639750(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__3317639750
}

var cache_Data_Functor_void__1880795472 gopurs_runtime.Value
var once_Data_Functor_void__1880795472 sync.Once
func Get_Data_Functor_void__1880795472() gopurs_runtime.Value {
	once_Data_Functor_void__1880795472.Do(func() {
		cache_Data_Functor_void__1880795472 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__1880795472(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__1880795472
}

var cache_Data_Functor_void__460922713 gopurs_runtime.Value
var once_Data_Functor_void__460922713 sync.Once
func Get_Data_Functor_void__460922713() gopurs_runtime.Value {
	once_Data_Functor_void__460922713.Do(func() {
		cache_Data_Functor_void__460922713 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__460922713(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__460922713
}

var cache_Data_Functor_void__3340832483 gopurs_runtime.Value
var once_Data_Functor_void__3340832483 sync.Once
func Get_Data_Functor_void__3340832483() gopurs_runtime.Value {
	once_Data_Functor_void__3340832483.Do(func() {
		cache_Data_Functor_void__3340832483 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__3340832483(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__3340832483
}

var cache_Data_Functor_void__982491637 gopurs_runtime.Value
var once_Data_Functor_void__982491637 sync.Once
func Get_Data_Functor_void__982491637() gopurs_runtime.Value {
	once_Data_Functor_void__982491637.Do(func() {
		cache_Data_Functor_void__982491637 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__982491637(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__982491637
}

var cache_Data_Functor_void__2681811132 gopurs_runtime.Value
var once_Data_Functor_void__2681811132 sync.Once
func Get_Data_Functor_void__2681811132() gopurs_runtime.Value {
	once_Data_Functor_void__2681811132.Do(func() {
		cache_Data_Functor_void__2681811132 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__2681811132(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__2681811132
}

var cache_Data_Functor_void__2380011973 gopurs_runtime.Value
var once_Data_Functor_void__2380011973 sync.Once
func Get_Data_Functor_void__2380011973() gopurs_runtime.Value {
	once_Data_Functor_void__2380011973.Do(func() {
		cache_Data_Functor_void__2380011973 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__2380011973(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__2380011973
}

var cache_Data_Functor_void__2241417683 gopurs_runtime.Value
var once_Data_Functor_void__2241417683 sync.Once
func Get_Data_Functor_void__2241417683() gopurs_runtime.Value {
	once_Data_Functor_void__2241417683.Do(func() {
		cache_Data_Functor_void__2241417683 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__2241417683(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__2241417683
}

var cache_Data_Functor_void__3536814106 gopurs_runtime.Value
var once_Data_Functor_void__3536814106 sync.Once
func Get_Data_Functor_void__3536814106() gopurs_runtime.Value {
	once_Data_Functor_void__3536814106.Do(func() {
		cache_Data_Functor_void__3536814106 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__3536814106(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__3536814106
}

var cache_Data_Functor_void__1927686381 gopurs_runtime.Value
var once_Data_Functor_void__1927686381 sync.Once
func Get_Data_Functor_void__1927686381() gopurs_runtime.Value {
	once_Data_Functor_void__1927686381.Do(func() {
		cache_Data_Functor_void__1927686381 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__1927686381(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__1927686381
}

var cache_Data_Functor_void__2080935483 gopurs_runtime.Value
var once_Data_Functor_void__2080935483 sync.Once
func Get_Data_Functor_void__2080935483() gopurs_runtime.Value {
	once_Data_Functor_void__2080935483.Do(func() {
		cache_Data_Functor_void__2080935483 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__2080935483(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__2080935483
}

var cache_Data_Functor_void__2449968466 gopurs_runtime.Value
var once_Data_Functor_void__2449968466 sync.Once
func Get_Data_Functor_void__2449968466() gopurs_runtime.Value {
	once_Data_Functor_void__2449968466.Do(func() {
		cache_Data_Functor_void__2449968466 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__2449968466(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__2449968466
}

var cache_Data_Functor_void__143291040 gopurs_runtime.Value
var once_Data_Functor_void__143291040 sync.Once
func Get_Data_Functor_void__143291040() gopurs_runtime.Value {
	once_Data_Functor_void__143291040.Do(func() {
		cache_Data_Functor_void__143291040 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__143291040(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__143291040
}

var cache_Data_Functor_void__672263030 gopurs_runtime.Value
var once_Data_Functor_void__672263030 sync.Once
func Get_Data_Functor_void__672263030() gopurs_runtime.Value {
	once_Data_Functor_void__672263030.Do(func() {
		cache_Data_Functor_void__672263030 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__672263030(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__672263030
}

var cache_Data_Functor_void__355975199 gopurs_runtime.Value
var once_Data_Functor_void__355975199 sync.Once
func Get_Data_Functor_void__355975199() gopurs_runtime.Value {
	once_Data_Functor_void__355975199.Do(func() {
		cache_Data_Functor_void__355975199 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__355975199(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__355975199
}

var cache_Data_Functor_void__1926861677 gopurs_runtime.Value
var once_Data_Functor_void__1926861677 sync.Once
func Get_Data_Functor_void__1926861677() gopurs_runtime.Value {
	once_Data_Functor_void__1926861677.Do(func() {
		cache_Data_Functor_void__1926861677 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__1926861677(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__1926861677
}

var cache_Data_Functor_void__4242248443 gopurs_runtime.Value
var once_Data_Functor_void__4242248443 sync.Once
func Get_Data_Functor_void__4242248443() gopurs_runtime.Value {
	once_Data_Functor_void__4242248443.Do(func() {
		cache_Data_Functor_void__4242248443 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__4242248443(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__4242248443
}

var cache_Data_Functor_void__3690139954 gopurs_runtime.Value
var once_Data_Functor_void__3690139954 sync.Once
func Get_Data_Functor_void__3690139954() gopurs_runtime.Value {
	once_Data_Functor_void__3690139954.Do(func() {
		cache_Data_Functor_void__3690139954 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__3690139954(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__3690139954
}

var cache_Data_Functor_void__2838993739 gopurs_runtime.Value
var once_Data_Functor_void__2838993739 sync.Once
func Get_Data_Functor_void__2838993739() gopurs_runtime.Value {
	once_Data_Functor_void__2838993739.Do(func() {
		cache_Data_Functor_void__2838993739 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__2838993739(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__2838993739
}

var cache_Data_Functor_void__929958557 gopurs_runtime.Value
var once_Data_Functor_void__929958557 sync.Once
func Get_Data_Functor_void__929958557() gopurs_runtime.Value {
	once_Data_Functor_void__929958557.Do(func() {
		cache_Data_Functor_void__929958557 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__929958557(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__929958557
}

var cache_Data_Functor_void__1885598516 gopurs_runtime.Value
var once_Data_Functor_void__1885598516 sync.Once
func Get_Data_Functor_void__1885598516() gopurs_runtime.Value {
	once_Data_Functor_void__1885598516.Do(func() {
		cache_Data_Functor_void__1885598516 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__1885598516(__eta_norm_0_unused_0_box)
})
	})
	return cache_Data_Functor_void__1885598516
}

var cache_Data_Functor_void__565086979 gopurs_runtime.Value
var once_Data_Functor_void__565086979 sync.Once
func Get_Data_Functor_void__565086979() gopurs_runtime.Value {
	once_Data_Functor_void__565086979.Do(func() {
		cache_Data_Functor_void__565086979 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_void__565086979(__eta_norm_0_0_box)
})
	})
	return cache_Data_Functor_void__565086979
}

var cache_Data_Functor_voidLeft gopurs_runtime.Value
var once_Data_Functor_voidLeft sync.Once
func Get_Data_Functor_voidLeft() gopurs_runtime.Value {
	once_Data_Functor_voidLeft.Do(func() {
		cache_Data_Functor_voidLeft = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_voidLeft(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, x_2_box)
})
	})
	return cache_Data_Functor_voidLeft
}

var cache_Data_Functor_functorArray gopurs_runtime.Value
var once_Data_Functor_functorArray sync.Once
func Get_Data_Functor_functorArray() gopurs_runtime.Value {
	once_Data_Functor_functorArray.Do(func() {
		cache_Data_Functor_functorArray = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, Get_Data_Functor_arrayMap()}))}
	})
	return cache_Data_Functor_functorArray
}

var cache_Data_Functor_voidLeft__1381072619 gopurs_runtime.Value
var once_Data_Functor_voidLeft__1381072619 sync.Once
func Get_Data_Functor_voidLeft__1381072619() gopurs_runtime.Value {
	once_Data_Functor_voidLeft__1381072619.Do(func() {
		cache_Data_Functor_voidLeft__1381072619 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_voidLeft__1381072619(f_0_box, x_1_box.StrVal())
})
	})
	return cache_Data_Functor_voidLeft__1381072619
}

var cache_Data_Functor_voidRight gopurs_runtime.Value
var once_Data_Functor_voidRight sync.Once
func Get_Data_Functor_voidRight() gopurs_runtime.Value {
	once_Data_Functor_voidRight.Do(func() {
		cache_Data_Functor_voidRight = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_voidRight(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), x_1_box)
})
	})
	return cache_Data_Functor_voidRight
}

var cache_Data_Functor_functorProxy gopurs_runtime.Value
var once_Data_Functor_functorProxy sync.Once
func Get_Data_Functor_functorProxy() gopurs_runtime.Value {
	once_Data_Functor_functorProxy.Do(func() {
		cache_Data_Functor_functorProxy = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_562069347_2812149806((&Constructor_Data_Functor_Functor[uint32]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Functor_functorProxy
}

var cache_Data_Functor_functorFn gopurs_runtime.Value
var once_Data_Functor_functorFn sync.Once
func Get_Data_Functor_functorFn() gopurs_runtime.Value {
	once_Data_Functor_functorFn.Do(func() {
		cache_Data_Functor_functorFn = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn()))}))}
	})
	return cache_Data_Functor_functorFn
}

var cache_Data_Functor_flap gopurs_runtime.Value
var once_Data_Functor_flap sync.Once
func Get_Data_Functor_flap() gopurs_runtime.Value {
	once_Data_Functor_flap.Do(func() {
		cache_Data_Functor_flap = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, ff_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_flap(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), ff_1_box, x_2_box)
})
	})
	return cache_Data_Functor_flap
}

type Constructor_Data_Functor_Functor[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[929368378] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Functor_Functor[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "map": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Functor_Functor: " + key)
		}
	}
}


func Call_Data_Functor_Functor_dollar_Dict(x_0_loop struct{
	go__map gopurs_runtime.Value
}) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
var x_0 struct{
	go__map gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict1("map", orig.go__map)
				}())
}

func Call_Data_Functor_go__map(dict_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_Data_Functor_map__151010626(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) struct{V0 gopurs_runtime.Value; V1 bool} {
map__151010626:
for {
if false { continue map__151010626 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, Get_Data_DateTime_DateTime(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_2280409795_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Functor_map__1192782812(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1192782812:
for {
if false { continue map__1192782812 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, Get_Effect_Aff_effectCanceler(), __eta_norm_0_1)
}
}

func Call_Data_Functor_map__3844050306(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[int64]) struct{V0 gopurs_runtime.Value; V1 bool} {
map__3844050306:
for {
if false { continue map__3844050306 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, Get_Data_Date_Date(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1170268447_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Functor_map__3996806813(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[int64]) struct{V0 gopurs_runtime.Value; V1 bool} {
map__3996806813:
for {
if false { continue map__3996806813 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, __eta_norm_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1170268447_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Functor_map__1141856137(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[int64]) struct{V0 gopurs_runtime.Value; V1 bool} {
map__1141856137:
for {
if false { continue map__1141856137 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, Get_Data_Date_exactDate(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1170268447_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Functor_map__3907495522(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[int64]) struct{V0 gopurs_runtime.Value; V1 bool} {
map__3907495522:
for {
if false { continue map__3907495522 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, Get_Data_Time_Time(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1170268447_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Functor_map__1365758973(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[int64]) struct{V0 gopurs_runtime.Value; V1 bool} {
map__1365758973:
for {
if false { continue map__1365758973 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, __eta_norm_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1170268447_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Functor_map__2983705738(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2983705738:
for {
if false { continue map__2983705738 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, Get_Data_DateTime_Instant_toDateTime(), Get_Effect_Now_now())
}
}

func Call_Data_Functor_map__1969116699(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1969116699:
for {
if false { continue map__1969116699 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Effect_Now_now())
}
}

func Call_Data_Functor_map__3293112731(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__3293112731:
for {
if false { continue map__3293112731 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Effect_Now_now())
}
}

func Call_Data_Functor_map__4431069(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[struct{
	head int64
	tail string
}]) struct{V0 gopurs_runtime.Value; V1 bool} {
map__4431069:
for {
if false { continue map__4431069 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[struct{
	head int64
	tail string
}] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Rebox_Data_Functor_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe())).V0, __eta_norm_1_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Functor_1387998409_3094389156(__eta_norm_0_1))})
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Functor_map__2404595521(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2404595521:
for {
if false { continue map__2404595521 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_AckermannFFICheatcode_act())
}
}

func Call_Data_Functor_map__1563461719(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1563461719:
for {
if false { continue map__1563461719 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_AckermannFFI_act())
}
}

func Call_Data_Functor_map__1555135774(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1555135774:
for {
if false { continue map__1555135774 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_Ackermann_act())
}
}

func Call_Data_Functor_map__2588336262(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2588336262:
for {
if false { continue map__2588336262 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_ArrayOpsFFICheatcode_act())
}
}

func Call_Data_Functor_map__1205579984(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1205579984:
for {
if false { continue map__1205579984 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_ArrayOpsFFI_act())
}
}

func Call_Data_Functor_map__296345017(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__296345017:
for {
if false { continue map__296345017 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_ArrayOps_act())
}
}

func Call_Data_Functor_map__2235814739(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2235814739:
for {
if false { continue map__2235814739 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_AstTreeFFICheatcode_act())
}
}

func Call_Data_Functor_map__1877032517(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1877032517:
for {
if false { continue map__1877032517 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_AstTreeFFI_act())
}
}

func Call_Data_Functor_map__2061311692(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2061311692:
for {
if false { continue map__2061311692 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_AstTree_act())
}
}

func Call_Data_Functor_map__1572513044(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1572513044:
for {
if false { continue map__1572513044 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_ChurchFFICheatcode_act())
}
}

func Call_Data_Functor_map__4036895682(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__4036895682:
for {
if false { continue map__4036895682 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_ChurchFFI_act())
}
}

func Call_Data_Functor_map__683888683(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__683888683:
for {
if false { continue map__683888683 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_Church_act())
}
}

func Call_Data_Functor_map__1002236606(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1002236606:
for {
if false { continue map__1002236606 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_FibFFICheatcode_act())
}
}

func Call_Data_Functor_map__3097389096(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__3097389096:
for {
if false { continue map__3097389096 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_FibFFI_act())
}
}

func Call_Data_Functor_map__999202337(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__999202337:
for {
if false { continue map__999202337 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_Fib_act())
}
}

func Call_Data_Functor_map__395294283(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__395294283:
for {
if false { continue map__395294283 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_LazyEvaluationFFICheatcode_act())
}
}

func Call_Data_Functor_map__3468671773(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__3468671773:
for {
if false { continue map__3468671773 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_LazyEvaluationFFI_act())
}
}

func Call_Data_Functor_map__1648975668(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1648975668:
for {
if false { continue map__1648975668 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_LazyEvaluation_act())
}
}

func Call_Data_Functor_map__102065405(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__102065405:
for {
if false { continue map__102065405 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_ListOpsFFICheatcode_act())
}
}

func Call_Data_Functor_map__794090731(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__794090731:
for {
if false { continue map__794090731 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_ListOpsFFI_act())
}
}

func Call_Data_Functor_map__484586722(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__484586722:
for {
if false { continue map__484586722 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_ListOps_act())
}
}

func Call_Data_Functor_map__478722022(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__478722022:
for {
if false { continue map__478722022 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_PolymorphismFFICheatcode_act())
}
}

func Call_Data_Functor_map__105413808(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__105413808:
for {
if false { continue map__105413808 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_PolymorphismFFI_act())
}
}

func Call_Data_Functor_map__3791346649(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__3791346649:
for {
if false { continue map__3791346649 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_Polymorphism_act())
}
}

func Call_Data_Functor_map__4079021475(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__4079021475:
for {
if false { continue map__4079021475 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_PrimesFFICheatcode_act())
}
}

func Call_Data_Functor_map__1223398517(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1223398517:
for {
if false { continue map__1223398517 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_PrimesFFI_act())
}
}

func Call_Data_Functor_map__2625193884(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2625193884:
for {
if false { continue map__2625193884 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_Primes_act())
}
}

func Call_Data_Functor_map__2185106885(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2185106885:
for {
if false { continue map__2185106885 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_RBTreeFFICheatcode_act())
}
}

func Call_Data_Functor_map__293830419(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__293830419:
for {
if false { continue map__293830419 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_RBTreeFFI_act())
}
}

func Call_Data_Functor_map__982714682(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__982714682:
for {
if false { continue map__982714682 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_RBTree_act())
}
}

func Call_Data_Functor_map__4131672173(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__4131672173:
for {
if false { continue map__4131672173 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_RecordsFFICheatcode_act())
}
}

func Call_Data_Functor_map__4187694843(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__4187694843:
for {
if false { continue map__4187694843 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_RecordsFFI_act())
}
}

func Call_Data_Functor_map__4013754482(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__4013754482:
for {
if false { continue map__4013754482 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_Records_act())
}
}

func Call_Data_Functor_map__1723234304(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__1723234304:
for {
if false { continue map__1723234304 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_RowToListFFICheatcode_act())
}
}

func Call_Data_Functor_map__4039752726(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__4039752726:
for {
if false { continue map__4039752726 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_RowToListFFI_act())
}
}

func Call_Data_Functor_map__2414686431(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2414686431:
for {
if false { continue map__2414686431 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_RowToList_act())
}
}

func Call_Data_Functor_map__3441199853(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__3441199853:
for {
if false { continue map__3441199853 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_StateMonadFFICheatcode_act())
}
}

func Call_Data_Functor_map__2477564603(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2477564603:
for {
if false { continue map__2477564603 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_StateMonadFFI_act())
}
}

func Call_Data_Functor_map__196852562(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__196852562:
for {
if false { continue map__196852562 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_StateMonad_act())
}
}

func Call_Data_Functor_map__2501627275(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2501627275:
for {
if false { continue map__2501627275 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_TCOFFICheatcode_act())
}
}

func Call_Data_Functor_map__3939840413(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__3939840413:
for {
if false { continue map__3939840413 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_TCOFFI_act())
}
}

func Call_Data_Functor_map__2061195028(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
map__2061195028:
for {
if false { continue map__2061195028 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, __eta_norm_1_0, Get_Test_TCO_act())
}
}

func Call_Data_Functor_mapFlipped(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_Data_Functor_void(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(dictFunctor_0.V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
}

func Call_Data_Functor_void__3393626650(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__3393626650:
for {
if false { continue void__3393626650 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_Aff_functorAff()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), __eta_norm_0_0)
}
}

func Call_Data_Functor_void__990419905(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__990419905:
for {
if false { continue void__990419905 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_AckermannFFICheatcode_act())
}
}

func Call_Data_Functor_void__37931671(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__37931671:
for {
if false { continue void__37931671 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_AckermannFFI_act())
}
}

func Call_Data_Functor_void__2259642750(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__2259642750:
for {
if false { continue void__2259642750 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Ackermann_act())
}
}

func Call_Data_Functor_void__4174605798(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__4174605798:
for {
if false { continue void__4174605798 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ArrayOpsFFICheatcode_act())
}
}

func Call_Data_Functor_void__1062770544(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__1062770544:
for {
if false { continue void__1062770544 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ArrayOpsFFI_act())
}
}

func Call_Data_Functor_void__2350790585(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__2350790585:
for {
if false { continue void__2350790585 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ArrayOps_act())
}
}

func Call_Data_Functor_void__220841491(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__220841491:
for {
if false { continue void__220841491 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_AstTreeFFICheatcode_act())
}
}

func Call_Data_Functor_void__3758797381(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__3758797381:
for {
if false { continue void__3758797381 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_AstTreeFFI_act())
}
}

func Call_Data_Functor_void__3813769196(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__3813769196:
for {
if false { continue void__3813769196 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_AstTree_act())
}
}

func Call_Data_Functor_void__3599204660(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__3599204660:
for {
if false { continue void__3599204660 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ChurchFFICheatcode_act())
}
}

func Call_Data_Functor_void__1324701602(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__1324701602:
for {
if false { continue void__1324701602 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ChurchFFI_act())
}
}

func Call_Data_Functor_void__715030123(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__715030123:
for {
if false { continue void__715030123 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Church_act())
}
}

func Call_Data_Functor_void__380936478(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__380936478:
for {
if false { continue void__380936478 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_FibFFICheatcode_act())
}
}

func Call_Data_Functor_void__380628936(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__380628936:
for {
if false { continue void__380628936 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_FibFFI_act())
}
}

func Call_Data_Functor_void__2589128993(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__2589128993:
for {
if false { continue void__2589128993 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Fib_act())
}
}

func Call_Data_Functor_void__4146435083(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__4146435083:
for {
if false { continue void__4146435083 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_LazyEvaluationFFICheatcode_act())
}
}

func Call_Data_Functor_void__1659429917(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__1659429917:
for {
if false { continue void__1659429917 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_LazyEvaluationFFI_act())
}
}

func Call_Data_Functor_void__2038934356(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__2038934356:
for {
if false { continue void__2038934356 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_LazyEvaluation_act())
}
}

func Call_Data_Functor_void__234143613(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__234143613:
for {
if false { continue void__234143613 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ListOpsFFICheatcode_act())
}
}

func Call_Data_Functor_void__2094920235(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__2094920235:
for {
if false { continue void__2094920235 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ListOpsFFI_act())
}
}

func Call_Data_Functor_void__4269283266(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__4269283266:
for {
if false { continue void__4269283266 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ListOps_act())
}
}

func Call_Data_Functor_void__3317639750(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__3317639750:
for {
if false { continue void__3317639750 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_PolymorphismFFICheatcode_act())
}
}

func Call_Data_Functor_void__1880795472(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__1880795472:
for {
if false { continue void__1880795472 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_PolymorphismFFI_act())
}
}

func Call_Data_Functor_void__460922713(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__460922713:
for {
if false { continue void__460922713 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Polymorphism_act())
}
}

func Call_Data_Functor_void__3340832483(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__3340832483:
for {
if false { continue void__3340832483 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_PrimesFFICheatcode_act())
}
}

func Call_Data_Functor_void__982491637(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__982491637:
for {
if false { continue void__982491637 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_PrimesFFI_act())
}
}

func Call_Data_Functor_void__2681811132(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__2681811132:
for {
if false { continue void__2681811132 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Primes_act())
}
}

func Call_Data_Functor_void__2380011973(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__2380011973:
for {
if false { continue void__2380011973 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RBTreeFFICheatcode_act())
}
}

func Call_Data_Functor_void__2241417683(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__2241417683:
for {
if false { continue void__2241417683 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RBTreeFFI_act())
}
}

func Call_Data_Functor_void__3536814106(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__3536814106:
for {
if false { continue void__3536814106 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RBTree_act())
}
}

func Call_Data_Functor_void__1927686381(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__1927686381:
for {
if false { continue void__1927686381 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RecordsFFICheatcode_act())
}
}

func Call_Data_Functor_void__2080935483(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__2080935483:
for {
if false { continue void__2080935483 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RecordsFFI_act())
}
}

func Call_Data_Functor_void__2449968466(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__2449968466:
for {
if false { continue void__2449968466 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Records_act())
}
}

func Call_Data_Functor_void__143291040(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__143291040:
for {
if false { continue void__143291040 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RowToListFFICheatcode_act())
}
}

func Call_Data_Functor_void__672263030(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__672263030:
for {
if false { continue void__672263030 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RowToListFFI_act())
}
}

func Call_Data_Functor_void__355975199(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__355975199:
for {
if false { continue void__355975199 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RowToList_act())
}
}

func Call_Data_Functor_void__1926861677(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__1926861677:
for {
if false { continue void__1926861677 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_StateMonadFFICheatcode_act())
}
}

func Call_Data_Functor_void__4242248443(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__4242248443:
for {
if false { continue void__4242248443 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_StateMonadFFI_act())
}
}

func Call_Data_Functor_void__3690139954(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__3690139954:
for {
if false { continue void__3690139954 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_StateMonad_act())
}
}

func Call_Data_Functor_void__2838993739(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__2838993739:
for {
if false { continue void__2838993739 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_TCOFFICheatcode_act())
}
}

func Call_Data_Functor_void__929958557(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__929958557:
for {
if false { continue void__929958557 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_TCOFFI_act())
}
}

func Call_Data_Functor_void__1885598516(__eta_norm_0_unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__1885598516:
for {
if false { continue void__1885598516 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_TCO_act())
}
}

func Call_Data_Functor_void__565086979(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
void__565086979:
for {
if false { continue void__565086979 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), __eta_norm_0_0)
}
}

func Call_Data_Functor_voidLeft(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), f_1)
}

func Call_Data_Functor_voidLeft__1381072619(f_0_loop gopurs_runtime.Value, x_1_loop string) gopurs_runtime.Value {
voidLeft__1381072619:
for {
if false { continue voidLeft__1381072619 }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 string = x_1_loop
_ = x_1
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
arr_val_arrayMap0 := f_0
_ = arr_val_arrayMap0
arr_go_arrayMap0 := (*[]gopurs_runtime.Value)(arr_val_arrayMap0.UnsafePtr)
_ = arr_go_arrayMap0
res_go_arrayMap0 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap0))
_ = res_go_arrayMap0
for i_arrayMap0, v_arrayMap0 := range *arr_go_arrayMap0 {
res_go_arrayMap0[i_arrayMap0] = gopurs_runtime.Apply(gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(x_1)
}), v_arrayMap0)
}
return gopurs_runtime.Array(res_go_arrayMap0)
}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
}
}

func Call_Data_Functor_voidRight(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictFunctor_0.V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Functor_flap(dictFunctor_0_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], ff_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var ff_1 gopurs_runtime.Value = ff_1_loop
_ = ff_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, x_2)
}), ff_1)
}

func Rebox_Data_Functor_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Functor_1387998409_3094389156(in *Constructor_Data_Maybe_Just[struct{
	head int64
	tail string
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Str(orig.tail))
				}()
	return out
}

func Rebox_Data_Functor_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Functor_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Functor_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Functor_3094389156_1387998409(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
	head int64
	tail string
}] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[struct{
	head int64
	tail string
}]{}
		out.V0 = func() struct{
	head int64
	tail string
} {
					orig := in.V0
					_ = orig
					clone := struct{
	head int64
	tail string
}{}
					clone.head = gopurs_runtime.RecordGet(orig, "head").IntVal
					clone.tail = gopurs_runtime.RecordGet(orig, "tail").StrVal()
					return clone
				}()
	return out
}

func Rebox_Data_Functor_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}

func Rebox_Data_Functor_562069347_2812149806(in *Constructor_Data_Functor_Functor[uint32]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Get_Data_Functor_arrayMap() gopurs_runtime.Value {
	return _Gopurs_Data_Functor_ArrayMap
}
