package main

import (
	"fmt"
	"time"
)

type Node struct {
	size  int
	color int
	key   int
	value int
	left  *Node
	right *Node
}

func insert(node *Node, key, value int) *Node {
	if node == nil {
		return &Node{1, 1, key, value, nil, nil}
	}
	if key < node.key {
		left := insert(node.left, key, value)
		return balance(node.key, node.value, left, node.right)
	} else if key > node.key {
		right := insert(node.right, key, value)
		return balance(node.key, node.value, node.left, right)
	} else {
		return &Node{node.size, node.color, key, value, node.left, node.right}
	}
}

func balance(key, value int, left, right *Node) *Node {
	// Simplified just to simulate allocation and shallow copying
	return &Node{1, 1, key, value, left, right}
}

func main() {
	start := time.Now()
	var root *Node
	for i := 0; i < 100000; i++ {
		root = insert(root, i, i)
	}
	fmt.Printf("Go native tree insertions: %d us\n", time.Since(start).Microseconds())
}
