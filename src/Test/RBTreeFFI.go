package Test_RBTreeFFI

type color bool

const (
	red   color = false
	black color = true
)

type tree struct {
	color color
	left  *tree
	value int
	right *tree
}

var pool []tree
var poolIdx int

func alloc(c color, l *tree, v int, r *tree) *tree {
	pool[poolIdx] = tree{color: c, left: l, value: v, right: r}
	p := &pool[poolIdx]
	poolIdx++
	return p
}

func balance(c color, a *tree, x int, b *tree) *tree {
	if c == black {
		if a != nil && a.color == red {
			if a.left != nil && a.left.color == red {
				return alloc(red, alloc(black, a.left.left, a.left.value, a.left.right), a.value, alloc(black, a.right, x, b))
			}
			if a.right != nil && a.right.color == red {
				return alloc(red, alloc(black, a.left, a.value, a.right.left), a.right.value, alloc(black, a.right.right, x, b))
			}
		}
		if b != nil && b.color == red {
			if b.left != nil && b.left.color == red {
				return alloc(red, alloc(black, a, x, b.left.left), b.left.value, alloc(black, b.left.right, b.value, b.right))
			}
			if b.right != nil && b.right.color == red {
				return alloc(red, alloc(black, a, x, b.left), b.value, alloc(black, b.right.left, b.right.value, b.right.right))
			}
		}
	}
	return alloc(c, a, x, b)
}

func ins(x int, t *tree) *tree {
	if t == nil {
		return alloc(red, nil, x, nil)
	}
	if x < t.value {
		return balance(t.color, ins(x, t.left), t.value, t.right)
	} else if x > t.value {
		return balance(t.color, t.left, t.value, ins(x, t.right))
	}
	return t
}

func insert(x int, t *tree) *tree {
	res := ins(x, t)
	return alloc(black, res.left, res.value, res.right)
}

func buildTree(n int, acc *tree) *tree {
	pool = make([]tree, 10000000) 
	poolIdx = 0

	for i := n; i > 0; i-- {
		acc = insert(i, acc)
	}
	return acc
}

func depth(t *tree) int {
	if t == nil {
		return 0
	}
	ld := depth(t.left)
	rd := depth(t.right)
	if ld > rd {
		return 1 + ld
	}
	return 1 + rd
}

func RunRBTreeFFI(limit float64) float64 {
	t := buildTree(int(limit), nil)
	return float64(depth(t))
}
