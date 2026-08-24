package Test_ListOpsFFI


type ListOpsList interface {
	isList()
}
type ListOpsNil struct{}
func (ListOpsNil) isList() {}
type ListOpsCons struct {
	value0 int
	value1 ListOpsList
}
func (ListOpsCons) isList() {}

func rangeListOps(start int, end int) ListOpsList {
	var goFunc func(int, ListOpsList) ListOpsList
	goFunc = func(curr int, acc ListOpsList) ListOpsList {
		if curr < start {
			return acc
		}
		return goFunc(curr-1, ListOpsCons{value0: curr, value1: acc})
	}
	return goFunc(end, ListOpsNil{})
}

func filterEvens(lst ListOpsList) ListOpsList {
	var goFunc func(ListOpsList, ListOpsList) ListOpsList
	goFunc = func(list ListOpsList, acc ListOpsList) ListOpsList {
		switch l := list.(type) {
		case ListOpsNil:
			return acc
		case ListOpsCons:
			x := l.value0
			xs := l.value1
			if x%2 == 0 {
				return goFunc(xs, ListOpsCons{value0: x, value1: acc})
			} else {
				return goFunc(xs, acc)
			}
		}
		return ListOpsNil{}
	}
	return goFunc(lst, ListOpsNil{})
}

func foldl(f func(int) func(int) int, acc int, lst ListOpsList) int {
	var goFunc func(ListOpsList, int) int
	goFunc = func(list ListOpsList, a int) int {
		switch l := list.(type) {
		case ListOpsNil:
			return a
		case ListOpsCons:
			return goFunc(l.value1, f(a)(l.value0))
		}
		return a
	}
	return goFunc(lst, acc)
}

func RunListOpsFFI(limit int) int {
	n := int(limit)
	res := foldl(func(acc int) func(int) int {
		return func(x int) int {
			return acc + x
		}
	}, 0, filterEvens(rangeListOps(1, n)))
	return (res)
}
