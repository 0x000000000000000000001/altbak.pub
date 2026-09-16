package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Maybe_identity gopurs_runtime.Value
var once_Data_Maybe_identity sync.Once
func Get_Data_Maybe_identity() gopurs_runtime.Value {
	once_Data_Maybe_identity.Do(func() {
		cache_Data_Maybe_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Data_Maybe_identity
}

var cache_Data_Maybe_identity1 gopurs_runtime.Value
var once_Data_Maybe_identity1 sync.Once
func Get_Data_Maybe_identity1() gopurs_runtime.Value {
	once_Data_Maybe_identity1.Do(func() {
		cache_Data_Maybe_identity1 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Data_Maybe_identity1
}

var cache_Data_Maybe_Nothing gopurs_runtime.Value
var once_Data_Maybe_Nothing sync.Once
func Get_Data_Maybe_Nothing() gopurs_runtime.Value {
	once_Data_Maybe_Nothing.Do(func() {
		cache_Data_Maybe_Nothing = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
	})
	return cache_Data_Maybe_Nothing
}

var cache_Data_Maybe_Just gopurs_runtime.Value
var once_Data_Maybe_Just sync.Once
func Get_Data_Maybe_Just() gopurs_runtime.Value {
	once_Data_Maybe_Just.Do(func() {
		cache_Data_Maybe_Just = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, value0}))}
})
	})
	return cache_Data_Maybe_Just
}

