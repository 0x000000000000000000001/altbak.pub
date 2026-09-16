package Test_PrimesFFI


type List interface {
	isList()
}

type Nil struct{}

func (Nil) isList() {}

type Cons struct {
	value0 int
	value1 List
}

func (Cons) isList() {}

func rangeList(start int, end int) List {
	var goFunc func(int, List) List
	goFunc = func(curr int, acc List) List {
		if curr < start {
			return acc
		}
		return goFunc(curr-1, Cons{value0: curr, value1: acc})
	}
	return goFunc(end, Nil{})
}

func filter(p func(int) bool, lst List) List {
	var goFunc func(List, List) List
	goFunc = func(list List, acc List) List {
		switch l := list.(type) {
		case Nil:
			return reverse(acc)
		case Cons:
			x := l.value0
			xs := l.value1
			if p(x) {
				return goFunc(xs, Cons{value0: x, value1: acc})
			} else {
				return goFunc(xs, acc)
			}
		}
		return Nil{}
	}
	return goFunc(lst, Nil{})
}

func reverse(lst List) List {
	var goFunc func(List, List) List
	goFunc = func(list List, acc List) List {
		switch l := list.(type) {
		case Nil:
			return acc
		case Cons:
			return goFunc(l.value1, Cons{value0: l.value0, value1: acc})
		}
		return Nil{}
	}
	return goFunc(lst, Nil{})
}

func sieve(lst List) List {
	switch l := lst.(type) {
	case Nil:
		return Nil{}
	case Cons:
		p := l.value0
		xs := l.value1
		return Cons{
			value0: p,
			value1: sieve(filter(func(x int) bool {
				return x%p != 0
			}, xs)),
		}
	}
	return Nil{}
}

func sumList(lst List) int {
	var goFunc func(List, int) int
	goFunc = func(list List, acc int) int {
		switch l := list.(type) {
		case Nil:
			return acc
		case Cons:
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
