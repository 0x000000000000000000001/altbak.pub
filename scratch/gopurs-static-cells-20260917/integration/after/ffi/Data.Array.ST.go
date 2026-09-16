package Data_Array_ST

import "sort"

func NewImpl(s interface{}) *[]interface{} {
	arr := make([]interface{}, 0)
	return &arr
}

func PeekImpl(just func(interface{}) interface{}, nothing interface{}, i int64, arr *[]interface{}) interface{} {
	if i >= 0 && i < int64(len(*arr)) {
		return just((*arr)[i])
	}
	return nothing
}

func PokeImpl(i int64, a interface{}, arr *[]interface{}) interface{} {
	if i >= 0 && i < int64(len(*arr)) {
		(*arr)[i] = a
		return true
	}
	return false
}

func LengthImpl(arr *[]interface{}) interface{} {
	return int64(len(*arr))
}

func PopImpl(just func(interface{}) interface{}, nothing interface{}, arr *[]interface{}) interface{} {
	if len(*arr) > 0 {
		last := (*arr)[len(*arr)-1]
		*arr = (*arr)[:len(*arr)-1]
		return just(last)
	}
	return nothing
}

func PushAllImpl(xs []interface{}, arr *[]interface{}) interface{} {
	*arr = append(*arr, xs...)
	return int64(len(*arr))
}

func PushImpl(x interface{}, arr *[]interface{}) interface{} {
	*arr = append(*arr, x)
	return int64(len(*arr))
}

func ShiftImpl(just func(interface{}) interface{}, nothing interface{}, arr *[]interface{}) interface{} {
	if len(*arr) > 0 {
		first := (*arr)[0]
		*arr = (*arr)[1:]
		return just(first)
	}
	return nothing
}

func UnshiftAllImpl(xs []interface{}, arr *[]interface{}) interface{} {
	*arr = append(xs, *arr...)
	return int64(len(*arr))
}

func UnshiftImpl(x interface{}, arr *[]interface{}) interface{} {
	*arr = append([]interface{}{x}, *arr...)
	return int64(len(*arr))
}

func SpliceImpl(start int64, count int64, xs []interface{}, arr *[]interface{}) interface{} {
	removed := make([]interface{}, count)
	copy(removed, (*arr)[start:start+count])
	
	newArr := make([]interface{}, 0, len(*arr) - int(count) + len(xs))
	newArr = append(newArr, (*arr)[:start]...)
	newArr = append(newArr, xs...)
	newArr = append(newArr, (*arr)[start+count:]...)
	*arr = newArr
	return removed
}

func UnsafeFreezeImpl(arr *[]interface{}) []interface{} {
	return *arr
}

func UnsafeThawImpl(xs []interface{}) *[]interface{} {
	return &xs
}

func FreezeImpl(arr *[]interface{}) []interface{} {
	return *arr
}

func ThawImpl(xs []interface{}) *[]interface{} {
	return &xs
}

func CloneImpl(arr *[]interface{}) *[]interface{} {
	res := make([]interface{}, len(*arr))
	copy(res, *arr)
	return &res
}

func SortByImpl(f func(interface{}, interface{}) interface{}, toInt func(interface{}) int64, arr *[]interface{}) interface{} {
	sort.SliceStable(*arr, func(i, j int) bool {
		ord := f((*arr)[i], (*arr)[j])
		return toInt(ord) < 0
	})
	return arr
}

func ToAssocArrayImpl(arr *[]interface{}) interface{} {
	res := make([]interface{}, len(*arr))
	for i, v := range *arr {
		res[i] = map[string]interface{}{
			"value": v,
			"index": int64(i),
		}
	}
	return res
}