var cache_Data_Maybe_Just__237222391 gopurs_runtime.Value
var once_Data_Maybe_Just__237222391 sync.Once
func Get_Data_Maybe_Just__237222391() gopurs_runtime.Value {
	once_Data_Maybe_Just__237222391.Do(func() {
		cache_Data_Maybe_Just__237222391 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__237222391(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__237222391
}

var cache_Data_Maybe_Just__3686219312 gopurs_runtime.Value
var once_Data_Maybe_Just__3686219312 sync.Once
func Get_Data_Maybe_Just__3686219312() gopurs_runtime.Value {
	once_Data_Maybe_Just__3686219312.Do(func() {
		cache_Data_Maybe_Just__3686219312 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3686219312(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__3686219312
}

var cache_Data_Maybe_Just__2614824046 gopurs_runtime.Value
var once_Data_Maybe_Just__2614824046 sync.Once
func Get_Data_Maybe_Just__2614824046() gopurs_runtime.Value {
	once_Data_Maybe_Just__2614824046.Do(func() {
		cache_Data_Maybe_Just__2614824046 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2614824046(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__2614824046
}

var cache_Data_Maybe_Just__3785071805 gopurs_runtime.Value
var once_Data_Maybe_Just__3785071805 sync.Once
func Get_Data_Maybe_Just__3785071805() gopurs_runtime.Value {
	once_Data_Maybe_Just__3785071805.Do(func() {
		cache_Data_Maybe_Just__3785071805 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3785071805(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__3785071805
}

var cache_Data_Maybe_Just__2633198667 gopurs_runtime.Value
var once_Data_Maybe_Just__2633198667 sync.Once
func Get_Data_Maybe_Just__2633198667() gopurs_runtime.Value {
	once_Data_Maybe_Just__2633198667.Do(func() {
		cache_Data_Maybe_Just__2633198667 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2633198667(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__2633198667
}

var cache_Data_Maybe_Just__105528475 gopurs_runtime.Value
var once_Data_Maybe_Just__105528475 sync.Once
func Get_Data_Maybe_Just__105528475() gopurs_runtime.Value {
	once_Data_Maybe_Just__105528475.Do(func() {
		cache_Data_Maybe_Just__105528475 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__105528475(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__105528475
}

var cache_Data_Maybe_Just__4249312965 gopurs_runtime.Value
var once_Data_Maybe_Just__4249312965 sync.Once
func Get_Data_Maybe_Just__4249312965() gopurs_runtime.Value {
	once_Data_Maybe_Just__4249312965.Do(func() {
		cache_Data_Maybe_Just__4249312965 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__4249312965(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__4249312965
}

var cache_Data_Maybe_Just__4573476 gopurs_runtime.Value
var once_Data_Maybe_Just__4573476 sync.Once
func Get_Data_Maybe_Just__4573476() gopurs_runtime.Value {
	once_Data_Maybe_Just__4573476.Do(func() {
		cache_Data_Maybe_Just__4573476 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__4573476(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__4573476
}

var cache_Data_Maybe_Just__496498436 gopurs_runtime.Value
var once_Data_Maybe_Just__496498436 sync.Once
func Get_Data_Maybe_Just__496498436() gopurs_runtime.Value {
	once_Data_Maybe_Just__496498436.Do(func() {
		cache_Data_Maybe_Just__496498436 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__496498436(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__496498436
}

var cache_Data_Maybe_Just__3591011291 gopurs_runtime.Value
var once_Data_Maybe_Just__3591011291 sync.Once
func Get_Data_Maybe_Just__3591011291() gopurs_runtime.Value {
	once_Data_Maybe_Just__3591011291.Do(func() {
		cache_Data_Maybe_Just__3591011291 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3591011291(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__3591011291
}

var cache_Data_Maybe_Just__139360115 gopurs_runtime.Value
var once_Data_Maybe_Just__139360115 sync.Once
func Get_Data_Maybe_Just__139360115() gopurs_runtime.Value {
	once_Data_Maybe_Just__139360115.Do(func() {
		cache_Data_Maybe_Just__139360115 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__139360115(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__139360115
}

var cache_Data_Maybe_Just__1605594846 gopurs_runtime.Value
var once_Data_Maybe_Just__1605594846 sync.Once
func Get_Data_Maybe_Just__1605594846() gopurs_runtime.Value {
	once_Data_Maybe_Just__1605594846.Do(func() {
		cache_Data_Maybe_Just__1605594846 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__1605594846(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__1605594846
}

var cache_Data_Maybe_Just__2013662992 gopurs_runtime.Value
var once_Data_Maybe_Just__2013662992 sync.Once
func Get_Data_Maybe_Just__2013662992() gopurs_runtime.Value {
	once_Data_Maybe_Just__2013662992.Do(func() {
		cache_Data_Maybe_Just__2013662992 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2013662992(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__2013662992
}

var cache_Data_Maybe_Just__764130881 gopurs_runtime.Value
var once_Data_Maybe_Just__764130881 sync.Once
func Get_Data_Maybe_Just__764130881() gopurs_runtime.Value {
	once_Data_Maybe_Just__764130881.Do(func() {
		cache_Data_Maybe_Just__764130881 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__764130881(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__764130881
}

var cache_Data_Maybe_Just__2900182732 gopurs_runtime.Value
var once_Data_Maybe_Just__2900182732 sync.Once
func Get_Data_Maybe_Just__2900182732() gopurs_runtime.Value {
	once_Data_Maybe_Just__2900182732.Do(func() {
		cache_Data_Maybe_Just__2900182732 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2900182732(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__2900182732
}

var cache_Data_Maybe_Just__1960558085 gopurs_runtime.Value
var once_Data_Maybe_Just__1960558085 sync.Once
func Get_Data_Maybe_Just__1960558085() gopurs_runtime.Value {
	once_Data_Maybe_Just__1960558085.Do(func() {
		cache_Data_Maybe_Just__1960558085 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__1960558085(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__1960558085
}

var cache_Data_Maybe_Just__555433957 gopurs_runtime.Value
var once_Data_Maybe_Just__555433957 sync.Once
func Get_Data_Maybe_Just__555433957() gopurs_runtime.Value {
	once_Data_Maybe_Just__555433957.Do(func() {
		cache_Data_Maybe_Just__555433957 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__555433957(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__555433957
}

var cache_Data_Maybe_Just__98062938 gopurs_runtime.Value
var once_Data_Maybe_Just__98062938 sync.Once
func Get_Data_Maybe_Just__98062938() gopurs_runtime.Value {
	once_Data_Maybe_Just__98062938.Do(func() {
		cache_Data_Maybe_Just__98062938 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__98062938(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__98062938
}

var cache_Data_Maybe_Just__3422863715 gopurs_runtime.Value
var once_Data_Maybe_Just__3422863715 sync.Once
func Get_Data_Maybe_Just__3422863715() gopurs_runtime.Value {
	once_Data_Maybe_Just__3422863715.Do(func() {
		cache_Data_Maybe_Just__3422863715 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3422863715(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__3422863715
}

var cache_Data_Maybe_Just__165672699 gopurs_runtime.Value
var once_Data_Maybe_Just__165672699 sync.Once
func Get_Data_Maybe_Just__165672699() gopurs_runtime.Value {
	once_Data_Maybe_Just__165672699.Do(func() {
		cache_Data_Maybe_Just__165672699 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__165672699(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](__eta_norm_0_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__165672699
}

var cache_Data_Maybe_Just__1852526845 gopurs_runtime.Value
var once_Data_Maybe_Just__1852526845 sync.Once
func Get_Data_Maybe_Just__1852526845() gopurs_runtime.Value {
	once_Data_Maybe_Just__1852526845.Do(func() {
		cache_Data_Maybe_Just__1852526845 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__1852526845(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__1852526845
}

var cache_Data_Maybe_Just__819016571 gopurs_runtime.Value
var once_Data_Maybe_Just__819016571 sync.Once
func Get_Data_Maybe_Just__819016571() gopurs_runtime.Value {
	once_Data_Maybe_Just__819016571.Do(func() {
		cache_Data_Maybe_Just__819016571 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__819016571(Rebox_Data_Maybe_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_0_0_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__819016571
}

var cache_Data_Maybe_Just__3139849822 gopurs_runtime.Value
var once_Data_Maybe_Just__3139849822 sync.Once
func Get_Data_Maybe_Just__3139849822() gopurs_runtime.Value {
	once_Data_Maybe_Just__3139849822.Do(func() {
		cache_Data_Maybe_Just__3139849822 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3139849822(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__3139849822
}

var cache_Data_Maybe_Just__1323530841 gopurs_runtime.Value
var once_Data_Maybe_Just__1323530841 sync.Once
func Get_Data_Maybe_Just__1323530841() gopurs_runtime.Value {
	once_Data_Maybe_Just__1323530841.Do(func() {
		cache_Data_Maybe_Just__1323530841 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__1323530841(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__1323530841
}

var cache_Data_Maybe_Just__821644466 gopurs_runtime.Value
var once_Data_Maybe_Just__821644466 sync.Once
func Get_Data_Maybe_Just__821644466() gopurs_runtime.Value {
	once_Data_Maybe_Just__821644466.Do(func() {
		cache_Data_Maybe_Just__821644466 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__821644466(uint32(__eta_norm_0_unused_0_box.IntVal))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__821644466
}

var cache_Data_Maybe_Just__1425602363 gopurs_runtime.Value
var once_Data_Maybe_Just__1425602363 sync.Once
func Get_Data_Maybe_Just__1425602363() gopurs_runtime.Value {
	once_Data_Maybe_Just__1425602363.Do(func() {
		cache_Data_Maybe_Just__1425602363 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__1425602363(Rebox_Data_Maybe_138441832_3363075976(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_0_0_box)))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__1425602363
}

var cache_Data_Maybe_Just__635767739 gopurs_runtime.Value
var once_Data_Maybe_Just__635767739 sync.Once
func Get_Data_Maybe_Just__635767739() gopurs_runtime.Value {
	once_Data_Maybe_Just__635767739.Do(func() {
		cache_Data_Maybe_Just__635767739 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__635767739((__eta_norm_0_0_box.IntVal) != (0))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__635767739
}

var cache_Data_Maybe_Just__2688222971 gopurs_runtime.Value
var once_Data_Maybe_Just__2688222971 sync.Once
func Get_Data_Maybe_Just__2688222971() gopurs_runtime.Value {
	once_Data_Maybe_Just__2688222971.Do(func() {
		cache_Data_Maybe_Just__2688222971 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2688222971(__eta_norm_0_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__2688222971
}

var cache_Data_Maybe_Just__2649619419 gopurs_runtime.Value
var once_Data_Maybe_Just__2649619419 sync.Once
func Get_Data_Maybe_Just__2649619419() gopurs_runtime.Value {
	once_Data_Maybe_Just__2649619419.Do(func() {
		cache_Data_Maybe_Just__2649619419 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2649619419(__eta_norm_0_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__2649619419
}

var cache_Data_Maybe_Just__2582300155 gopurs_runtime.Value
var once_Data_Maybe_Just__2582300155 sync.Once
func Get_Data_Maybe_Just__2582300155() gopurs_runtime.Value {
	once_Data_Maybe_Just__2582300155.Do(func() {
		cache_Data_Maybe_Just__2582300155 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__2582300155(__eta_norm_0_0_box.FloatVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__2582300155
}

var cache_Data_Maybe_Just__651901051 gopurs_runtime.Value
var once_Data_Maybe_Just__651901051 sync.Once
func Get_Data_Maybe_Just__651901051() gopurs_runtime.Value {
	once_Data_Maybe_Just__651901051.Do(func() {
		cache_Data_Maybe_Just__651901051 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__651901051(func() struct{
	head string
	tail string
} {
					orig := __eta_norm_0_0_box
					_ = orig
					clone := struct{
	head string
	tail string
}{}
					clone.head = gopurs_runtime.RecordGet(orig, "head").StrVal()
					clone.tail = gopurs_runtime.RecordGet(orig, "tail").StrVal()
					return clone
				}())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__651901051
}

var cache_Data_Maybe_Just__1601057723 gopurs_runtime.Value
var once_Data_Maybe_Just__1601057723 sync.Once
func Get_Data_Maybe_Just__1601057723() gopurs_runtime.Value {
	once_Data_Maybe_Just__1601057723.Do(func() {
		cache_Data_Maybe_Just__1601057723 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__1601057723(func() struct{
	head int64
	tail string
} {
					orig := __eta_norm_0_0_box
					_ = orig
					clone := struct{
	head int64
	tail string
}{}
					clone.head = gopurs_runtime.RecordGet(orig, "head").IntVal
					clone.tail = gopurs_runtime.RecordGet(orig, "tail").StrVal()
					return clone
				}())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__1601057723
}

var cache_Data_Maybe_Just__1492091643 gopurs_runtime.Value
var once_Data_Maybe_Just__1492091643 sync.Once
func Get_Data_Maybe_Just__1492091643() gopurs_runtime.Value {
	once_Data_Maybe_Just__1492091643.Do(func() {
		cache_Data_Maybe_Just__1492091643 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__1492091643(__eta_norm_0_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__1492091643
}

var cache_Data_Maybe_Just__3221933828 gopurs_runtime.Value
var once_Data_Maybe_Just__3221933828 sync.Once
func Get_Data_Maybe_Just__3221933828() gopurs_runtime.Value {
	once_Data_Maybe_Just__3221933828.Do(func() {
		cache_Data_Maybe_Just__3221933828 = gopurs_runtime.Func(func(__eta_norm_0_unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_Just__3221933828(__eta_norm_0_unused_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_Just__3221933828
}

var cache_Data_Maybe_showMaybe gopurs_runtime.Value
var once_Data_Maybe_showMaybe sync.Once
func Get_Data_Maybe_showMaybe() gopurs_runtime.Value {
	once_Data_Maybe_showMaybe.Do(func() {
		cache_Data_Maybe_showMaybe = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_showMaybe(dictShow_0_box)
})
	})
	return cache_Data_Maybe_showMaybe
}

var cache_Data_Maybe_semigroupMaybe gopurs_runtime.Value
var once_Data_Maybe_semigroupMaybe sync.Once
func Get_Data_Maybe_semigroupMaybe() gopurs_runtime.Value {
	once_Data_Maybe_semigroupMaybe.Do(func() {
		cache_Data_Maybe_semigroupMaybe = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_semigroupMaybe(dictSemigroup_0_box)
})
	})
	return cache_Data_Maybe_semigroupMaybe
}

var cache_Data_Maybe_optional gopurs_runtime.Value
var once_Data_Maybe_optional sync.Once
func Get_Data_Maybe_optional() gopurs_runtime.Value {
	once_Data_Maybe_optional.Do(func() {
		cache_Data_Maybe_optional = gopurs_runtime.Func(func(dictAlt_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_optional(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](dictAlt_0_box))
})
	})
	return cache_Data_Maybe_optional
}

var cache_Data_Maybe_monoidMaybe gopurs_runtime.Value
var once_Data_Maybe_monoidMaybe sync.Once
func Get_Data_Maybe_monoidMaybe() gopurs_runtime.Value {
	once_Data_Maybe_monoidMaybe.Do(func() {
		cache_Data_Maybe_monoidMaybe = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_monoidMaybe(dictSemigroup_0_box)
})
	})
	return cache_Data_Maybe_monoidMaybe
}

var cache_Data_Maybe_maybe_prime_ gopurs_runtime.Value
var once_Data_Maybe_maybe_prime_ sync.Once
func Get_Data_Maybe_maybe_prime_() gopurs_runtime.Value {
	once_Data_Maybe_maybe_prime_.Do(func() {
		cache_Data_Maybe_maybe_prime_ = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe_prime_(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe_prime_
}

var cache_Data_Maybe_maybe gopurs_runtime.Value
var once_Data_Maybe_maybe sync.Once
func Get_Data_Maybe_maybe() gopurs_runtime.Value {
	once_Data_Maybe_maybe.Do(func() {
		cache_Data_Maybe_maybe = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_maybe(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_Data_Maybe_maybe
}

var cache_Data_Maybe_maybe__4265365554 gopurs_runtime.Value
var once_Data_Maybe_maybe__4265365554 sync.Once
func Get_Data_Maybe_maybe__4265365554() gopurs_runtime.Value {
	once_Data_Maybe_maybe__4265365554.Do(func() {
		cache_Data_Maybe_maybe__4265365554 = gopurs_runtime.Func3(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_maybe__4265365554(uint32(v_unused_0_box.IntVal), v1_unused_1_box, Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box)))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_maybe__4265365554
}

var cache_Data_Maybe_maybe__1711083959 gopurs_runtime.Value
var once_Data_Maybe_maybe__1711083959 sync.Once
func Get_Data_Maybe_maybe__1711083959() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1711083959.Do(func() {
		cache_Data_Maybe_maybe__1711083959 = gopurs_runtime.Func3(func(v_unused_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_maybe__1711083959(uint32(v_unused_0_box.IntVal), v1_unused_1_box, Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box)))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_maybe__1711083959
}

var cache_Data_Maybe_maybe__2146642512 gopurs_runtime.Value
var once_Data_Maybe_maybe__2146642512 sync.Once
func Get_Data_Maybe_maybe__2146642512() gopurs_runtime.Value {
	once_Data_Maybe_maybe__2146642512.Do(func() {
		cache_Data_Maybe_maybe__2146642512 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Maybe_maybe__2146642512(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box), v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Maybe_maybe__2146642512
}

var cache_Data_Maybe_maybe__2726997443 gopurs_runtime.Value
var once_Data_Maybe_maybe__2726997443 sync.Once
func Get_Data_Maybe_maybe__2726997443() gopurs_runtime.Value {
	once_Data_Maybe_maybe__2726997443.Do(func() {
		cache_Data_Maybe_maybe__2726997443 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__2726997443((v_0_box.IntVal) != (0), v1_1_box, Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box))))
})
	})
	return cache_Data_Maybe_maybe__2726997443
}

var cache_Data_Maybe_maybe__4252741795 gopurs_runtime.Value
var once_Data_Maybe_maybe__4252741795 sync.Once
func Get_Data_Maybe_maybe__4252741795() gopurs_runtime.Value {
	once_Data_Maybe_maybe__4252741795.Do(func() {
		cache_Data_Maybe_maybe__4252741795 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__4252741795((v_0_box.IntVal) != (0), v1_1_box, Rebox_Data_Maybe_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box))))
})
	})
	return cache_Data_Maybe_maybe__4252741795
}

var cache_Data_Maybe_maybe__3642269347 gopurs_runtime.Value
var once_Data_Maybe_maybe__3642269347 sync.Once
func Get_Data_Maybe_maybe__3642269347() gopurs_runtime.Value {
	once_Data_Maybe_maybe__3642269347.Do(func() {
		cache_Data_Maybe_maybe__3642269347 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_maybe__3642269347((v_0_box.IntVal) != (0), v1_1_box, Rebox_Data_Maybe_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box))))
})
	})
	return cache_Data_Maybe_maybe__3642269347
}

var cache_Data_Maybe_maybe__2295219908 gopurs_runtime.Value
var once_Data_Maybe_maybe__2295219908 sync.Once
func Get_Data_Maybe_maybe__2295219908() gopurs_runtime.Value {
	once_Data_Maybe_maybe__2295219908.Do(func() {
		cache_Data_Maybe_maybe__2295219908 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_maybe__2295219908(v_0_box.IntVal, v1_1_box, Rebox_Data_Maybe_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box))))
})
	})
	return cache_Data_Maybe_maybe__2295219908
}

var cache_Data_Maybe_maybe__3140274526 gopurs_runtime.Value
var once_Data_Maybe_maybe__3140274526 sync.Once
func Get_Data_Maybe_maybe__3140274526() gopurs_runtime.Value {
	once_Data_Maybe_maybe__3140274526.Do(func() {
		cache_Data_Maybe_maybe__3140274526 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_unused_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_maybe__3140274526(v_0_box.IntVal, v1_unused_1_box, Rebox_Data_Maybe_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box))))
})
	})
	return cache_Data_Maybe_maybe__3140274526
}

var cache_Data_Maybe_maybe__1210512692 gopurs_runtime.Value
var once_Data_Maybe_maybe__1210512692 sync.Once
func Get_Data_Maybe_maybe__1210512692() gopurs_runtime.Value {
	once_Data_Maybe_maybe__1210512692.Do(func() {
		cache_Data_Maybe_maybe__1210512692 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_maybe__1210512692(v_0_box.IntVal, v1_1_box, Rebox_Data_Maybe_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2_box))))
})
	})
	return cache_Data_Maybe_maybe__1210512692
}

var cache_Data_Maybe_isNothing gopurs_runtime.Value
var once_Data_Maybe_isNothing sync.Once
func Get_Data_Maybe_isNothing() gopurs_runtime.Value {
	once_Data_Maybe_isNothing.Do(func() {
		cache_Data_Maybe_isNothing = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_Data_Maybe_isNothing
}

var cache_Data_Maybe_isNothing__2716061749 gopurs_runtime.Value
var once_Data_Maybe_isNothing__2716061749 sync.Once
func Get_Data_Maybe_isNothing__2716061749() gopurs_runtime.Value {
	once_Data_Maybe_isNothing__2716061749.Do(func() {
		cache_Data_Maybe_isNothing__2716061749 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing__2716061749(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_0_box))))
})
	})
	return cache_Data_Maybe_isNothing__2716061749
}

var cache_Data_Maybe_isNothing__1937207052 gopurs_runtime.Value
var once_Data_Maybe_isNothing__1937207052 sync.Once
func Get_Data_Maybe_isNothing__1937207052() gopurs_runtime.Value {
	once_Data_Maybe_isNothing__1937207052.Do(func() {
		cache_Data_Maybe_isNothing__1937207052 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isNothing__1937207052(Rebox_Data_Maybe_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_0_box))))
})
	})
	return cache_Data_Maybe_isNothing__1937207052
}

var cache_Data_Maybe_isJust gopurs_runtime.Value
var once_Data_Maybe_isJust sync.Once
func Get_Data_Maybe_isJust() gopurs_runtime.Value {
	once_Data_Maybe_isJust.Do(func() {
		cache_Data_Maybe_isJust = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_Data_Maybe_isJust
}

var cache_Data_Maybe_isJust__1937207052 gopurs_runtime.Value
var once_Data_Maybe_isJust__1937207052 sync.Once
func Get_Data_Maybe_isJust__1937207052() gopurs_runtime.Value {
	once_Data_Maybe_isJust__1937207052.Do(func() {
		cache_Data_Maybe_isJust__1937207052 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust__1937207052(Rebox_Data_Maybe_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_0_box))))
})
	})
	return cache_Data_Maybe_isJust__1937207052
}

var cache_Data_Maybe_isJust__3089511036 gopurs_runtime.Value
var once_Data_Maybe_isJust__3089511036 sync.Once
func Get_Data_Maybe_isJust__3089511036() gopurs_runtime.Value {
	once_Data_Maybe_isJust__3089511036.Do(func() {
		cache_Data_Maybe_isJust__3089511036 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Maybe_isJust__3089511036(Rebox_Data_Maybe_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_0_box))))
})
	})
	return cache_Data_Maybe_isJust__3089511036
}

var cache_Data_Maybe_genericMaybe gopurs_runtime.Value
var once_Data_Maybe_genericMaybe sync.Once
func Get_Data_Maybe_genericMaybe() gopurs_runtime.Value {
	once_Data_Maybe_genericMaybe.Do(func() {
		cache_Data_Maybe_genericMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1921946594, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_732223057_2818661616((&Constructor_Data_Generic_Rep_Generic[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_497217223_1323331594((&Constructor_Data_Generic_Rep_Inl[uint32, gopurs_runtime.Value]{1, 1454898258})))}
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_0)
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_660678937_2687169876((&Constructor_Data_Generic_Rep_Inr[uint32, gopurs_runtime.Value]{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_0.UnsafePtr).V0})))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (x_0.Type == 9 && x_0.IntVal == 3478632216) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 492034566) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(*Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(x_0.UnsafePtr).V0, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})})))}
	})
	return cache_Data_Maybe_genericMaybe
}

var cache_Data_Maybe_functorMaybe gopurs_runtime.Value
var once_Data_Maybe_functorMaybe sync.Once
func Get_Data_Maybe_functorMaybe() gopurs_runtime.Value {
	once_Data_Maybe_functorMaybe.Do(func() {
		cache_Data_Maybe_functorMaybe = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3689823567_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_1)
_ = __t_tag_0
if (__t_tag_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(v_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})})))}
	})
	return cache_Data_Maybe_functorMaybe
}

var cache_Data_Maybe_invariantMaybe gopurs_runtime.Value
var once_Data_Maybe_invariantMaybe sync.Once
func Get_Data_Maybe_invariantMaybe() gopurs_runtime.Value {
	once_Data_Maybe_invariantMaybe.Do(func() {
		cache_Data_Maybe_invariantMaybe = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3818752775_2241343270((&Constructor_Data_Functor_Invariant_Invariant[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Get_Data_Functor_Invariant_imapF(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3689823567_2812149806(Rebox_Data_Maybe_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe()))))})})))}
	})
	return cache_Data_Maybe_invariantMaybe
}

var cache_Data_Maybe_fromMaybe_prime_ gopurs_runtime.Value
var once_Data_Maybe_fromMaybe_prime_ sync.Once
func Get_Data_Maybe_fromMaybe_prime_() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe_prime_.Do(func() {
		cache_Data_Maybe_fromMaybe_prime_ = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_fromMaybe_prime_(a_0_box)
})
	})
	return cache_Data_Maybe_fromMaybe_prime_
}

var cache_Data_Maybe_fromMaybe gopurs_runtime.Value
var once_Data_Maybe_fromMaybe sync.Once
func Get_Data_Maybe_fromMaybe() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe.Do(func() {
		cache_Data_Maybe_fromMaybe = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_fromMaybe(a_0_box)
})
	})
	return cache_Data_Maybe_fromMaybe
}

var cache_Data_Maybe_fromMaybe__3286440201 gopurs_runtime.Value
var once_Data_Maybe_fromMaybe__3286440201 sync.Once
func Get_Data_Maybe_fromMaybe__3286440201() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe__3286440201.Do(func() {
		cache_Data_Maybe_fromMaybe__3286440201 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_fromMaybe__3286440201(uint32(a_unused_0_box.IntVal), Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_fromMaybe__3286440201
}

var cache_Data_Maybe_fromMaybe__516259212 gopurs_runtime.Value
var once_Data_Maybe_fromMaybe__516259212 sync.Once
func Get_Data_Maybe_fromMaybe__516259212() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe__516259212.Do(func() {
		cache_Data_Maybe_fromMaybe__516259212 = gopurs_runtime.Func2(func(a_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_fromMaybe__516259212(uint32(a_unused_0_box.IntVal), Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box)))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_fromMaybe__516259212
}

var cache_Data_Maybe_fromMaybe__816552133 gopurs_runtime.Value
var once_Data_Maybe_fromMaybe__816552133 sync.Once
func Get_Data_Maybe_fromMaybe__816552133() gopurs_runtime.Value {
	once_Data_Maybe_fromMaybe__816552133.Do(func() {
		cache_Data_Maybe_fromMaybe__816552133 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_fromMaybe__816552133(a_0_box.IntVal, Rebox_Data_Maybe_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1_box))))
})
	})
	return cache_Data_Maybe_fromMaybe__816552133
}

var cache_Data_Maybe_fromJust gopurs_runtime.Value
var once_Data_Maybe_fromJust sync.Once
func Get_Data_Maybe_fromJust() gopurs_runtime.Value {
	once_Data_Maybe_fromJust.Do(func() {
		cache_Data_Maybe_fromJust = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_fromJust(_dollar___unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_Data_Maybe_fromJust
}

var cache_Data_Maybe_fromJust__2539482 gopurs_runtime.Value
var once_Data_Maybe_fromJust__2539482 sync.Once
func Get_Data_Maybe_fromJust__2539482() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__2539482.Do(func() {
		cache_Data_Maybe_fromJust__2539482 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_fromJust__2539482(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box)))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_fromJust__2539482
}

var cache_Data_Maybe_fromJust__1001929434 gopurs_runtime.Value
var once_Data_Maybe_fromJust__1001929434 sync.Once
func Get_Data_Maybe_fromJust__1001929434() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__1001929434.Do(func() {
		cache_Data_Maybe_fromJust__1001929434 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Maybe_fromJust__1001929434(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box)))), UnsafePtr: nil}
})
	})
	return cache_Data_Maybe_fromJust__1001929434
}

var cache_Data_Maybe_fromJust__171824954 gopurs_runtime.Value
var once_Data_Maybe_fromJust__171824954 sync.Once
func Get_Data_Maybe_fromJust__171824954() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__171824954.Do(func() {
		cache_Data_Maybe_fromJust__171824954 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_fromJust__171824954(Rebox_Data_Maybe_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))))}
})
	})
	return cache_Data_Maybe_fromJust__171824954
}

var cache_Data_Maybe_fromJust__3704324326 gopurs_runtime.Value
var once_Data_Maybe_fromJust__3704324326 sync.Once
func Get_Data_Maybe_fromJust__3704324326() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__3704324326.Do(func() {
		cache_Data_Maybe_fromJust__3704324326 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer(Call_Data_Maybe_fromJust__3704324326(Rebox_Data_Maybe_3094389156_3839235747(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))))}
})
	})
	return cache_Data_Maybe_fromJust__3704324326
}

var cache_Data_Maybe_fromJust__55301350 gopurs_runtime.Value
var once_Data_Maybe_fromJust__55301350 sync.Once
func Get_Data_Maybe_fromJust__55301350() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__55301350.Do(func() {
		cache_Data_Maybe_fromJust__55301350 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Maybe_fromJust__55301350(Rebox_Data_Maybe_3094389156_1731162461(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box)))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Maybe_fromJust__55301350
}

var cache_Data_Maybe_fromJust__1673906054 gopurs_runtime.Value
var once_Data_Maybe_fromJust__1673906054 sync.Once
func Get_Data_Maybe_fromJust__1673906054() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__1673906054.Do(func() {
		cache_Data_Maybe_fromJust__1673906054 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Maybe_fromJust__1673906054(Rebox_Data_Maybe_3094389156_1495236409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box)))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Maybe_fromJust__1673906054
}

var cache_Data_Maybe_fromJust__218925574 gopurs_runtime.Value
var once_Data_Maybe_fromJust__218925574 sync.Once
func Get_Data_Maybe_fromJust__218925574() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__218925574.Do(func() {
		cache_Data_Maybe_fromJust__218925574 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Maybe_fromJust__218925574(Rebox_Data_Maybe_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))))
})
	})
	return cache_Data_Maybe_fromJust__218925574
}

var cache_Data_Maybe_fromJust__2463559942 gopurs_runtime.Value
var once_Data_Maybe_fromJust__2463559942 sync.Once
func Get_Data_Maybe_fromJust__2463559942() gopurs_runtime.Value {
	once_Data_Maybe_fromJust__2463559942.Do(func() {
		cache_Data_Maybe_fromJust__2463559942 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Maybe_fromJust__2463559942(Rebox_Data_Maybe_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))))
})
	})
	return cache_Data_Maybe_fromJust__2463559942
}

