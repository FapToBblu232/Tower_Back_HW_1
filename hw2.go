package main

import "fmt"

type BST_Node struct {
	value        int
	less_node    *BST_Node
	greater_node *BST_Node
}

type BST struct {
	head_node *BST_Node
}

func (node *BST_Node) Find(value int) bool {
	if node == nil {
		return false
	}
	if node.value == value {
		return true
	} else if value < node.value {
		return node.less_node.Find(value)
	} else {
		return node.greater_node.Find(value)
	}
}

func (tree *BST) Find(value int) bool {
	return tree.head_node.Find(value)
}

func (node *BST_Node) Add(number int) *BST_Node {
	var newNode *BST_Node = &BST_Node{value: number}
	if node == nil {
		return newNode
	}
	if number >= node.value {
		node.greater_node = node.greater_node.Add(number)
	} else {
		node.less_node = node.less_node.Add(number)
	}
	return node
}

func (tree *BST) Add(value int) {
	tree.head_node = tree.head_node.Add(value)
}

func (node *BST_Node) Mini() *BST_Node {
	if node == nil {
		return node
	}
	var cur *BST_Node = node
	for cur.less_node != nil {
		cur = cur.less_node
	}
	return cur
}

func (node *BST_Node) Remove(value int) *BST_Node {
	if node == nil {
		return node
	}

	if value > node.value {
		node.greater_node = node.greater_node.Remove(value)
	} else if value < node.value {
		node.less_node = node.less_node.Remove(value)
	} else {
		if node.greater_node == nil {
			return node.less_node
		} else if node.less_node == nil {
			return node.greater_node
		}
		var right_mini = node.greater_node.Mini()
		node.value = right_mini.value
		node.greater_node = node.greater_node.Remove(right_mini.value)
	}
	return node
}

func (tree *BST) Remove(value int) {
	tree.head_node.Remove(value)
}

func (node *BST_Node) Print() {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.value)
	node.less_node.Print()
	node.greater_node.Print()
}

func (tree *BST) Print() {
	tree.head_node.Print()
	fmt.Println()
}

func main() {
	//fmt.Println("Goodbay, World!")
	var bst BST
	bst.Add(5)
	bst.Add(3)
	bst.Add(4)
	bst.Add(9)
	bst.Add(8)
	bst.Print()
	bst.Remove(3)
	bst.Print()
}
