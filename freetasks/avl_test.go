package freetasks

import (
	"fmt"
	"testing"
)

func TestAVL(t *testing.T) {
	tree := NewAVL()

	tree.Insert(1)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	tree.Insert(2)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	tree.Insert(3)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	tree.Insert(4)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	tree.Insert(5)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	tree.Insert(6)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	tree.Insert(7)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	tree.Remove(7)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	tree.Remove(3)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	tree.Remove(4)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	tree.Remove(0)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())
}
