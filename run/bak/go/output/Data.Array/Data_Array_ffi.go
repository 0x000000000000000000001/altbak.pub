package Data_Array



func RangeImpl(start int, end int) []int {
	step := 1
	if start > end {
		step = -1
	}
	size := (end - start) * step + 1
	result := make([]int, size)
	i := start
	n := 0
	for i != end {
		result[n] = i
		n++
		i += step
	}
	result[n] = i
	return result
}

func ReplicateImpl(count int, value interface{}) []interface{} {
	if count < 1 {
		return make([]interface{}, 0)
	}
	result := make([]interface{}, count)
	for i := 0; i < count; i++ {
		result[i] = value
	}
	return result
}

func Length(xs []interface{}) int {
	return len(xs)
}

func UnconsImpl(empty func(interface{}) interface{}, next func(interface{}) func([]interface{}) interface{}, xs []interface{}) interface{} {
	if len(xs) == 0 {
		return empty(nil)
	}
	head := xs[0]
	tail := make([]interface{}, len(xs)-1)
	copy(tail, xs[1:])
	return next(head)(tail)
}

func IndexImpl(just func(interface{}) interface{}, nothing interface{}, xs []interface{}, i int) interface{} {
	if i < 0 || i >= len(xs) {
		return nothing
	}
	return just(xs[i])
}

func _UpdateAt(just func([]interface{}) interface{}, nothing interface{}, i int, a interface{}, xs []interface{}) interface{} {
	if i < 0 || i >= len(xs) {
		return nothing
	}
	l1 := make([]interface{}, len(xs))
	copy(l1, xs)
	l1[i] = a
	return just(l1)
}

func _InsertAt(just func([]interface{}) interface{}, nothing interface{}, i int, a interface{}, xs []interface{}) interface{} {
	if i < 0 || i > len(xs) {
		return nothing
	}
	l1 := make([]interface{}, 0, len(xs)+1)
	l1 = append(l1, xs[:i]...)
	l1 = append(l1, a)
	l1 = append(l1, xs[i:]...)
	return just(l1)
}

func _DeleteAt(just func([]interface{}) interface{}, nothing interface{}, i int, xs []interface{}) interface{} {
	if i < 0 || i >= len(xs) {
		return nothing
	}
	l1 := make([]interface{}, 0, len(xs)-1)
	l1 = append(l1, xs[:i]...)
	l1 = append(l1, xs[i+1:]...)
	return just(l1)
}

func Reverse(xs []interface{}) []interface{} {
	l := len(xs)
	l1 := make([]interface{}, l)
	for i := 0; i < l; i++ {
		l1[i] = xs[l-1-i]
	}
	return l1
}

func Concat(xss [][]interface{}) []interface{} {
	var result []interface{}
	for _, xs := range xss {
		result = append(result, xs...)
	}
	return result
}

func FilterImpl(f func(interface{}) bool, xs []interface{}) []interface{} {
	var result []interface{}
	for _, x := range xs {
		if f(x) {
			result = append(result, x)
		}
	}
	return result
}

func SliceImpl(s int, e int, l []interface{}) []interface{} {
	if s < 0 {
		s = len(l) + s
	}
	if e < 0 {
		e = len(l) + e
	}
	if s < 0 { s = 0 }
	if e > len(l) { e = len(l) }
	if s > e { s = e }
	
	res := make([]interface{}, e-s)
	copy(res, l[s:e])
	return res
}

func ZipWithImpl(f func(interface{}) func(interface{}) interface{}, xs []interface{}, ys []interface{}) []interface{} {
	length := len(xs)
	if len(ys) < length {
		length = len(ys)
	}
	result := make([]interface{}, length)
	for i := 0; i < length; i++ {
		result[i] = f(xs[i])(ys[i])
	}
	return result
}

func UnsafeIndexImpl(xs []interface{}, n int) interface{} {
	return xs[n]
}

func SortByImpl(compare func(interface{}) func(interface{}) interface{}, fromOrdering func(interface{}) int, xs []interface{}) []interface{} {
	if len(xs) < 2 {
		return xs
	}
	out := make([]interface{}, len(xs))
	copy(out, xs)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			c := fromOrdering(compare(out[i])(out[j]))
			if c > 0 { // GT
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}

func ScanrImpl(f func(interface{}) func(interface{}) interface{}, b interface{}, xs []interface{}) []interface{} {
	out := make([]interface{}, len(xs))
	acc := b
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(xs[i])(acc)
		out[i] = acc
	}
	return out
}

func ScanlImpl(f func(interface{}) func(interface{}) interface{}, b interface{}, xs []interface{}) []interface{} {
	out := make([]interface{}, len(xs))
	acc := b
	for i := 0; i < len(xs); i++ {
		acc = f(acc)(xs[i])
		out[i] = acc
	}
	return out
}

func PartitionImpl(f func(interface{}) bool, xs []interface{}) map[string]interface{} {
	var yes []interface{}
	var no []interface{}
	for _, x := range xs {
		if f(x) {
			yes = append(yes, x)
		} else {
			no = append(no, x)
		}
	}
	return map[string]interface{}{
		"yes": yes,
		"no":  no,
	}
}

func FromFoldableImpl(foldr interface{}, xsVal interface{}) []interface{} {
	panic("Not implemented: FromFoldableImpl (complex callback)")
}

func FindMapImpl(nothing interface{}, isJust func(interface{}) bool, f func(interface{}) interface{}, xs []interface{}) interface{} {
	for _, x := range xs {
		res := f(x)
		if isJust(res) {
			return res
		}
	}
	return nothing
}

func FindLastIndexImpl(just func(int) interface{}, nothing interface{}, f func(interface{}) bool, xs []interface{}) interface{} {
	for i := len(xs) - 1; i >= 0; i-- {
		if f(xs[i]) {
			return just(i)
		}
	}
	return nothing
}

func FindIndexImpl(just func(int) interface{}, nothing interface{}, f func(interface{}) bool, xs []interface{}) interface{} {
	for i := 0; i < len(xs); i++ {
		if f(xs[i]) {
			return just(i)
		}
	}
	return nothing
}

func AnyImpl(p func(interface{}) bool, xs []interface{}) bool {
	for _, x := range xs {
		if p(x) {
			return true
		}
	}
	return false
}

func AllImpl(p func(interface{}) bool, xs []interface{}) bool {
	for _, x := range xs {
		if !p(x) {
			return false
		}
	}
	return true
}