var cache_Data_Maybe_extendMaybe gopurs_runtime.Value
var once_Data_Maybe_extendMaybe sync.Once
func Get_Data_Maybe_extendMaybe() gopurs_runtime.Value {
	once_Data_Maybe_extendMaybe.Do(func() {
		cache_Data_Maybe_extendMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_570732504_3290176857((&Constructor_Control_Extend_Extend[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3689823567_2812149806(Rebox_Data_Maybe_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_1)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_1))}), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})})))}
	})
	return cache_Data_Maybe_extendMaybe
}

var cache_Data_Maybe_eqMaybe gopurs_runtime.Value
var once_Data_Maybe_eqMaybe sync.Once
func Get_Data_Maybe_eqMaybe() gopurs_runtime.Value {
	once_Data_Maybe_eqMaybe.Do(func() {
		cache_Data_Maybe_eqMaybe = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_eqMaybe(dictEq_0_box)
})
	})
	return cache_Data_Maybe_eqMaybe
}

var cache_Data_Maybe_ordMaybe gopurs_runtime.Value
var once_Data_Maybe_ordMaybe sync.Once
func Get_Data_Maybe_ordMaybe() gopurs_runtime.Value {
	once_Data_Maybe_ordMaybe.Do(func() {
		cache_Data_Maybe_ordMaybe = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_ordMaybe(dictOrd_0_box)
})
	})
	return cache_Data_Maybe_ordMaybe
}

