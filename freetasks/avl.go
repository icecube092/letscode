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

func (t *AVL) Remove(value int) {
	t.root = t.remove(t.root, value)
}

func (t *AVL) String() string {
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

		for _, child := range []*Node{n.left, n.right} {
			if _, ok := visited[child]; child != nil && ok {
				continue
			}

			nodes = append(nodes, child)
			visited[child] = true
		}
	}

	return strings.TrimSpace(b.String())
}

func (t *AVL) ASCII() string {
	if t.root == nil {
		return "nil"
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

	n = t.balance(n)

	return n
}

func (t *AVL) balance(n *Node) *Node {
	n.updateHeight()

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

func (t *AVL) remove(n *Node, value int) *Node {
	if n == nil {
		return n
	}

	if value < n.value {
		n.left = t.remove(n.left, value)
	} else if value > n.value {
		n.right = t.remove(n.right, value)
	} else {
		left := n.left
		right := n.right

		n.left = nil
		n.right = nil
		n = nil

		if right == nil {
			return left
		}

		minRight := t.findMin(right)
		minRight.right = t.removeMin(right)
		minRight.left = left
		return t.balance(minRight)
	}

	return t.balance(n)
}

func (t *AVL) findMin(n *Node) *Node {
	if n.left != nil {
		return t.findMin(n.left)
	}

	return n
}

func (t *AVL) removeMin(n *Node) *Node {
	if n.left == nil {
		return n.right
	}
	n.left = t.removeMin(n.left)

	return t.balance(n)
}

func (t *AVL) rotateLeft(n *Node) *Node {
	next := n.right

	if next != nil {
		n.right = next.left
		next.left = n
	} else {
		next = n
	}

	n.updateHeight()
	next.updateHeight()

	return next
}

func (t *AVL) rotateRight(n *Node) *Node {
	next := n.left

	if next != nil {
		n.left = next.right
		next.right = n
	} else {
		next = n
	}

	n.updateHeight()
	next.updateHeight()

	return next
}

func (n *Node) updateHeight() {
	if n == nil {
		return
	}

	if n.left.getHeight() > n.right.getHeight() {
		n.height = n.left.getHeight() + 1
	} else {
		n.height = n.right.getHeight() + 1
	}
}

func (n *Node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) diff() int {
	return n.right.getHeight() - n.left.getHeight()
}
