package main

import "fmt"

type node struct {
	value  int
	bet    int
	height int
	left   *node
	right  *node
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

func inorderTran(n *node, i *int, sum *int) {
	if n == nil {
		fmt.Println("Tree is empty")
		return
	}
	if n.left != nil {
		inorderTran(n.left, i, sum)
	}
	*sum += n.bet * *i
	*i += 1
	if n.right != nil {
		inorderTran(n.right, i, sum)
	}
}

// HERE

func insertNode(n *node, value int, bet int) *node {
	if n == nil {
		n = newNode(value, bet)
		return n
	}
	search := true
	c := n
	v := make([]*node, 0)
	for search {
		switch {
		case value < c.value:
			v = append(v, c)
			if c.left == nil {
				c.left = newNode(value, bet)
				search = false
			} else {
				c = c.left
			}
		case value > c.value:
			v = append(v, c)
			if c.right == nil {
				c.right = newNode(value, bet)
				search = false
			} else {
				c = c.right
			}
		}
		for i := len(v) - 1; i >= 0; i-- {
			node := v[i]
			node.height = newHeight(node)
			balance := height(node.left) - height(node.right)
			switch {
			case balance > 1:
				// base case: if node.left.right is nil then the nodes are
				// lined up and only one rotation is necessary. Using the height
				// of the branches generalizes the base case
				if height(node.left.left) > height(node.left.right) {
					if i == 0 {
						// if your are rotating the root node set the return to
						// the new root after rotating
						n = rotateRight(node)
					} else if i > 0 {
						// if your are not at the root node go one level up and
						// set the previous left || right node to the new one
						if v[i-1].left == node {
							v[i-1].left = rotateRight(node)
						} else if v[i-1].right == node {
							v[i-1].right = rotateRight(node)
						}
					}
					// base case: if node.left.left is nil then the nodes are
					// stagared and two rotations are necessary. Using the height
					// of the branches generalizes the base case
				} else if height(node.left.left) < height(node.left.right) {
					node.left = rotateLeft(node.left)
					if i == 0 {
						n = rotateRight(node)
					} else if i > 0 {
						if v[i-1].left == node {
							v[i-1].left = rotateRight(node)
						} else if v[i-1].right == node {
							v[i-1].right = rotateRight(node)
						}
					}
				}

			case balance < -1:
				if height(node.right.right) > height(node.right.left) {
					if i == 0 {
						n = rotateLeft(node)
					} else if i > 0 {
						if v[i-1].left == node {
							v[i-1].left = rotateLeft(node)
						} else if v[i-1].right == node {
							v[i-1].right = rotateLeft(node)
						}
					}
				} else if height(node.right.right) < height(node.right.left) {
					node.right = rotateRight(node.right)
					if i == 0 {
						n = rotateLeft(node)
					} else if i > 0 {
						if v[i-1].left == node {
							v[i-1].left = rotateLeft(node)
						} else if v[i-1].right == node {
							v[i-1].right = rotateLeft(node)
						}
					}
				}
			}
		}
	}
	return n
}

func rotateRight(n *node) *node {
	t := n.left
	t1 := t.right
	t.right = n
	n.left = t1
	n.height = newHeight(n)
	t.height = newHeight(t)
	return t
}

func rotateLeft(n *node) *node {
	t := n.right
	t1 := n.right.left
	t.left = n
	n.right = t1
	n.height = newHeight(n)
	t.height = newHeight(t)
	return t
}

func newHeight(n *node) int {
	return max(height(n.left), height(n.right)) + 1
}

func height(n *node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