var cache_Data_Maybe_eq1Maybe gopurs_runtime.Value
var once_Data_Maybe_eq1Maybe sync.Once
func Get_Data_Maybe_eq1Maybe() gopurs_runtime.Value {
	once_Data_Maybe_eq1Maybe.Do(func() {
		cache_Data_Maybe_eq1Maybe = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1662389854_1766074591((&Constructor_Data_Eq_Eq1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Maybe_eqMaybe(dictEq_0)))
})})))}
	})
	return cache_Data_Maybe_eq1Maybe
}

var cache_Data_Maybe_ord1Maybe gopurs_runtime.Value
var once_Data_Maybe_ord1Maybe sync.Once
func Get_Data_Maybe_ord1Maybe() gopurs_runtime.Value {
	once_Data_Maybe_ord1Maybe.Do(func() {
		cache_Data_Maybe_ord1Maybe = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_170825214_3985601471((&Constructor_Data_Ord_Ord1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1662389854_1766074591(Rebox_Data_Maybe_1766074591_1662389854(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Maybe_eq1Maybe()))))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Maybe_ordMaybe(dictOrd_0)))
})})))}
	})
	return cache_Data_Maybe_ord1Maybe
}

var cache_Data_Maybe_boundedMaybe gopurs_runtime.Value
var once_Data_Maybe_boundedMaybe sync.Once
func Get_Data_Maybe_boundedMaybe() gopurs_runtime.Value {
	once_Data_Maybe_boundedMaybe.Do(func() {
		cache_Data_Maybe_boundedMaybe = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_boundedMaybe(dictBounded_0_box)
})
	})
	return cache_Data_Maybe_boundedMaybe
}

