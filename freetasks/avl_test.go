package freetasks

import (
	"fmt"
	"testing"
)

func TestAVL(t *testing.T) {
	tree := NewAVL()

	for i := range 8 {
		tree.Insert(i)
	}

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

	tree.Insert(0)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	for i := range 10 {
		tree.Insert(i)
	}

	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	tree.Remove(3)
	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())

	for i := range 3 {
		tree.Remove(i)
	}

	fmt.Println(tree.String())
	fmt.Println(tree.ASCII())
}
