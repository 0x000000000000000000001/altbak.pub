package main

import (
	"fmt"
	"time"
)

type Color bool
const (
	R Color = false
	B Color = true
)

type Node struct {
	color Color
	left  *Node
	value int
	right *Node
}

func balance(color Color, a *Node, x int, b *Node) *Node {
	if color == B {
		if a != nil && a.color == R {
			if a.left != nil && a.left.color == R {
				return &Node{R, &Node{B, a.left.left, a.left.value, a.left.right}, a.value, &Node{B, a.right, x, b}}
			}
			if a.right != nil && a.right.color == R {
				return &Node{R, &Node{B, a.left, a.value, a.right.left}, a.right.value, &Node{B, a.right.right, x, b}}
			}
		}
		if b != nil && b.color == R {
			if b.left != nil && b.left.color == R {
				return &Node{R, &Node{B, a, x, b.left.left}, b.left.value, &Node{B, b.left.right, b.value, b.right}}
			}
			if b.right != nil && b.right.color == R {
				return &Node{R, &Node{B, a, x, b.left}, b.value, &Node{B, b.right.left, b.right.value, b.right.right}}
			}
		}
	}
	return &Node{color, a, x, b}
}

func ins(x int, n *Node) *Node {
	if n == nil {
		return &Node{R, nil, x, nil}
	}
	if x < n.value {
		return balance(n.color, ins(x, n.left), n.value, n.right)
	} else if x > n.value {
		return balance(n.color, n.left, n.value, ins(x, n.right))
	} else {
		return &Node{n.color, n.left, n.value, n.right}
	}
}

func insert(x int, s *Node) *Node {
	n := ins(x, s)
	return &Node{B, n.left, n.value, n.right}
}

func buildTree(n int, acc *Node) *Node {
	for n > 0 {
		acc = insert(n, acc)
		n--
	}
	return acc
}

func depth(n *Node) int {
	if n == nil {
		return 0
	}
	l := depth(n.left)
	r := depth(n.right)
	if l > r {
		return 1 + l
	}
	return 1 + r
}

func main() {
	start := time.Now()
	res := depth(buildTree(100000, nil))
	dt := time.Since(start).Microseconds()
	fmt.Printf("Red-Black Tree: %d, Execution time: %d μs\n", res, dt)
}
