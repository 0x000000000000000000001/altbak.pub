func OrdBooleanImpl(lt interface{}, eq interface{}, gt interface{}, x bool, y bool) interface{} {
	if !x && y {
		return lt
	} else if x == y {
		return eq
	}
	return gt
}
func OrdIntImpl(lt interface{}, eq interface{}, gt interface{}, x int64, y int64) interface{} {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdCharImpl(lt interface{}, eq interface{}, gt interface{}, x string, y string) interface{} {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdStringImpl(lt interface{}, eq interface{}, gt interface{}, x string, y string) interface{} {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdNumberImpl(lt interface{}, eq interface{}, gt interface{}, x float64, y float64) interface{} {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdArrayImpl(f func(interface{}, interface{}) int64, xs []interface{}, ys []interface{}) int64 {
	xlen := len(xs)
	ylen := len(ys)
	for i := 0; i < xlen && i < ylen; i++ {
		o := f(xs[i], ys[i])
		if o != 0 {
			return o
		}
	}
	if xlen == ylen {
		return 0
	} else if xlen > ylen {
		return -1
	} else {
		return 1
	}
}
