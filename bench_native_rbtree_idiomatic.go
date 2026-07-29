package main

import (
	"fmt"
	"time"
)

type Color bool

const (
	RED   Color = false
	BLACK Color = true
)

type Node struct {
	color Color
	value int
	left  *Node
	right *Node
	parent *Node
}

type RBTree struct {
	root *Node
}

func (t *RBTree) Insert(value int) {
	node := &Node{color: RED, value: value}
	
	var y *Node = nil
	x := t.root

	for x != nil {
		y = x
		if node.value < x.value {
			x = x.left
		} else {
			x = x.right
		}
	}

	node.parent = y
	if y == nil {
		t.root = node
	} else if node.value < y.value {
		y.left = node
	} else {
		y.right = node
	}

	t.fixInsert(node)
}

func (t *RBTree) fixInsert(k *Node) {
	for k.parent != nil && k.parent.color == RED {
		if k.parent == k.parent.parent.left {
			u := k.parent.parent.right
			if u != nil && u.color == RED {
				k.parent.color = BLACK
				u.color = BLACK
				k.parent.parent.color = RED
				k = k.parent.parent
			} else {
				if k == k.parent.right {
					k = k.parent
					t.leftRotate(k)
				}
				k.parent.color = BLACK
				k.parent.parent.color = RED
				t.rightRotate(k.parent.parent)
			}
		} else {
			u := k.parent.parent.left
			if u != nil && u.color == RED {
				k.parent.color = BLACK
				u.color = BLACK
				k.parent.parent.color = RED
				k = k.parent.parent
			} else {
				if k == k.parent.left {
					k = k.parent
					t.rightRotate(k)
				}
				k.parent.color = BLACK
				k.parent.parent.color = RED
				t.leftRotate(k.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func (t *RBTree) leftRotate(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RBTree) rightRotate(y *Node) {
	x := y.left
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}
	x.right = y
	y.parent = x
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
	
	tree := &RBTree{}
	// Worst case: inserting sequentially
	for i := 100000; i > 0; i-- {
		tree.Insert(i)
	}

	res := depth(tree.root)
	dt := time.Since(start).Microseconds()
	fmt.Printf("Red-Black Tree: %d, Execution time: %d μs\n", res, dt)
}
