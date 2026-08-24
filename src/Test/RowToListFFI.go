package Test_RowToListFFI


type RecordKeys interface {
	keysImpl(interface{}) int
}

type dictNil struct{}
func (dictNil) keysImpl(_ interface{}) int {
	return 0
}

type dictCons struct {
	tail RecordKeys
}
func (d dictCons) keysImpl(_ interface{}) int {
	return 1 + d.tail.keysImpl(nil)
}

func RunRowToListFFI(limit float64) float64 {
	// dummy := int(limit)
	// rec is not even used in keysImpl, it's just for the type signature
	dict := dictCons{tail: dictCons{tail: dictCons{tail: dictCons{tail: dictCons{tail: dictNil{}}}}}}
	return float64(dict.keysImpl(nil))
}