var cache_Data_Maybe_applyMaybe gopurs_runtime.Value
var once_Data_Maybe_applyMaybe sync.Once
func Get_Data_Maybe_applyMaybe() gopurs_runtime.Value {
	once_Data_Maybe_applyMaybe.Do(func() {
		cache_Data_Maybe_applyMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3552963512_3741347833((&Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3689823567_2812149806(Rebox_Data_Maybe_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
_ = __t_tag_0
if (__t_tag_0 != nil) {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_1)
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_2:
__t4 = __t2
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
_ = __t_tag_3
if (__t_tag_3 == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
})})))}
	})
	return cache_Data_Maybe_applyMaybe
}

var cache_Data_Maybe_bindMaybe gopurs_runtime.Value
var once_Data_Maybe_bindMaybe sync.Once
func Get_Data_Maybe_bindMaybe() gopurs_runtime.Value {
	once_Data_Maybe_bindMaybe.Do(func() {
		cache_Data_Maybe_bindMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3887487416_2748095225((&Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3552963512_3741347833(Rebox_Data_Maybe_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
_ = __t_tag_0
if (__t_tag_0 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
})})))}
	})
	return cache_Data_Maybe_bindMaybe
}

var cache_Data_Maybe_semiringMaybe gopurs_runtime.Value
var once_Data_Maybe_semiringMaybe sync.Once
func Get_Data_Maybe_semiringMaybe() gopurs_runtime.Value {
	once_Data_Maybe_semiringMaybe.Do(func() {
		cache_Data_Maybe_semiringMaybe = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Maybe_semiringMaybe(dictSemiring_0_box)
})
	})
	return cache_Data_Maybe_semiringMaybe
}

var cache_Data_Maybe_applicativeMaybe gopurs_runtime.Value
var once_Data_Maybe_applicativeMaybe sync.Once
func Get_Data_Maybe_applicativeMaybe() gopurs_runtime.Value {
	once_Data_Maybe_applicativeMaybe.Do(func() {
		cache_Data_Maybe_applicativeMaybe = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_649684152_1439734649((&Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3552963512_3741347833(Rebox_Data_Maybe_3741347833_3552963512(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Maybe_applyMaybe()))))}
}), Get_Data_Maybe_Just()})))}
	})
	return cache_Data_Maybe_applicativeMaybe
}

var cache_Data_Maybe_monadMaybe gopurs_runtime.Value
var once_Data_Maybe_monadMaybe sync.Once
func Get_Data_Maybe_monadMaybe() gopurs_runtime.Value {
	once_Data_Maybe_monadMaybe.Do(func() {
		cache_Data_Maybe_monadMaybe = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1642601656_2568689657((&Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_649684152_1439734649(Rebox_Data_Maybe_1439734649_649684152(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Maybe_applicativeMaybe()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3887487416_2748095225(Rebox_Data_Maybe_2748095225_3887487416(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Maybe_bindMaybe()))))}
})})))}
	})
	return cache_Data_Maybe_monadMaybe
}

var cache_Data_Maybe_altMaybe gopurs_runtime.Value
var once_Data_Maybe_altMaybe sync.Once
func Get_Data_Maybe_altMaybe() gopurs_runtime.Value {
	once_Data_Maybe_altMaybe.Do(func() {
		cache_Data_Maybe_altMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3904200120_3421983481((&Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3689823567_2812149806(Rebox_Data_Maybe_2812149806_3689823567(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Maybe_functorMaybe()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_1)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
})})))}
	})
	return cache_Data_Maybe_altMaybe
}

var cache_Data_Maybe_plusMaybe gopurs_runtime.Value
var once_Data_Maybe_plusMaybe sync.Once
func Get_Data_Maybe_plusMaybe() gopurs_runtime.Value {
	once_Data_Maybe_plusMaybe.Do(func() {
		cache_Data_Maybe_plusMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_267397720_3706288089((&Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3904200120_3421983481(Rebox_Data_Maybe_3421983481_3904200120(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Get_Data_Maybe_altMaybe()))))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}})))}
	})
	return cache_Data_Maybe_plusMaybe
}

var cache_Data_Maybe_alternativeMaybe gopurs_runtime.Value
var once_Data_Maybe_alternativeMaybe sync.Once
func Get_Data_Maybe_alternativeMaybe() gopurs_runtime.Value {
	once_Data_Maybe_alternativeMaybe.Do(func() {
		cache_Data_Maybe_alternativeMaybe = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_673668088_2307501113((&Constructor_Control_Alternative_Alternative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_649684152_1439734649(Rebox_Data_Maybe_1439734649_649684152(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Maybe_applicativeMaybe()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_267397720_3706288089(Rebox_Data_Maybe_3706288089_267397720(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_Maybe_plusMaybe()))))}
})})))}
	})
	return cache_Data_Maybe_alternativeMaybe
}

