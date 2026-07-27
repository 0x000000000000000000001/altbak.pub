package Effect_Unsafe

func UnsafePerformEffect(f func() interface{}) interface{} { return f() }
