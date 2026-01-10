package freetasks

import (
	"strconv"
	"strings"
)

type AVL struct {
	root *Node
}

type Node struct {
	left   *Node
	right  *Node
	value  int
	height int
}

func NewAVL() *AVL {
	return &AVL{}
}

func (t *AVL) Insert(value int) {
	t.root = t.insert(t.root, value)
}

func (t *AVL) Print() string {
	var nodes []*Node
	visited := make(map[*Node]bool)
	visited[t.root] = true
	nodes = append(nodes, t.root)
	b := strings.Builder{}

	for len(nodes) > 0 {
		n := nodes[0]
		if n == nil {
			b.WriteString("nil ")
		} else {
			b.WriteString(strconv.Itoa(n.value) + " ")
		}

		nodes = nodes[1:]

		if n == nil {
			continue
		}

		for _, l := range []*Node{n.left, n.right} {
			if _, ok := visited[l]; l != nil && ok {
				continue
			}

			nodes = append(nodes, l)
			visited[l] = true
		}
	}

	return b.String()
}

func (t *AVL) ASCII() string {
	if t.root == nil {
		return "<empty>\n"
	}

	var b strings.Builder
	writeNode(&b, t.root, "", true)
	return b.String()
}

func writeNode(b *strings.Builder, n *Node, prefix string, isTail bool) {
	if n == nil {
		return
	}

	if n.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		writeNode(b, n.right, newPrefix, false)
	}

	b.WriteString(prefix)
	if isTail {
		b.WriteString("└── ")
	} else {
		b.WriteString("┌── ")
	}
	b.WriteString(strconv.Itoa(n.value))
	b.WriteByte('\n')

	if n.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		writeNode(b, n.left, newPrefix, true)
	}
}

func (t *AVL) insert(n *Node, value int) *Node {
	if n == nil {
		return &Node{value: value, height: 1}
	}

	if value < n.value {
		n.left = t.insert(n.left, value)
	} else if value > n.value {
		n.right = t.insert(n.right, value)
	} else {
		return n
	}

	n.fixHeight()

	switch n.diff() {
	case 2:
		if n.right.diff() < 0 {
			n.right = t.rotateRight(n.right)
		}
		return t.rotateLeft(n)
	case -2:
		if n.left.diff() < 0 {
			n.left = t.rotateLeft(n.left)
		}
		return t.rotateRight(n)
	}

	return n
}

func (t *AVL) rotateLeft(n *Node) *Node {
	p := n.right
	n.right = p.left
	p.left = n

	n.fixHeight()
	p.fixHeight()

	return p
}

func (t *AVL) rotateRight(n *Node) *Node {
	q := n.left
	n.left = q.right
	q.right = n

	n.fixHeight()
	q.fixHeight()

	return q
}

func (n *Node) fixHeight() {
	if n.left.Height() > n.right.Height() {
		n.height = n.left.Height() + 1
	} else {
		n.height = n.right.Height() + 1
	}
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) diff() int {
	return n.right.Height() - n.left.Height()
}
