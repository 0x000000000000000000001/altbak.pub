package Effect_Ref

func _New(val interface{}) func() interface{} {
	return func() interface{} {
		return map[string]interface{}{"value": val}
	}
}
func NewWithSelf(f func(interface{}) interface{}) func() interface{} {
	return func() interface{} {
		ref := map[string]interface{}{}
		ref["value"] = f(ref)
		return ref
	}
}
func Read(ref interface{}) func() interface{} {
	return func() interface{} {
		return ref.(map[string]interface{})["value"]
	}
}
func ModifyImpl(f func(interface{}) interface{}, ref interface{}) func() interface{} {
	return func() interface{} {
		t := f(ref.(map[string]interface{})["value"]).(map[string]interface{})
		ref.(map[string]interface{})["value"] = t["state"]
		return t["value"]
	}
}
func Write(val interface{}, ref interface{}) func() interface{} {
	return func() interface{} {
		ref.(map[string]interface{})["value"] = val
		return nil
	}
}
