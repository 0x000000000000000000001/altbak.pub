package main

import (
	"fmt"
	"time"
)

// ==============================================================================
// HYPOTHESIS 1: Unboxed ADTs + Defunctionalization (Typed Pointers, No Closures)
// ==============================================================================

type Color int32
const (
	R Color = 0
	B Color = 1
)

type Tree struct {
	color Color
	left  *Tree
	key   int64
	right *Tree
}

func balance(c Color, a *Tree, x int64, b *Tree) *Tree {
	if c == B {
		// balance B (T R (T R a x b) y c) z d = T R (T B a x b) y (T B c z d)
		if a != nil && a.color == R {
			if a.left != nil && a.left.color == R {
				return &Tree{R, &Tree{B, a.left.left, a.left.key, a.left.right}, a.key, &Tree{B, a.right, x, b}}
			}
			if a.right != nil && a.right.color == R {
				return &Tree{R, &Tree{B, a.left, a.key, a.right.left}, a.right.key, &Tree{B, a.right.right, x, b}}
			}
		}
		// balance B a x (T R (T R b y c) z d) = T R (T B a x b) y (T B c z d)
		// balance B a x (T R b y (T R c z d)) = T R (T B a x b) y (T B c z d)
		if b != nil && b.color == R {
			if b.left != nil && b.left.color == R {
				return &Tree{R, &Tree{B, a, x, b.left.left}, b.left.key, &Tree{B, b.left.right, b.key, b.right}}
			}
			if b.right != nil && b.right.color == R {
				return &Tree{R, &Tree{B, a, x, b.left}, b.key, &Tree{B, b.right.left, b.right.key, b.right.right}}
			}
		}
	}
	return &Tree{c, a, x, b}
}

func ins(x int64, t *Tree) *Tree {
	if t == nil {
		return &Tree{R, nil, x, nil}
	}
	if x < t.key {
		return balance(t.color, ins(x, t.left), t.key, t.right)
	} else if x > t.key {
		return balance(t.color, t.left, t.key, ins(x, t.right))
	}
	return t
}

func makeBlack(t *Tree) *Tree {
	if t != nil {
		return &Tree{B, t.left, t.key, t.right}
	}
	return nil
}

func insert(x int64, t *Tree) *Tree {
	return makeBlack(ins(x, t))
}

func buildTree(n int64, acc *Tree) *Tree {
	for n > 0 {
		acc = insert(n, acc)
		n--
	}
	return acc
}

func depth(t *Tree) int {
	if t == nil {
		return 0
	}
	d1 := depth(t.left)
	d2 := depth(t.right)
	if d1 > d2 {
		return 1 + d1
	}
	return 1 + d2
}

// ==============================================================================
// HYPOTHESIS 2: Native Arrays (Array-backed Trees with Integer Indices)
// ==============================================================================

type NodeIdx int32

type ArrayNode struct {
	color Color
	left  NodeIdx
	key   int64
	right NodeIdx
}

var arena []ArrayNode
var arenaIdx NodeIdx = 1 // 0 is nil

func allocArrayNode(c Color, l NodeIdx, k int64, r NodeIdx) NodeIdx {
	idx := arenaIdx
	arenaIdx++
	arena[idx] = ArrayNode{c, l, k, r}
	return idx
}

func balanceArr(c Color, a NodeIdx, x int64, b NodeIdx) NodeIdx {
	if c == B {
		if a != 0 && arena[a].color == R {
			if arena[a].left != 0 && arena[arena[a].left].color == R {
				return allocArrayNode(R, allocArrayNode(B, arena[arena[a].left].left, arena[arena[a].left].key, arena[arena[a].left].right), arena[a].key, allocArrayNode(B, arena[a].right, x, b))
			}
			if arena[a].right != 0 && arena[arena[a].right].color == R {
				return allocArrayNode(R, allocArrayNode(B, arena[a].left, arena[a].key, arena[arena[a].right].left), arena[arena[a].right].key, allocArrayNode(B, arena[arena[a].right].right, x, b))
			}
		}
		if b != 0 && arena[b].color == R {
			if arena[b].left != 0 && arena[arena[b].left].color == R {
				return allocArrayNode(R, allocArrayNode(B, a, x, arena[arena[b].left].left), arena[arena[b].left].key, allocArrayNode(B, arena[arena[b].left].right, arena[b].key, arena[b].right))
			}
			if arena[b].right != 0 && arena[arena[b].right].color == R {
				return allocArrayNode(R, allocArrayNode(B, a, x, arena[b].left), arena[b].key, allocArrayNode(B, arena[arena[b].right].left, arena[arena[b].right].key, arena[arena[b].right].right))
			}
		}
	}
	return allocArrayNode(c, a, x, b)
}

func insArr(x int64, t NodeIdx) NodeIdx {
	if t == 0 {
		return allocArrayNode(R, 0, x, 0)
	}
	if x < arena[t].key {
		return balanceArr(arena[t].color, insArr(x, arena[t].left), arena[t].key, arena[t].right)
	} else if x > arena[t].key {
		return balanceArr(arena[t].color, arena[t].left, arena[t].key, insArr(x, arena[t].right))
	}
	return t
}

func makeBlackArr(t NodeIdx) NodeIdx {
	if t != 0 {
		return allocArrayNode(B, arena[t].left, arena[t].key, arena[t].right)
	}
	return 0
}

func insertArr(x int64, t NodeIdx) NodeIdx {
	return makeBlackArr(insArr(x, t))
}

func buildTreeArr(n int64, acc NodeIdx) NodeIdx {
	for n > 0 {
		acc = insertArr(n, acc)
		n--
	}
	return acc
}

func depthArr(t NodeIdx) int {
	if t == 0 {
		return 0
	}
	d1 := depthArr(arena[t].left)
	d2 := depthArr(arena[t].right)
	if d1 > d2 {
		return 1 + d1
	}
	return 1 + d2
}

func main() {
	fmt.Println("=== BENCHMARK HYPOTHESES ===")
	
	// 1. Unboxed ADTs
	t0 := time.Now()
	res1 := buildTree(100000, nil)
	d1 := depth(res1)
	time1 := time.Since(t0)
	fmt.Printf("Hypothesis 1 (Unboxed Typed Pointers & Defunctionalized): %d\n", d1)
	fmt.Printf("Execution time: %.2f ms\n\n", float64(time1.Microseconds())/1000.0)
	
	// 2. Native Arrays
	arena = make([]ArrayNode, 10000000)
	t0 = time.Now()
	res2 := buildTreeArr(100000, 0)
	d2 := depthArr(res2)
	time2 := time.Since(t0)
	fmt.Printf("Hypothesis 2 (Native Arrays / Slices): %d\n", d2)
	fmt.Printf("Execution time: %.2f ms\n\n", float64(time2.Microseconds())/1000.0)
	
	fmt.Println("Note: Current gopurs implementation takes ~60 ms for this same code.")
}
