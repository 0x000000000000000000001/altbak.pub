package main

import (
	"fmt"
	"math"
	"sort"
	"testing"

	ps "gopurs/output/purescript"
)

type tree = ps.Constructor_Test_RBTree_T

func black() uint32 { return uint32(ps.Get_Test_RBTree_B().IntVal) }
func red() uint32   { return uint32(ps.Get_Test_RBTree_R().IntVal) }

func render(node *tree) string {
	if node == nil {
		return "E"
	}
	return fmt.Sprintf("%d:%d(%s)(%s)", node.V0, node.V2, render(node.V1), render(node.V3))
}

func validate(t *testing.T, root *tree, expectedKeys []int64) {
	t.Helper()
	if root != nil && root.V0 != black() {
		t.Fatal("root is not black")
	}
	seen := map[*tree]bool{}
	keys := []int64{}
	var walk func(*tree, int64, int64) int
	walk = func(node *tree, lower, upper int64) int {
		if node == nil {
			return 1
		}
		if seen[node] {
			t.Fatal("cycle or node repeated within one tree")
		}
		seen[node] = true
		if node.V2 <= lower || node.V2 >= upper {
			t.Fatalf("key %d is outside (%d, %d)", node.V2, lower, upper)
		}
		if node.V0 != red() && node.V0 != black() {
			t.Fatalf("unknown color %d", node.V0)
		}
		if node.V0 == red() && ((node.V1 != nil && node.V1.V0 == red()) || (node.V3 != nil && node.V3.V0 == red())) {
			t.Fatal("red parent has red child")
		}
		leftHeight := walk(node.V1, lower, node.V2)
		keys = append(keys, node.V2)
		rightHeight := walk(node.V3, node.V2, upper)
		if leftHeight != rightHeight {
			t.Fatalf("different black heights under key %d", node.V2)
		}
		if node.V0 == black() {
			leftHeight++
		}
		return leftHeight
	}
	walk(root, math.MinInt64, math.MaxInt64)
	if fmt.Sprint(keys) != fmt.Sprint(expectedKeys) {
		t.Fatalf("keys %v; expected %v", keys, expectedKeys)
	}
}

func TestFourRotations(t *testing.T) {
	for _, order := range [][]int64{{3, 2, 1}, {3, 1, 2}, {1, 3, 2}, {1, 2, 3}} {
		t.Run(fmt.Sprint(order), func(t *testing.T) {
			var root *tree
			for _, key := range order {
				root = ps.Call_Test_RBTree_insert(key, root)
			}
			validate(t, root, []int64{1, 2, 3})
			if root.V2 != 2 || root.V1.V0 != black() || root.V3.V0 != black() {
				t.Fatal("unexpected three-key rotation result")
			}
		})
	}
}

func TestPersistentHistories(t *testing.T) {
	orders := [][]int64{{}, {1}, {3, 1, 2, 1, 3, 2}}
	ascending, descending, mixed := []int64{}, []int64{}, []int64{}
	for i := int64(1); i <= 127; i++ {
		ascending = append(ascending, i)
		descending = append(descending, 128-i)
		mixed = append(mixed, (i*37)%127)
	}
	orders = append(orders, ascending, descending, mixed)
	for index, order := range orders {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			var root *tree
			roots := []*tree{nil}
			images := []string{"E"}
			keySet := map[int64]bool{}
			for _, key := range order {
				root = ps.Call_Test_RBTree_insert(key, root)
				keySet[key] = true
				keys := []int64{}
				for item := range keySet {
					keys = append(keys, item)
				}
				sort.Slice(keys, func(a, b int) bool { return keys[a] < keys[b] })
				validate(t, root, keys)
				for i, old := range roots {
					if render(old) != images[i] {
						t.Fatalf("snapshot %d changed after inserting %d", i, key)
					}
				}
				roots = append(roots, root)
				images = append(images, render(root))
			}
		})
	}
}

func TestRetainedSubtreeAndSharedChild(t *testing.T) {
	root := ps.Call_Test_RBTree_buildTree(31, nil)
	child := root.V1
	beforeRoot, beforeChild := render(root), render(child)
	changed := ps.Call_Test_RBTree_insert(-1, root)
	if render(root) != beforeRoot || render(child) != beforeChild {
		t.Fatal("retained root or subtree changed")
	}
	keys := []int64{-1}
	for i := int64(1); i <= 31; i++ {
		keys = append(keys, i)
	}
	validate(t, changed, keys)

	// New parent, externally shared child: fresh parent is not deep uniqueness.
	leaf := &tree{Rc: 1, V0: black(), V2: 10}
	parent := &tree{Rc: 1, V0: black(), V1: leaf, V2: 50, V3: &tree{Rc: 1, V0: black(), V2: 90}}
	beforeLeaf := render(leaf)
	updated := ps.Call_Test_RBTree_insert(5, parent)
	validate(t, updated, []int64{5, 10, 50, 90})
	if render(leaf) != beforeLeaf {
		t.Fatal("shared child under fresh parent changed")
	}
}

func TestNominalBuild(t *testing.T) {
	root := buildFresh(100000)
	var count func(*tree) int
	count = func(node *tree) int {
		if node == nil {
			return 0
		}
		return 1 + count(node.V1) + count(node.V3)
	}
	nodes := count(root)
	depth := ps.Call_Test_RBTree_depth(root)
	if nodes != 100000 || depth != 22 {
		t.Fatalf("nodes=%d depth=%d; expected 100000 and 22", nodes, depth)
	}
	t.Logf("reachable nodes=%d depth=%d", nodes, depth)
}

func TestOwnedFreshSequences(t *testing.T) {
	orders := [][]int64{{}, {3, 2, 1}, {3, 1, 2}, {1, 3, 2}, {1, 2, 3}, {3, 1, 2, 3, 1, 2}}
	ascending, descending, mixed := []int64{}, []int64{}, []int64{}
	for i := int64(1); i <= 127; i++ {
		ascending = append(ascending, i)
		descending = append(descending, 128-i)
		mixed = append(mixed, (i*37)%127)
	}
	orders = append(orders, ascending, descending, mixed)
	for index, order := range orders {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			// Only this root owns the live tree. No prior node reference escapes
			// validate, and no snapshots/subtrees are retained across insertFresh.
			var root *tree
			keySet := map[int64]bool{}
			for _, key := range order {
				root = insertFresh(key, root)
				keySet[key] = true
				keys := []int64{}
				for item := range keySet {
					keys = append(keys, item)
				}
				sort.Slice(keys, func(a, b int) bool { return keys[a] < keys[b] })
				validate(t, root, keys)
			}
		})
	}
}

func TestFreshBuildMatchesPersistent(t *testing.T) {
	for _, count := range []int64{0, 1, 2, 3, 17, 127, 4096} {
		owned := buildFresh(count)
		persistent := ps.Call_Test_RBTree_buildTree(count, nil)
		if render(owned) != render(persistent) {
			t.Fatalf("different tree shape/colors for count %d", count)
		}
	}
}
