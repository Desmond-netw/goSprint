package main

import "fmt"

// use struct for custome types
type TreeNode struct {
	data  string
	left  *TreeNode
	right *TreeNode
}

// function
func main() {
	// crate root and instance of TreeNode
	root := &TreeNode{data: "R"}
	nodeA := &TreeNode{data: "A"}
	nodeB := &TreeNode{data: "B"}
	nodeC := &TreeNode{data: "C"}
	nodeD := &TreeNode{data: "D"}
	nodeE := &TreeNode{data: "E"}
	nodeF := &TreeNode{data: "F"}
	nodeG := &TreeNode{data: "G"}

	//linke the nodes
	root.left = nodeA
	root.right = nodeB

	nodeA.left = nodeC
	nodeA.right = nodeD

	nodeB.left = nodeE
	nodeB.right = nodeF

	nodeE.left = nodeG

	//Print from linked
	fmt.Println("Get data room root.Right.Left.Data: ", root.right.left.data)

	/*
			R
			/\
		   /  \
		  A    B
		 / \  / \
		C   D E  F
			  /
			 G

	*/

}
