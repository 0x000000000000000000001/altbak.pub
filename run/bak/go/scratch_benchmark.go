package main

import (
	"fmt"
	"time"
)

// --- Definitions ---

type Color bool
const (
	R Color = false
	B Color = true
)

type Node struct {
	color Color
	left  *Node
	key   int64
	right *Node
}

type Maybe struct {
	isJust bool
	value  int64
}

// 1. Current gopurs translation (Returns *Maybe, allocating on the heap)
//go:noinline
func lookup_allocated(key int64, node *Node) *Maybe {
	if node == nil {
		return &Maybe{isJust: false}
	}
	if key < node.key {
		return lookup_allocated(key, node.left)
	} else if key > node.key {
		return lookup_allocated(key, node.right)
	}
	return &Maybe{isJust: true, value: node.key}
}

// 2. Ideal TAST unboxed / Scrutinee Fusion (Returns value, ok, zero allocation)
//go:noinline
func lookup_unboxed(key int64, node *Node) (int64, bool) {
	if node == nil {
		return 0, false
	}
	if key < node.key {
		return lookup_unboxed(key, node.left)
	} else if key > node.key {
		return lookup_unboxed(key, node.right)
	}
	return node.key, true
}

func main() {
	// Build a dummy tree of depth 20 (about 1M nodes)
	var tree *Node = buildTree(1, 1000000)

	iterations := 10000000
	targetKey := int64(999999) // Deep in the tree

	// 1. Test Allocated
	startAllocated := time.Now()
	for i := 0; i < iterations; i++ {
		res := lookup_allocated(targetKey, tree)
		if res.isJust && res.value == -1 {
			// Prevent optimization
			fmt.Println("Impossible")
		}
	}
	durationAllocated := time.Since(startAllocated)

	// 2. Test Unboxed
	startUnboxed := time.Now()
	for i := 0; i < iterations; i++ {
		val, ok := lookup_unboxed(targetKey, tree)
		if ok && val == -1 {
			// Prevent optimization
			fmt.Println("Impossible")
		}
	}
	durationUnboxed := time.Since(startUnboxed)

	fmt.Printf("Iterations: %d\n\n", iterations)
	fmt.Printf("Current gopurs (Allocating Maybe) : %v\n", durationAllocated)
	fmt.Printf("Unboxed TAST (No Allocation)      : %v\n", durationUnboxed)
	fmt.Printf("\nSpeedup factor: %.2fx\n", float64(durationAllocated)/float64(durationUnboxed))
}

func buildTree(min, max int64) *Node {
	if min > max {
		return nil
	}
	mid := (min + max) / 2
	return &Node{
		color: B,
		left:  buildTree(min, mid-1),
		key:   mid,
		right: buildTree(mid+1, max),
	}
}
