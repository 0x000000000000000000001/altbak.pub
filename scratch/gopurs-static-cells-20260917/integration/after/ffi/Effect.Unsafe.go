package Effect_Unsafe

func UnsafePerformEffect(f func(interface{}) interface{}) interface{} {
    return f(nil)
}
