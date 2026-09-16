package Data_Array_ST_Partial

func PeekImpl(i int64, xs *[]interface{}) interface{} {
	return (*xs)[i]
}
func PokeImpl(i int64, a interface{}, xs *[]interface{}) bool {
	(*xs)[i] = a
	return true
}
