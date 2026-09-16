

func RangeImpl(start int64, end int64) []int64 {
	step := int64(1)
	if start > end {
		step = -1
	}
	size := (end - start) * step + 1
	result := make([]int64, size)
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

func ReplicateImpl(count int64, value interface{}) []interface{} {
	if count < 1 {
		return make([]interface{}, 0)
	}
	result := make([]interface{}, count)
	for i := 0; i < int(count); i++ {
		result[i] = value
	}
	return result
}

func Length(xs []interface{}) int64 {
	return int64(len(xs))
}

func UnconsImpl(empty func(interface{}) interface{}, next func(interface{}, []interface{}) interface{}, xs []interface{}) interface{} {
	if len(xs) == 0 {
		return empty(nil)
	}
	head := xs[0]
	tail := make([]interface{}, len(xs)-1)
	copy(tail, xs[1:])
	return next(head, tail)
}

func IndexImpl(just func(interface{}) interface{}, nothing interface{}, xs []interface{}, i int64) interface{} {
	if i < 0 || int(i) >= len(xs) {
		return nothing
	}
	return just(xs[int(i)])
}

func _UpdateAt(just func([]interface{}) interface{}, nothing interface{}, i int64, a interface{}, xs []interface{}) interface{} {
	if i < 0 || int(i) >= len(xs) {
		return nothing
	}
	l1 := make([]interface{}, len(xs))
	copy(l1, xs)
	l1[int(i)] = a
	return just(l1)
}

func _InsertAt(just func([]interface{}) interface{}, nothing interface{}, i int64, a interface{}, xs []interface{}) interface{} {
	if i < 0 || int(i) > len(xs) {
		return nothing
	}
	l1 := make([]interface{}, 0, len(xs)+1)
	l1 = append(l1, xs[:int(i)]...)
	l1 = append(l1, a)
	l1 = append(l1, xs[int(i):]...)
	return just(l1)
}

func _DeleteAt(just func([]interface{}) interface{}, nothing interface{}, i int64, xs []interface{}) interface{} {
	if i < 0 || int(i) >= len(xs) {
		return nothing
	}
	l1 := make([]interface{}, 0, len(xs)-1)
	l1 = append(l1, xs[:int(i)]...)
	l1 = append(l1, xs[int(i)+1:]...)
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

func SliceImpl(s int64, e int64, l []interface{}) []interface{} {
	sInt := int(s)
	eInt := int(e)
	if sInt < 0 {
		sInt = len(l) + sInt
	}
	if eInt < 0 {
		eInt = len(l) + eInt
	}
	if sInt < 0 { sInt = 0 }
	if eInt > len(l) { eInt = len(l) }
	if sInt > eInt { sInt = eInt }
	
	res := make([]interface{}, eInt-sInt)
	copy(res, l[sInt:eInt])
	return res
}

func ZipWithImpl(f func(interface{}, interface{}) interface{}, xs []interface{}, ys []interface{}) []interface{} {
	length := len(xs)
	if len(ys) < length {
		length = len(ys)
	}
	result := make([]interface{}, length)
	for i := 0; i < length; i++ {
		result[i] = f(xs[i], ys[i])
	}
	return result
}

func UnsafeIndexImpl(xs []interface{}, n int64) interface{} {
	return xs[int(n)]
}

func SortByImpl(compare func(interface{}, interface{}) interface{}, fromOrdering func(interface{}) int64, xs []interface{}) []interface{} {
	if len(xs) < 2 {
		return xs
	}
	out := make([]interface{}, len(xs))
	copy(out, xs)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			c := fromOrdering(compare(out[i], out[j]))
			if c > 0 { // GT
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}

func ScanrImpl(f func(interface{}, interface{}) interface{}, b interface{}, xs []interface{}) []interface{} {
	out := make([]interface{}, len(xs))
	acc := b
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(xs[i], acc)
		out[i] = acc
	}
	return out
}

func ScanlImpl(f func(interface{}, interface{}) interface{}, b interface{}, xs []interface{}) []interface{} {
	out := make([]interface{}, len(xs))
	acc := b
	for i := 0; i < len(xs); i++ {
		acc = f(acc, xs[i])
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

type consList struct {
	head interface{}
	tail interface{}
}

func FromFoldableImpl(foldr func(func(interface{}) func(interface{}) interface{}, interface{}, interface{}) interface{}, xsVal interface{}) []interface{} {
	var emptyList interface{} = nil

	curryCons := func(head interface{}) func(interface{}) interface{} {
		return func(tail interface{}) interface{} {
			return &consList{head: head, tail: tail}
		}
	}

	list := foldr(curryCons, emptyList, xsVal)

	var unboxAny func(interface{}) interface{}
	unboxAny = func(v interface{}) interface{} {
		if val, ok := v.(gopurs_runtime.Value); ok && val.Type == gopurs_runtime.TypeAny {
			if val.UnsafePtr != nil {
				return unboxAny(*(*any)(val.UnsafePtr))
			}
			return nil
		}
		return v
	}

	list = unboxAny(list)

	var result []interface{}
	curr, ok := list.(*consList)
	for ok && curr != nil {
		result = append(result, curr.head)
		curr, ok = unboxAny(curr.tail).(*consList)
	}
	
	if result == nil {
		return make([]interface{}, 0)
	}
	return result
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

func FindLastIndexImpl(just func(int64) interface{}, nothing interface{}, f func(interface{}) bool, xs []interface{}) interface{} {
	for i := len(xs) - 1; i >= 0; i-- {
		if f(xs[i]) {
			return just(int64(i))
		}
	}
	return nothing
}

func FindIndexImpl(just func(int64) interface{}, nothing interface{}, f func(interface{}) bool, xs []interface{}) interface{} {
	for i := 0; i < len(xs); i++ {
		if f(xs[i]) {
			return just(int64(i))
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
