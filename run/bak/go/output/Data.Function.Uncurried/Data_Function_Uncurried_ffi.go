package Data_Function_Uncurried




func identity(fn any) any {
	return fn
}

var RunFn2 = identity
var RunFn3 = identity
var RunFn4 = identity
var RunFn5 = identity
var RunFn6 = identity
var RunFn7 = identity
var RunFn8 = identity
var RunFn9 = identity
var RunFn10 = identity

var MkFn2 = identity
var MkFn3 = identity
var MkFn4 = identity
var MkFn5 = identity
var MkFn6 = identity
var MkFn7 = identity
var MkFn8 = identity
var MkFn9 = identity
var MkFn10 = identity

func MkFn0(f any) any {
	return f
}

func RunFn0(f func(any) any) any {
	return f(nil)
}
