package Data_Array_ST_Partial

func PeekImpl(i int, xs []interface{}) func() interface{} { return func() interface{} { return xs[i] } }
func PokeImpl(i int, a interface{}, xs []interface{}) func() bool { return func() bool { xs[i] = a; return true } }
