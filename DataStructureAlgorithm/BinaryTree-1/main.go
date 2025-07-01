package main

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

}
