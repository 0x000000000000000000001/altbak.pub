package Data_Foldable

func FoldrArray(f func(interface{}) func(interface{}) interface{}, init interface{}, xs []interface{}) interface{} {
	acc := init
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(xs[i])(acc)
	}
	return acc
}

func FoldlArray(f func(interface{}) func(interface{}) interface{}, init interface{}, xs []interface{}) interface{} {
	acc := init
	for i := 0; i < len(xs); i++ {
		acc = f(acc)(xs[i])
	}
	return acc
}
