package main

import "fmt"


type node struct {
	value  int
	bet    int
	height int
	left   *node
	right  *node
}

func getHeight(n *node) int {
	if n != nil {
		return n.height
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insertNode(n *node, v int, b int) *node {
	if n == nil {
		return newNode(v, b)
	} else if v < n.value {
		n.left = insertNode(n.left, v, b)
	} else if v > n.value {
		n.right = insertNode(n.right, v, b)
	} else {
		return n
	}

	n.height = 1 + max(getHeight(n.left), getHeight(n.right))
	balance := getHeight(n.left) - getHeight(n.right)

	// right heavy
	if balance < -1 {
		if v > n.right.value {
			return rotateLeft(n)
		} else if v < n.right.value {
			n.right = rotateRight(n.right)
			return rotateLeft(n)
		}
	}

	if balance > 1 {
		if v < n.left.value {
			return rotateRight(n)
		} else if v > n.left.value {
			n.left = rotateLeft(n.left)
			return rotateRight(n)
		}
	}

	return n
}

func rotateRight(n *node) *node {
	t := n.left
	t1 := t.right
	n.left = t1
	t.right = n
	n.height = 1 + max(getHeight(n.left), getHeight(n.right))
	t.height = 1 + max(getHeight(n.left), getHeight(n.right))

	return t
}

func rotateLeft(n *node) *node {
	t := n.right
	t1 := t.left
	n.right = t1
	t.left = n
	n.height = 1 + max(getHeight(n.left), getHeight(n.right))
	t.height = 1 + max(getHeight(n.left), getHeight(n.right))

	return t
}

func newNode(v, b int) *node {
	return &node{
		value:  v,
		bet:    b,
		height: 1,
		left:   nil,
		right:  nil,
	}
}

func inorderTran(n *node, i *int) {
	if n.left != nil {
		inorderTran(n.left, i)
	}
	fmt.Println("Val:", n.value, "Ind: ", *i)
	*i += 1
	if n.right != nil {
		inorderTran(n.right, i)
	}
}
