package main

import "fmt"

// get struct of the Node

type Node struct {
	data  int
	left  *Node
	right *Node
}

// search for node in Root Left Right
func PrintPreOrder(node *Node) {
	// edge case such as node being nil
	if node == nil {
		return
	}

	// 1. deal with current node and print it
	fmt.Printf("%d ", node.data)

	// Recure on the left side of tree
	PrintPreOrder(node.left)

	//3. Recur the right side of tree
	PrintPreOrder(node.right)
}
func main() {
	// Construct the binary tree
	//     1
	//    / \
	//   2   3
	//  / \   \
	// 4   5   6
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.right = &Node{data: 6}

	fmt.Println("Preorder Traversal:")
	// Call the function to print the preorder traversal of the created tree.
	PrintPreOrder(root)
	fmt.Println() // Print a final newline for clean output.
}
