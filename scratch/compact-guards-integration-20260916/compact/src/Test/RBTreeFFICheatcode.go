package Test_RBTreeFFICheatcode

type color_cheatcode bool

const (
	red_cheatcode   color_cheatcode = false
	black_cheatcode color_cheatcode = true
)

type tree_cheatcode struct {
	color_cheatcode color_cheatcode
	left            *tree_cheatcode
	value           int
	right           *tree_cheatcode
}

func isRed_cheatcode(t *tree_cheatcode) bool {
	return t != nil && t.color_cheatcode == red_cheatcode
}

// Apply the same Okasaki rotations, updating nodes owned by this tree.
func balance_cheatcode(t *tree_cheatcode) *tree_cheatcode {
	if t.color_cheatcode != black_cheatcode {
		return t
	}
	left := t.left
	right := t.right
	if isRed_cheatcode(left) {
		if isRed_cheatcode(left.left) {
			t.left = left.right
			left.right = t
			left.left.color_cheatcode = black_cheatcode
			left.color_cheatcode = red_cheatcode
			return left
		}
		if isRed_cheatcode(left.right) {
			middle := left.right
			left.right = middle.left
			t.left = middle.right
			middle.left = left
			middle.right = t
			left.color_cheatcode = black_cheatcode
			middle.color_cheatcode = red_cheatcode
			return middle
		}
	}
	if isRed_cheatcode(right) {
		if isRed_cheatcode(right.left) {
			middle := right.left
			t.right = middle.left
			right.left = middle.right
			middle.left = t
			middle.right = right
			right.color_cheatcode = black_cheatcode
			middle.color_cheatcode = red_cheatcode
			return middle
		}
		if isRed_cheatcode(right.right) {
			t.right = right.left
			right.left = t
			right.right.color_cheatcode = black_cheatcode
			right.color_cheatcode = red_cheatcode
			return right
		}
	}
	return t
}

func ins_cheatcode(x int, t *tree_cheatcode) *tree_cheatcode {
	if t == nil {
		return &tree_cheatcode{color_cheatcode: red_cheatcode, value: x}
	}
	if x < t.value {
		t.left = ins_cheatcode(x, t.left)
	} else if x > t.value {
		t.right = ins_cheatcode(x, t.right)
	} else {
		return t
	}
	return balance_cheatcode(t)
}

func insert_cheatcode(x int, t *tree_cheatcode) *tree_cheatcode {
	res := ins_cheatcode(x, t)
	res.color_cheatcode = black_cheatcode
	return res
}

func buildTree_cheatcode(n int, acc *tree_cheatcode) *tree_cheatcode {
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
	t := buildTree_cheatcode(limit, nil)
	return depth_cheatcode(t)
}
