package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	p := 0
	if len(args) == 1 || args[1] == "1" {
		p = 1
	} else if args[1] == "2" {
		p = 2
	}

	f, err := os.Open(args[0])
	defer f.Close()

	buf := make([]byte, 5745)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f.Read(buf)
	b := blocks(buf)
	fmt.Println("******************")
	fmt.Println(parseSeeds(b[0]))
	fmt.Println("******************")
	test := Tree{}
	in := []int{2, 1, 5, 4, 6, 3}

	for _, v := range in {
		test.insert(0, v, 0)
	}

	fmt.Println(test.root)
	fmt.Println(test.root.left, test.root.right)
	fmt.Println(test.root.left.left, test.root.left.right, test.root.right.left, test.root.right.right)

	switch p {
	case 1:
		fmt.Println("Part 1")
	case 2:
		fmt.Println("Part 2")
	}
}

func blocks(buf []byte) (chunks [][]byte) {
	mid := bytes.Split(buf, []byte("\n\n"))
	for i := 0; i < len(mid)-1; i++ {
		if len(mid[i]) > 0 {
			chunks = append(chunks, mid[i])
		}
	}
	return chunks
}

func parseSeeds(block []byte) (seeds []int, err error) {
	in := make([]int, 0)
	for i, char := range block {
		if char == ' ' {
			in = append(in, i)
		}
	}

	for i := 0; i < len(in); i++ {
		if i == len(in)-1 {
			n, err := strconv.ParseInt(string(block[in[i]+1:]), 10, 64)
			if err != nil {
				return seeds, err
			}
			seeds = append(seeds, int(n))
		} else {
			n, err := strconv.ParseInt(string(block[in[i]+1:in[i+1]]), 10, 64)
			if err != nil {
				return seeds, err
			}
			seeds = append(seeds, int(n))
		}
	}
	return seeds, nil
}

type Node struct {
	lower int
	//, upper, dest,
	height int
	left   *Node
	right  *Node
}

type Tree struct {
	root *Node
}

func newNode(dest, source, r int) (n *Node) {
	n = &Node{
		lower: source,
		// upper: source + r -1,
		//	dest:   dest,
		height: 0,
		left:   nil,
		right:  nil,
	}
	return
}

func (t *Tree) insert(dest, source, r int) {
	visit := make([]*Node, 0)
	empty, s, _ := emptyNode(source, t.root, &visit)
	fmt.Println("inserting: ", source, empty, s)
	switch {
	case empty == nil:
		t.root = newNode(dest, source, r)
	case s == 0:
		empty.left = newNode(dest, source, r)
	case s == 1:
		empty.right = newNode(dest, source, r)
	}


	// Go back up the path to adjust height and rebalance
	for i := len(visit) - 1; i >= 0; i-- {
		fmt.Println("GOING back: ")
		// heights
		updateHeight(visit[i])
		// balance
		// Left side heavy
		if getHeight(visit[i].left)-getHeight(visit[i].right) > 1 {
		}
		// Right side heavy
		if getHeight(visit[i].left)-getHeight(visit[i].right) < -1 {
		}
		//}
	}
}

func updateHeight(n *Node) {
	if n.left == nil {
		n.height = n.right.height + 1
	} else if n.right == nil {
		n.height = n.left.height + 1
	} else {
		n.height = max(n.left.height, n.right.height) + 1
	}
}

func emptyNode(source int, c *Node, visit *[]*Node) (*Node, int, error) {
	if c == nil {
		return c,0, nil
	}
	for true {
		// Insert
		if source >= c.lower && c.right == nil {
			*visit = append(*visit, c)
			return c, 1, nil
		}
		if source < c.lower && c.left == nil {
			*visit = append(*visit, c)
			return c, 0, nil
		}
		// Move to the next node
		if source >= c.lower && c.right != nil {
			*visit = append(*visit, c)
			c = c.right
		}
		if source < c.lower && c.left != nil {
			*visit = append(*visit, c)
			c = c.left
		}
	}
	return nil, 0,errors.New("No location found")
}

func getHeight(n *Node) (h int) {
	if n != nil {
		h = n.height
	}
	return
}

func max(a, b int) (m int) {
	if a >= b {
		m = a
	} else {
		m = b
	}
	return
}

func (t Tree) value(l int) (f int) {
	return f
}

func (t Tree) rebalance() {}
