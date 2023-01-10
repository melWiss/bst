package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	value      int
	parentNode *Node
	leftNode   *Node
	rightNode  *Node
}

type BinaryTree *Node

func addValueToTheTree(value int, tree BinaryTree) {
	if tree.value < 0 {
		tree.value = value
	} else if value <= tree.value {
		if tree.leftNode == nil {
			tree.leftNode = &Node{value: value, parentNode: tree}
		} else {
			addValueToTheTree(value, tree.leftNode)
		}
	} else {
		if tree.rightNode == nil {
			tree.rightNode = &Node{value: value, parentNode: tree}
		} else {
			addValueToTheTree(value, tree.rightNode)
		}
	}
}
func addNodeToTheTree(node *Node, tree BinaryTree) {
	if node == nil {
		return
	} else if tree.value < 0 {
		tree.value = node.value
	} else if node.value <= tree.value {
		if tree.leftNode == nil {
			tree.leftNode = node
		} else {
			addNodeToTheTree(node, tree.leftNode)
		}
	} else {
		if tree.rightNode == nil {
			tree.rightNode = node
		} else {
			addNodeToTheTree(node, tree.rightNode)
		}
	}
}

func searchForNodeByValue(value int, tree BinaryTree) BinaryTree {
	if value == tree.value {
		return tree
	} else if value < tree.value {
		if tree.leftNode != nil {
			return searchForNodeByValue(value, tree.leftNode)
		} else {
			return nil
		}
	} else {
		if tree.rightNode != nil {
			return searchForNodeByValue(value, tree.rightNode)
		} else {
			return nil
		}
	}
}
func goSearchForNodeByValue(value int, tree BinaryTree, channel chan BinaryTree) {
	if value == tree.value {
		channel <- tree
	} else if value < tree.value {
		if tree.leftNode != nil {
			go goSearchForNodeByValue(value, tree.leftNode, channel)
		}
	} else {
		if tree.rightNode != nil {
			go goSearchForNodeByValue(value, tree.rightNode, channel)
		}
	}
}

func removeValueFromTheTree(value int, tree BinaryTree) {
	node := searchForNodeByValue(value, tree)
	if node == nil {
		return
	}
	leftNode, rightNode, parentNode := node.leftNode, node.rightNode, node.parentNode
	if parentNode == nil {
		tree = rightNode
		addNodeToTheTree(leftNode, tree)
	} else {
		if parentNode.leftNode == node {
			parentNode.leftNode = nil
		} else {
			parentNode.rightNode = nil
		}
		addNodeToTheTree(leftNode, parentNode)
		addNodeToTheTree(rightNode, parentNode)
	}
}

func printTree(tree BinaryTree) {
	fmt.Println(tree.value)
	if tree.leftNode != nil {
		printTree(tree.leftNode)
	}
	if tree.rightNode != nil {
		printTree(tree.rightNode)
	}
}

func main() {
	var tree BinaryTree = &Node{value: -1}
	value := 0
	size := 20000000
	for size > 0 {
		value = rand.Intn(60000000)
		addValueToTheTree(value, tree)
		size--
	}
	// printTree(tree)
	valueChan := make(chan BinaryTree)
	value = 59999997
	println("------------searching for value 59999997 using goroutines ------------------")
	start := time.Now()
	go goSearchForNodeByValue(value, tree, valueChan)
	println(<-valueChan)
	println(time.Since(start).Nanoseconds(), " ns")
	println("------------searching for value 59999997 without using goroutines ------------------")
	start = time.Now()
	node := searchForNodeByValue(value, tree)
	println(node)
	println(time.Since(start).Nanoseconds(), " ns")
}
