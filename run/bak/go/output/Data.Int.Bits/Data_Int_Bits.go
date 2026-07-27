package Data_Int_Bits

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_and gopurs_runtime.Value
var once_and sync.Once
func Get_and() gopurs_runtime.Value {
	once_and.Do(func() {
		cache_and = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(And(arg0.IntVal, arg1.IntVal))
})
	})
	return cache_and
}

var cache_complement gopurs_runtime.Value
var once_complement sync.Once
func Get_complement() gopurs_runtime.Value {
	once_complement.Do(func() {
		cache_complement = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Complement(arg0.IntVal))
})
	})
	return cache_complement
}

var cache_or gopurs_runtime.Value
var once_or sync.Once
func Get_or() gopurs_runtime.Value {
	once_or.Do(func() {
		cache_or = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Or(arg0.IntVal, arg1.IntVal))
})
	})
	return cache_or
}

var cache_shl gopurs_runtime.Value
var once_shl sync.Once
func Get_shl() gopurs_runtime.Value {
	once_shl.Do(func() {
		cache_shl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Shl(arg0.IntVal, arg1.IntVal))
})
	})
	return cache_shl
}

var cache_shr gopurs_runtime.Value
var once_shr sync.Once
func Get_shr() gopurs_runtime.Value {
	once_shr.Do(func() {
		cache_shr = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Shr(arg0.IntVal, arg1.IntVal))
})
	})
	return cache_shr
}

var cache_xor gopurs_runtime.Value
var once_xor sync.Once
func Get_xor() gopurs_runtime.Value {
	once_xor.Do(func() {
		cache_xor = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Xor(arg0.IntVal, arg1.IntVal))
})
	})
	return cache_xor
}

var cache_zshr gopurs_runtime.Value
var once_zshr sync.Once
func Get_zshr() gopurs_runtime.Value {
	once_zshr.Do(func() {
		cache_zshr = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Zshr(arg0.IntVal, arg1.IntVal))
})
	})
	return cache_zshr
}


