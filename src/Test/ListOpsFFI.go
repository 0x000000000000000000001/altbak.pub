package Test_ListOpsFFI

type ListOpsList[A any] interface {
	isList()
}
type ListOpsNil[A any] struct{}

func (ListOpsNil[A]) isList() {}

type ListOpsCons[A any] struct {
	value0 A
	value1 ListOpsList[A]
}

func (ListOpsCons[A]) isList() {}

func rangeListOps(start int, end int) ListOpsList[int] {
	var goFunc func(int, ListOpsList[int]) ListOpsList[int]
	goFunc = func(curr int, acc ListOpsList[int]) ListOpsList[int] {
		if curr < start {
			return acc
		}
		return goFunc(curr-1, ListOpsCons[int]{value0: curr, value1: acc})
	}
	return goFunc(end, ListOpsNil[int]{})
}

func filterEvens(lst ListOpsList[int]) ListOpsList[int] {
	var goFunc func(ListOpsList[int], ListOpsList[int]) ListOpsList[int]
	goFunc = func(list ListOpsList[int], acc ListOpsList[int]) ListOpsList[int] {
		switch l := list.(type) {
		case ListOpsNil[int]:
			return acc
		case ListOpsCons[int]:
			x := l.value0
			xs := l.value1
			if x%2 == 0 {
				return goFunc(xs, ListOpsCons[int]{value0: x, value1: acc})
			} else {
				return goFunc(xs, acc)
			}
		}
		return ListOpsNil[int]{}
	}
	return goFunc(lst, ListOpsNil[int]{})
}

func foldl[A, B any](f func(B) func(A) B, acc B, lst ListOpsList[A]) B {
	var goFunc func(ListOpsList[A], B) B
	goFunc = func(list ListOpsList[A], a B) B {
		switch l := list.(type) {
		case ListOpsNil[A]:
			return a
		case ListOpsCons[A]:
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
