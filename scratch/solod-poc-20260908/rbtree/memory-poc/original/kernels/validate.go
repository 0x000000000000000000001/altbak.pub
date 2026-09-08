package kernels

// Validation describes the completed tree. BlackHeight includes the nil leaf.
type Validation struct {
	Count       int64
	Sum         int64
	Depth       int64
	BlackHeight int64
}

// Validate checks the benchmark tree containing exactly the keys 1 through n.
// It is intended for untimed verification of both Go and Solod builds.
func Validate(root *Constructor_Test_RBTree_T, n int64) Validation {
	// This range also keeps the expected sum and exclusive upper bound safe.
	if n < 0 || n > 1000000000 {
		panic("RBTree validation: unsupported size")
	}
	if root != nil {
		if root.V0 != 1583507464 && root.V0 != 3668501016 {
			panic("RBTree validation: unknown root color")
		}
		if root.V0 != 1583507464 {
			panic("RBTree validation: root is not black")
		}
	}

	// A valid red-black tree has depth at most 2*log2(n+1).
	// The extra level makes this integer bound conservative and caps recursion
	// even if a corrupted tree contains a cycle or a long invalid chain.
	var levels int64 = 0
	for m := n + 1; m > 1; m = m / 2 {
		levels = levels + 1
	}
	result := ValidateNode(root, 0, n+1, 1, 2*levels+1)
	if result.Count != n {
		panic("RBTree validation: wrong node count")
	}
	if result.Sum != n*(n+1)/2 {
		panic("RBTree validation: wrong key sum")
	}
	return result
}

func ValidateNode(root *Constructor_Test_RBTree_T, lower int64, upper int64, level int64, maxDepth int64) Validation {
	if root == nil {
		return Validation{Count: 0, Sum: 0, Depth: 0, BlackHeight: 1}
	}
	if level > maxDepth {
		panic("RBTree validation: excessive depth or cycle")
	}
	if root.V0 != 1583507464 && root.V0 != 3668501016 {
		panic("RBTree validation: unknown node color")
	}
	if root.V2 <= lower || root.V2 >= upper {
		panic("RBTree validation: BST ordering violation")
	}
	if root.V0 == 3668501016 {
		if root.V1 != nil && root.V1.V0 == 3668501016 {
			panic("RBTree validation: red left child of red node")
		}
		if root.V3 != nil && root.V3.V0 == 3668501016 {
			panic("RBTree validation: red right child of red node")
		}
	}

	left := ValidateNode(root.V1, lower, root.V2, level+1, maxDepth)
	right := ValidateNode(root.V3, root.V2, upper, level+1, maxDepth)
	if left.BlackHeight != right.BlackHeight {
		panic("RBTree validation: unequal black heights")
	}
	depth := left.Depth
	if right.Depth > depth {
		depth = right.Depth
	}
	blackHeight := left.BlackHeight
	if root.V0 == 1583507464 {
		blackHeight = blackHeight + 1
	}
	return Validation{
		Count:       left.Count + right.Count + 1,
		Sum:         left.Sum + right.Sum + root.V2,
		Depth:       depth + 1,
		BlackHeight: blackHeight,
	}
}
