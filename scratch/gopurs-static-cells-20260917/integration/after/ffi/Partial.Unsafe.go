func _UnsafePartial(f func(interface{}) interface{}) interface{} {
	return f(nil)
}
