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
	refcount int
	color Color
	left  *Node
	value int
	right *Node
}

func balance(reuse *Node, color Color, a *Node, x int, b *Node) *Node {
	if color == B {
		if a != nil && a.color == R {
			if a.left != nil && a.left.color == R {
				if a.refcount == 1 && a.left.refcount == 1 {
					l := a.left
					l.color = B
					
					r := a
					r.color = B
					r.left = a.right
					r.value = x
					r.right = b
					
					if reuse != nil {
						reuse.color = R
						reuse.left = l
						reuse.value = a.value
						reuse.right = r
						return reuse
					}
					return &Node{1, R, l, a.value, r}
				}
			}
			if a.right != nil && a.right.color == R {
				if a.refcount == 1 && a.right.refcount == 1 {
					l := a
					l.color = B
					l.right = a.right.left
					
					r := a.right
					r.color = B
					r.left = a.right.right
					r.value = x
					r.right = b
					
					if reuse != nil {
						reuse.color = R
						reuse.left = l
						reuse.value = a.right.value
						reuse.right = r
						return reuse
					}
					return &Node{1, R, l, a.right.value, r}
				}
			}
		}
		if b != nil && b.color == R {
			if b.left != nil && b.left.color == R {
				if b.refcount == 1 && b.left.refcount == 1 {
					l := b.left
					l.color = B
					l.left = a
					l.value = x
					l.right = b.left.left
					
					r := b
					r.color = B
					r.left = b.left.right
					
					if reuse != nil {
						reuse.color = R
						reuse.left = l
						reuse.value = b.left.value
						reuse.right = r
						return reuse
					}
					return &Node{1, R, l, b.left.value, r}
				}
			}
			if b.right != nil && b.right.color == R {
				if b.refcount == 1 && b.right.refcount == 1 {
					l := b
					l.color = B
					l.left = a
					l.value = x
					l.right = b.left
					
					r := b.right
					r.color = B
					
					if reuse != nil {
						reuse.color = R
						reuse.left = l
						reuse.value = b.value
						reuse.right = r
						return reuse
					}
					return &Node{1, R, l, b.value, r}
				}
			}
		}
	}
	if reuse != nil {
		reuse.color = color
		reuse.left = a
		reuse.value = x
		reuse.right = b
		return reuse
	}
	return &Node{1, color, a, x, b}
}

func ins(x int, n *Node) *Node {
	if n == nil {
		return &Node{1, R, nil, x, nil}
	}
	
	var reuse *Node = nil
	if n.refcount == 1 {
		reuse = n
	}
	
	if x < n.value {
		return balance(reuse, n.color, ins(x, n.left), n.value, n.right)
	} else if x > n.value {
		return balance(reuse, n.color, n.left, n.value, ins(x, n.right))
	} else {
		return n
	}
}

func insert(x int, s *Node) *Node {
	n := ins(x, s)
	n.color = B
	return n
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
	fmt.Printf("Red-Black Tree FBIP: %d, Execution time: %d μs\n", res, dt)
}
