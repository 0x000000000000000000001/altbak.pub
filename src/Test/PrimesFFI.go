package Test_PrimesFFI

type List[A any] interface {
	isList()
}

type Nil[A any] struct{}

func (Nil[A]) isList() {}

type Cons[A any] struct {
	value0 A
	value1 List[A]
}

func (Cons[A]) isList() {}

func rangeList(start int, end int) List[int] {
	var goFunc func(int, List[int]) List[int]
	goFunc = func(curr int, acc List[int]) List[int] {
		if curr < start {
			return acc
		}
		return goFunc(curr-1, Cons[int]{value0: curr, value1: acc})
	}
	return goFunc(end, Nil[int]{})
}

func filter[A any](p func(A) bool, lst List[A]) List[A] {
	var goFunc func(List[A], List[A]) List[A]
	goFunc = func(list List[A], acc List[A]) List[A] {
		switch l := list.(type) {
		case Nil[A]:
			return reverse[A](acc)
		case Cons[A]:
			x := l.value0
			xs := l.value1
			if p(x) {
				return goFunc(xs, Cons[A]{value0: x, value1: acc})
			} else {
				return goFunc(xs, acc)
			}
		}
		return Nil[A]{}
	}
	return goFunc(lst, Nil[A]{})
}

func reverse[A any](lst List[A]) List[A] {
	var goFunc func(List[A], List[A]) List[A]
	goFunc = func(list List[A], acc List[A]) List[A] {
		switch l := list.(type) {
		case Nil[A]:
			return acc
		case Cons[A]:
			return goFunc(l.value1, Cons[A]{value0: l.value0, value1: acc})
		}
		return Nil[A]{}
	}
	return goFunc(lst, Nil[A]{})
}

func sieve(lst List[int]) List[int] {
	switch l := lst.(type) {
	case Nil[int]:
		return Nil[int]{}
	case Cons[int]:
		p := l.value0
		xs := l.value1
		return Cons[int]{
			value0: p,
			value1: sieve(filter(func(x int) bool {
				return x%p != 0
			}, xs)),
		}
	}
	return Nil[int]{}
}

func sumList(lst List[int]) int {
	var goFunc func(List[int], int) int
	goFunc = func(list List[int], acc int) int {
		switch l := list.(type) {
		case Nil[int]:
			return acc
		case Cons[int]:
			return goFunc(l.value1, acc+l.value0)
		}
		return acc
	}
	return goFunc(lst, 0)
}

func RunPrimesFFI(limit int) int {
	dummy := limit
	return (sumList(sieve(rangeList(2, dummy))))
}
