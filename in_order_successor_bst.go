package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var first = TreeNode{5, &second, &third}

var second = TreeNode{3, &firs, &seco}

var third = TreeNode{6, nil, nil}

var firs = TreeNode{2, &thi, nil}

var seco = TreeNode{4, nil, nil}

var thi = TreeNode{1, nil, nil}

func inOrderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var result *TreeNode
	var current = root

	for current != nil {
		if current.Val > p.Val {
			result = current
			current = current.Left
		} else if current.Val <= p.Val {
			current = current.Right
		}
	}
	return result
}

func main() {
	fmt.Print(inOrderSuccessor(&first, &third))
}
