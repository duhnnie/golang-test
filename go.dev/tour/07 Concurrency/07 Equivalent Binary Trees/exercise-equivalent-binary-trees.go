package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// fmt.Printf("always print %v - %v\n", t.Left, t.Right)

	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	res := make(chan bool)

	go Walk(t1, c1)
	go Walk(t2, c2)

	go func(c chan bool) {
		for i := 0; i < 10; i++ {
			x := <- c1
			y := <- c2

			fmt.Printf("%v == %v\n", x, y)

			if x != y {
				c <- false
			}
		}

		c <- true
	}(res)

	return <- res
}

func main() {
	treeOne := tree.New(1)
	treeTwo := tree.New(2)
	isSame := Same(treeOne, treeTwo)

	fmt.Println(isSame)
}