type Constructor_Data_Maybe_Nothing[T_a any] struct {
	Rc uint32
}


type Constructor_Data_Maybe_Just[T_a any] struct {
	Rc uint32
	V0 T_a
}


func Call_Data_Maybe_Just__237222391(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__237222391:
for {
if false { continue Just__237222391 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_April().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__3686219312(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3686219312:
for {
if false { continue Just__3686219312 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_August().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__2614824046(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2614824046:
for {
if false { continue Just__2614824046 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_December().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__3785071805(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3785071805:
for {
if false { continue Just__3785071805 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_February().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__2633198667(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2633198667:
for {
if false { continue Just__2633198667 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_January().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__105528475(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__105528475:
for {
if false { continue Just__105528475 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_July().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__4249312965(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__4249312965:
for {
if false { continue Just__4249312965 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_June().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__4573476(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__4573476:
for {
if false { continue Just__4573476 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_March().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__496498436(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__496498436:
for {
if false { continue Just__496498436 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_May().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__3591011291(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3591011291:
for {
if false { continue Just__3591011291 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_November().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__139360115(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__139360115:
for {
if false { continue Just__139360115 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_October().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__1605594846(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__1605594846:
for {
if false { continue Just__1605594846 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_September().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__2013662992(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2013662992:
for {
if false { continue Just__2013662992 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Friday().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__764130881(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__764130881:
for {
if false { continue Just__764130881 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Monday().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__2900182732(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2900182732:
for {
if false { continue Just__2900182732 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Saturday().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__1960558085(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__1960558085:
for {
if false { continue Just__1960558085 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Sunday().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__555433957(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__555433957:
for {
if false { continue Just__555433957 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Thursday().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__98062938(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__98062938:
for {
if false { continue Just__98062938 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Tuesday().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__3422863715(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3422863715:
for {
if false { continue Just__3422863715 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_Wednesday().IntVal)), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__165672699(__eta_norm_0_0_loop *Constructor_Data_Date_Date) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__165672699:
for {
if false { continue Just__165672699 }
var __eta_norm_0_0 *Constructor_Data_Date_Date = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_2280409795_3094389156(Rebox_Data_Maybe_3094389156_2280409795(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(__eta_norm_0_0)}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__1852526845(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__1852526845:
for {
if false { continue Just__1852526845 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__819016571(__eta_norm_0_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__819016571:
for {
if false { continue Just__819016571 }
var __eta_norm_0_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_470062885_3094389156(Rebox_Data_Maybe_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3123684004_1293498952(__eta_norm_0_0))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__3139849822(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3139849822:
for {
if false { continue Just__3139849822 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__1323530841(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__1323530841:
for {
if false { continue Just__1323530841 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__821644466(__eta_norm_0_unused_0_loop uint32) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__821644466:
for {
if false { continue Just__821644466 }
var __eta_norm_0_unused_0 uint32 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_622082505_3094389156(Rebox_Data_Maybe_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__1425602363(__eta_norm_0_0_loop *Constructor_Data_Tuple_Tuple[int64, int64]) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__1425602363:
for {
if false { continue Just__1425602363 }
var __eta_norm_0_0 *Constructor_Data_Tuple_Tuple[int64, int64] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1415037225_3094389156(Rebox_Data_Maybe_3094389156_1415037225(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3363075976_138441832(__eta_norm_0_0))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__635767739(__eta_norm_0_0_loop bool) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__635767739:
for {
if false { continue Just__635767739 }
var __eta_norm_0_0 bool = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1044698560_3094389156(Rebox_Data_Maybe_3094389156_1044698560(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Bool(__eta_norm_0_0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__2688222971(__eta_norm_0_0_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2688222971:
for {
if false { continue Just__2688222971 }
var __eta_norm_0_0 string = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_742090555_3094389156(Rebox_Data_Maybe_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Str(__eta_norm_0_0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__2649619419(__eta_norm_0_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2649619419:
for {
if false { continue Just__2649619419 }
var __eta_norm_0_0 int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1170268447_3094389156(Rebox_Data_Maybe_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(__eta_norm_0_0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__2582300155(__eta_norm_0_0_loop float64) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__2582300155:
for {
if false { continue Just__2582300155 }
var __eta_norm_0_0 float64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3240988860_3094389156(Rebox_Data_Maybe_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Float(__eta_norm_0_0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__651901051(__eta_norm_0_0_loop struct{
	head string
	tail string
}) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__651901051:
for {
if false { continue Just__651901051 }
var __eta_norm_0_0 struct{
	head string
	tail string
} = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_2377819597_3094389156(Rebox_Data_Maybe_3094389156_2377819597(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{func() gopurs_runtime.Value {
				orig := __eta_norm_0_0
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str(orig.head), gopurs_runtime.Str(orig.tail))
				}(), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__1601057723(__eta_norm_0_0_loop struct{
	head int64
	tail string
}) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__1601057723:
for {
if false { continue Just__1601057723 }
var __eta_norm_0_0 struct{
	head int64
	tail string
} = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1387998409_3094389156(Rebox_Data_Maybe_3094389156_1387998409(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{func() gopurs_runtime.Value {
				orig := __eta_norm_0_0
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Str(orig.tail))
				}(), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__1492091643(__eta_norm_0_0_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__1492091643:
for {
if false { continue Just__1492091643 }
var __eta_norm_0_0 string = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_742090555_3094389156(Rebox_Data_Maybe_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Str(__eta_norm_0_0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_Just__3221933828(__eta_norm_0_unused_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
Just__3221933828:
for {
if false { continue Just__3221933828 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{Get_Data_Unit_unit(), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_showMaybe(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_345859663_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 string
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
_ = __t_tag_0
if (__t_tag_0 != nil) {
__t2 = (("(Just ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t2 = "Nothing"
goto end_branch_2
} else {

}
}
{
__t2 = func() string { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Str(__t2)
})})))}
}

func Call_Data_Maybe_semigroupMaybe(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_125234255_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
goto end_branch_5
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
goto end_branch_5
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_1)
_ = __t_tag_2
var __t_and_4 bool = false
if (__t_tag_2 != nil) {

var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_2)
_ = __t_tag_3
__t_and_4 = (__t_tag_3 != nil)
}
if __t_and_4 {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_2.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
})})))}
}

func Call_Data_Maybe_optional(dictAlt_0_loop *Constructor_Control_Alt_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlt_0 *Constructor_Control_Alt_Alt[gopurs_runtime.Value] = dictAlt_0_loop
_ = dictAlt_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope25)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlt_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(dictApplicative_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictAlt_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, Get_Data_Maybe_Just(), a_3), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}))
})
}

func Call_Data_Maybe_monoidMaybe(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): semigroupMaybe1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope30)])])
semigroupMaybe1_1_0 := Rebox_Data_Maybe_4179793454_125234255(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Maybe_semigroupMaybe(dictSemigroup_0)))
_ = semigroupMaybe1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_1089607855_1201789390((&Constructor_Data_Monoid_Monoid[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_125234255_4179793454(semigroupMaybe1_1_0))}
}), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})})))}
}

func Call_Data_Maybe_maybe_prime_(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_maybe__4265365554(v_unused_0_loop uint32, v1_unused_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[uint32]) uint32 {
maybe__4265365554:
for {
if false { continue maybe__4265365554 }
var v_unused_0 uint32 = v_unused_0_loop
_ = v_unused_0
var v1_unused_1 gopurs_runtime.Value = v1_unused_1_loop
_ = v1_unused_1
var v2_2 *Constructor_Data_Maybe_Just[uint32] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_December().IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: int64((v2_2).V0), UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return uint32(__t0.IntVal)
}
}

func Call_Data_Maybe_maybe__1711083959(v_unused_0_loop uint32, v1_unused_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[uint32]) uint32 {
maybe__1711083959:
for {
if false { continue maybe__1711083959 }
var v_unused_0 uint32 = v_unused_0_loop
_ = v_unused_0
var v1_unused_1 gopurs_runtime.Value = v1_unused_1_loop
_ = v1_unused_1
var v2_2 *Constructor_Data_Maybe_Just[uint32] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_January().IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: int64((v2_2).V0), UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return uint32(__t0.IntVal)
}
}

func Call_Data_Maybe_maybe__2146642512(v_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
maybe__2146642512:
for {
if false { continue maybe__2146642512 }
var v_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (v2_2).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := __t0
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_Maybe_maybe__2726997443(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[uint32]) bool {
maybe__2726997443:
for {
if false { continue maybe__2726997443 }
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[uint32] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Bool(v_0)
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Value{Type: 9, IntVal: int64((v2_2).V0), UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}
}

func Call_Data_Maybe_maybe__4252741795(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[int64]) bool {
maybe__4252741795:
for {
if false { continue maybe__4252741795 }
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[int64] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Bool(v_0)
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Int((v2_2).V0))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}
}

func Call_Data_Maybe_maybe__3642269347(v_0_loop bool, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[float64]) bool {
maybe__3642269347:
for {
if false { continue maybe__3642269347 }
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[float64] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Bool(v_0)
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Float((v2_2).V0))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}
}

func Call_Data_Maybe_maybe__2295219908(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[int64]) int64 {
maybe__2295219908:
for {
if false { continue maybe__2295219908 }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[int64] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Int(v_0)
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Int((v2_2).V0))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}
}

func Call_Data_Maybe_maybe__3140274526(v_0_loop int64, v1_unused_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[int64]) int64 {
maybe__3140274526:
for {
if false { continue maybe__3140274526 }
var v_0 int64 = v_0_loop
_ = v_0
var v1_unused_1 gopurs_runtime.Value = v1_unused_1_loop
_ = v1_unused_1
var v2_2 *Constructor_Data_Maybe_Just[int64] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Int(v_0)
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Int((v2_2).V0))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}
}

func Call_Data_Maybe_maybe__1210512692(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just[int64]) int64 {
maybe__1210512692:
for {
if false { continue maybe__1210512692 }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just[int64] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = gopurs_runtime.Int(v_0)
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply(v1_1, gopurs_runtime.Int((v2_2).V0))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}
}

func Call_Data_Maybe_isNothing(v2_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) bool {
var v2_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_Data_Maybe_isNothing__2716061749(__eta_norm_0_0_loop *Constructor_Data_Maybe_Just[uint32]) bool {
isNothing__2716061749:
for {
if false { continue isNothing__2716061749 }
var __eta_norm_0_0 *Constructor_Data_Maybe_Just[uint32] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0 != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}
}

func Call_Data_Maybe_isNothing__1937207052(__eta_norm_0_0_loop *Constructor_Data_Maybe_Just[int64]) bool {
isNothing__1937207052:
for {
if false { continue isNothing__1937207052 }
var __eta_norm_0_0 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_0 == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0 != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}
}

func Call_Data_Maybe_isJust(v2_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) bool {
var v2_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (v2_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (v2_0 != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_Data_Maybe_isJust__1937207052(__eta_norm_0_0_loop *Constructor_Data_Maybe_Just[int64]) bool {
isJust__1937207052:
for {
if false { continue isJust__1937207052 }
var __eta_norm_0_0 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0 != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}
}

func Call_Data_Maybe_isJust__3089511036(__eta_norm_0_0_loop *Constructor_Data_Maybe_Just[float64]) bool {
isJust__3089511036:
for {
if false { continue isJust__3089511036 }
var __eta_norm_0_0 *Constructor_Data_Maybe_Just[float64] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_0 == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_0 != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}
}

func Call_Data_Maybe_fromMaybe_prime_(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply2(Get_Data_Maybe_maybe_prime_(), a_0, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Call_Data_Maybe_fromMaybe(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope77)] (TypeVar a$scope77))
__local_var_1_0 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_1_0
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t3 = a_0
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2)
_ = __t_tag_2
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
}

func Call_Data_Maybe_fromMaybe__3286440201(a_unused_0_loop uint32, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[uint32]) uint32 {
fromMaybe__3286440201:
for {
if false { continue fromMaybe__3286440201 }
var a_unused_0 uint32 = a_unused_0_loop
_ = a_unused_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[uint32] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_1 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_December().IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_1 != nil) {
__t0 = gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: int64((__eta_norm_0_1).V0), UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return uint32(__t0.IntVal)
}
}

func Call_Data_Maybe_fromMaybe__516259212(a_unused_0_loop uint32, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[uint32]) uint32 {
fromMaybe__516259212:
for {
if false { continue fromMaybe__516259212 }
var a_unused_0 uint32 = a_unused_0_loop
_ = a_unused_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[uint32] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_1 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Date_Component_January().IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_1 != nil) {
__t0 = gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Value{Type: 9, IntVal: int64((__eta_norm_0_1).V0), UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return uint32(__t0.IntVal)
}
}

func Call_Data_Maybe_fromMaybe__816552133(a_0_loop int64, __eta_norm_0_1_loop *Constructor_Data_Maybe_Just[int64]) int64 {
fromMaybe__816552133:
for {
if false { continue fromMaybe__816552133 }
var a_0 int64 = a_0_loop
_ = a_0
var __eta_norm_0_1 *Constructor_Data_Maybe_Just[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t0 gopurs_runtime.Value
{
if (__eta_norm_0_1 == nil) {
__t0 = gopurs_runtime.Int(a_0)
goto end_branch_0
} else {

}
}
{
if (__eta_norm_0_1 != nil) {
__t0 = gopurs_runtime.Apply(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), gopurs_runtime.Int((__eta_norm_0_1).V0))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}
}

func Call_Data_Maybe_fromJust(_dollar___unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var v_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (v_1 != nil) {
__t0 = (v_1).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_Data_Maybe_fromJust__2539482(v_0_loop *Constructor_Data_Maybe_Just[uint32]) uint32 {
fromJust__2539482:
for {
if false { continue fromJust__2539482 }
var v_0 *Constructor_Data_Maybe_Just[uint32] = v_0_loop
_ = v_0
var __t0 uint32
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_fromJust__1001929434(v_0_loop *Constructor_Data_Maybe_Just[uint32]) uint32 {
fromJust__1001929434:
for {
if false { continue fromJust__1001929434 }
var v_0 *Constructor_Data_Maybe_Just[uint32] = v_0_loop
_ = v_0
var __t0 uint32
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_fromJust__171824954(v_0_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Date_Date {
fromJust__171824954:
for {
if false { continue fromJust__171824954 }
var v_0 *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Date_Date
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() *Constructor_Data_Date_Date { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_fromJust__3704324326(v_0_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]) *Constructor_Data_Time_Time {
fromJust__3704324326:
for {
if false { continue fromJust__3704324326 }
var v_0 *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Time_Time
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() *Constructor_Data_Time_Time { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_fromJust__55301350(v_0_loop *Constructor_Data_Maybe_Just[[]string]) []string {
fromJust__55301350:
for {
if false { continue fromJust__55301350 }
var v_0 *Constructor_Data_Maybe_Just[[]string] = v_0_loop
_ = v_0
var __t0 []string
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() []string { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_fromJust__1673906054(v_0_loop *Constructor_Data_Maybe_Just[[]int64]) []int64 {
fromJust__1673906054:
for {
if false { continue fromJust__1673906054 }
var v_0 *Constructor_Data_Maybe_Just[[]int64] = v_0_loop
_ = v_0
var __t0 []int64
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() []int64 { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_fromJust__218925574(v_0_loop *Constructor_Data_Maybe_Just[int64]) int64 {
fromJust__218925574:
for {
if false { continue fromJust__218925574 }
var v_0 *Constructor_Data_Maybe_Just[int64] = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() int64 { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_fromJust__2463559942(v_0_loop *Constructor_Data_Maybe_Just[string]) string {
fromJust__2463559942:
for {
if false { continue fromJust__2463559942 }
var v_0 *Constructor_Data_Maybe_Just[string] = v_0_loop
_ = v_0
var __t0 string
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() string { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_Maybe_eqMaybe(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3508461103_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 bool
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1)
_ = __t_tag_3
if (__t_tag_3 == nil) {
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_2)
_ = __t_tag_4
__t5 = (__t_tag_4 == nil)
goto end_branch_5
} else {

}
}
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_1)
_ = __t_tag_0
var __t_and_2 bool = false
if (__t_tag_0 != nil) {

var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_2)
_ = __t_tag_1
__t_and_2 = ((__t_tag_1 != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0))
}
__t5 = __t_and_2
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
})})))}
}

func Call_Data_Maybe_ordMaybe(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqMaybe1_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope18)])])
eqMaybe1_1_0 := Rebox_Data_Maybe_3790796878_3508461103(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Maybe_eqMaybe(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))))
_ = eqMaybe1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_2155612431_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3508461103_3790796878(eqMaybe1_1_0))}
}), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 uint32
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_2)
_ = __t_tag_1
if (__t_tag_1 == nil) {
var __t3 uint32
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_3)
_ = __t_tag_2
if (__t_tag_2 == nil) {
__t3 = 902936544
goto end_branch_3
} else {

}
}
{
__t3 = 1527465420
}
end_branch_3:
__t8 = __t3
goto end_branch_8
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_3)
_ = __t_tag_4
if (__t_tag_4 == nil) {
__t8 = 380165415
goto end_branch_8
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_2)
_ = __t_tag_5
var __t_and_7 bool = false
if (__t_tag_5 != nil) {

var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_3)
_ = __t_tag_6
__t_and_7 = (__t_tag_6 != nil)
}
if __t_and_7 {
__t8 = uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal)
goto end_branch_8
} else {

}
}
{
__t8 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t8), UnsafePtr: nil}
})})))}
}

func Call_Data_Maybe_boundedMaybe(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): ordMaybe1_1_0 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope95)])])
ordMaybe1_1_0 := Rebox_Data_Maybe_4177771502_2155612431(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Maybe_ordMaybe(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{}))))
_ = ordMaybe1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_591333647_2094947566((&Constructor_Data_Bounded_Bounded[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_2155612431_4177771502(ordMaybe1_1_0))}
}), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordGet(dictBounded_0, "top"), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())})))}
}

func Call_Data_Maybe_semiringMaybe(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
// TAST (Let): mul_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope4), (TypeVar a$scope4)] (TypeVar a$scope4))
mul_1_0 := Call_Data_Semiring_mul(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](dictSemiring_0))
_ = mul_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_2047434479_2826095630((&Constructor_Data_Semiring_Semiring[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
goto end_branch_6
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
_ = __t_tag_2
if (__t_tag_2 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
goto end_branch_6
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_3
var __t_and_5 bool = false
if (__t_tag_3 != nil) {

var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
_ = __t_tag_4
__t_and_5 = (__t_tag_4 != nil)
}
if __t_and_5 {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
}), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](x_2)
_ = __t_tag_7
if (__t_tag_7 != nil) {
var __t9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](y_3)
_ = __t_tag_8
if (__t_tag_8 != nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply2(mul_1_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(y_3.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_9:
__t10 = __t9
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)}
}), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordGet(dictSemiring_0, "one"), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})})))}
}

func Rebox_Data_Maybe_1044698560_3094389156(in *Constructor_Data_Maybe_Just[bool]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Bool(in.V0)
	return out
}

func Rebox_Data_Maybe_1089607855_1201789390(in *Constructor_Data_Monoid_Monoid[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Maybe_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Maybe_125234255_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_1293498952_3123684004(in *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_138441832_3363075976(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[int64, int64]{}
		out.V0 = in.V0.IntVal
		out.V1 = in.V1.IntVal
	return out
}

func Rebox_Data_Maybe_1387998409_3094389156(in *Constructor_Data_Maybe_Just[struct{
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

func Rebox_Data_Maybe_1415037225_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3363075976_138441832(in.V0))}
	return out
}

func Rebox_Data_Maybe_1439734649_649684152(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_1642601656_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Monad[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_1662389854_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_170825214_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_1766074591_1662389854(in *Constructor_Data_Eq_Eq1[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_2047434479_2826095630(in *Constructor_Data_Semiring_Semiring[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V2)}
		out.V3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V3)}
	return out
}

func Rebox_Data_Maybe_2155612431_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_2280409795_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Maybe_2377819597_3094389156(in *Constructor_Data_Maybe_Just[struct{
	head string
	tail string
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str(orig.head), gopurs_runtime.Str(orig.tail))
				}()
	return out
}

func Rebox_Data_Maybe_267397720_3706288089(in *Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Plus_Plus[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Plus_Plus[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_2748095225_3887487416(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_2812149806_3689823567(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_3094389156_1044698560(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[bool] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[bool]{}
		out.V0 = (in.V0.IntVal) != (0)
	return out
}

func Rebox_Data_Maybe_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Maybe_3094389156_1387998409(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
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

func Rebox_Data_Maybe_3094389156_1415037225(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]{}
		out.V0 = Rebox_Data_Maybe_138441832_3363075976(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_Maybe_3094389156_1495236409(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[[]int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[[]int64]{}
		out.V0 = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(in.V0.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
	return out
}

func Rebox_Data_Maybe_3094389156_1731162461(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[[]string] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[[]string]{}
		out.V0 = func() []string {
					arr := *(*[]gopurs_runtime.Value)(in.V0.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
	return out
}

func Rebox_Data_Maybe_3094389156_2280409795(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Date_Date]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](in.V0)
	return out
}

func Rebox_Data_Maybe_3094389156_2377819597(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
	head string
	tail string
}] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[struct{
	head string
	tail string
}]{}
		out.V0 = func() struct{
	head string
	tail string
} {
					orig := in.V0
					_ = orig
					clone := struct{
	head string
	tail string
}{}
					clone.head = gopurs_runtime.RecordGet(orig, "head").StrVal()
					clone.tail = gopurs_runtime.RecordGet(orig, "tail").StrVal()
					return clone
				}()
	return out
}

func Rebox_Data_Maybe_3094389156_3240988860(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[float64]{}
		out.V0 = in.V0.FloatVal()
	return out
}

func Rebox_Data_Maybe_3094389156_3839235747(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Time_Time]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](in.V0)
	return out
}

func Rebox_Data_Maybe_3094389156_470062885(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = Rebox_Data_Maybe_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_Maybe_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[uint32]{}
		out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Data_Maybe_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[string]{}
		out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_Maybe_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_3240988860_3094389156(in *Constructor_Data_Maybe_Just[float64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Float(in.V0)
	return out
}

func Rebox_Data_Maybe_3363075976_138441832(in *Constructor_Data_Tuple_Tuple[int64, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Data_Maybe_3421983481_3904200120(in *Constructor_Control_Alt_Alt[gopurs_runtime.Value]) *Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_345859663_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_3508461103_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_3552963512_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_3689823567_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_3706288089_267397720(in *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) *Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Plus_Plus[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_3741347833_3552963512(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_3790796878_3508461103(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_3818752775_2241343270(in *Constructor_Data_Functor_Invariant_Invariant[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_3887487416_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_3904200120_3421983481(in *Constructor_Control_Alt_Alt[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Alt_Alt[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Alt_Alt[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_4177771502_2155612431(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_4179793454_125234255(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_470062885_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_Maybe_3123684004_1293498952(in.V0))}
	return out
}

func Rebox_Data_Maybe_497217223_1323331594(in *Constructor_Data_Generic_Rep_Inl[uint32, gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Generic_Rep_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Data_Maybe_570732504_3290176857(in *Constructor_Control_Extend_Extend[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Extend_Extend[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Extend_Extend[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_591333647_2094947566(in *Constructor_Data_Bounded_Bounded[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V2)}
	return out
}

func Rebox_Data_Maybe_622082505_3094389156(in *Constructor_Data_Maybe_Just[uint32]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}

func Rebox_Data_Maybe_649684152_1439734649(in *Constructor_Control_Applicative_Applicative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_660678937_2687169876(in *Constructor_Data_Generic_Rep_Inr[uint32, gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Generic_Rep_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Maybe_673668088_2307501113(in *Constructor_Control_Alternative_Alternative[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_732223057_2818661616(in *Constructor_Data_Generic_Rep_Generic[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Generic_Rep_Generic[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Maybe_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Str(in.V0)
	return out
}


