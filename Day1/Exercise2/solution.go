package main

import "fmt"

type node struct {
	 left *node
	 right *node
	 data uint8
}
func preOrder(root *node) {
	if root != nil {
		fmt.Printf("%c ", root.data)
		preOrder(root.left)
		preOrder(root.right)
	}
}
func postOrder(root *node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Printf("%c ", root.data)
}
func main () {
	a := "a+b-c"
	//fmt.Printf("%T",a[0])
	rootNode := node{nil,nil,a[1]}
	rootNode.left = &node{nil,nil, a[0]}
	rootNode.right = &node{nil, nil, a[3]}
	rootNode.right.left = &node{nil,nil, a[2]}
	rootNode.right.right = &node{nil,nil, a[4]}
	//fmt.Println(rootNode.data)
	fmt.Println("Preorder traversel is : ")
	preOrder(&rootNode)
	fmt.Println("\n Postorder Traversal is")
	postOrder(&rootNode)
}