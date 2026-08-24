package Test_RBTreeFFICheatcode

type color_cheatcode bool

const (
	red_cheatcode   color_cheatcode = false
	black_cheatcode color_cheatcode = true
)

type tree_cheatcode struct {
	color_cheatcode color_cheatcode
	left  *tree_cheatcode
	value int
	right *tree_cheatcode
}

var pool []tree_cheatcode
var poolIdx int

func alloc_cheatcode(c color_cheatcode, l *tree_cheatcode, v int, r *tree_cheatcode) *tree_cheatcode {
	pool[poolIdx] = tree_cheatcode{color_cheatcode: c, left: l, value: v, right: r}
	p := &pool[poolIdx]
	poolIdx++
	return p
}

func balance_cheatcode(c color_cheatcode, a *tree_cheatcode, x int, b *tree_cheatcode) *tree_cheatcode {
	if c == black_cheatcode {
		if a != nil && a.color_cheatcode == red_cheatcode {
			if a.left != nil && a.left.color_cheatcode == red_cheatcode {
				return alloc_cheatcode(red_cheatcode, alloc_cheatcode(black_cheatcode, a.left.left, a.left.value, a.left.right), a.value, alloc_cheatcode(black_cheatcode, a.right, x, b))
			}
			if a.right != nil && a.right.color_cheatcode == red_cheatcode {
				return alloc_cheatcode(red_cheatcode, alloc_cheatcode(black_cheatcode, a.left, a.value, a.right.left), a.right.value, alloc_cheatcode(black_cheatcode, a.right.right, x, b))
			}
		}
		if b != nil && b.color_cheatcode == red_cheatcode {
			if b.left != nil && b.left.color_cheatcode == red_cheatcode {
				return alloc_cheatcode(red_cheatcode, alloc_cheatcode(black_cheatcode, a, x, b.left.left), b.left.value, alloc_cheatcode(black_cheatcode, b.left.right, b.value, b.right))
			}
			if b.right != nil && b.right.color_cheatcode == red_cheatcode {
				return alloc_cheatcode(red_cheatcode, alloc_cheatcode(black_cheatcode, a, x, b.left), b.value, alloc_cheatcode(black_cheatcode, b.right.left, b.right.value, b.right.right))
			}
		}
	}
	return alloc_cheatcode(c, a, x, b)
}

func ins_cheatcode(x int, t *tree_cheatcode) *tree_cheatcode {
	if t == nil {
		return alloc_cheatcode(red_cheatcode, nil, x, nil)
	}
	if x < t.value {
		return balance_cheatcode(t.color_cheatcode, ins_cheatcode(x, t.left), t.value, t.right)
	} else if x > t.value {
		return balance_cheatcode(t.color_cheatcode, t.left, t.value, ins_cheatcode(x, t.right))
	}
	return t
}

func insert_cheatcode(x int, t *tree_cheatcode) *tree_cheatcode {
	res := ins_cheatcode(x, t)
	return alloc_cheatcode(black_cheatcode, res.left, res.value, res.right)
}

func buildTree_cheatcode(n int, acc *tree_cheatcode) *tree_cheatcode {
	pool = make([]tree_cheatcode, 10000000) 
	poolIdx = 0

	for i := n; i > 0; i-- {
		acc = insert_cheatcode(i, acc)
	}
	return acc
}

func depth_cheatcode(t *tree_cheatcode) int {
	if t == nil {
		return 0
	}
	ld := depth_cheatcode(t.left)
	rd := depth_cheatcode(t.right)
	if ld > rd {
		return 1 + ld
	}
	return 1 + rd
}

func RunRBTreeFFICheatcode(limit int) int {
	t := buildTree_cheatcode(int(limit), nil)
	return (depth_cheatcode(t))
}
