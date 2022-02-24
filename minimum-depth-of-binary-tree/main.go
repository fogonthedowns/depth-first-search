package main

import (
	"container/list"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	node  *TreeNode
	depth int
}

// minDepth with stack
func minDepth(root *TreeNode) int {
	l := list.New()
	if root == nil {
		return 0
	} else {
		l.PushBack(&Pair{root, 1})
	}

	min_depth := math.MaxInt32

	for l.Len() > 0 {
		e := l.Front()
		l.Remove(e)
		node := e.Value.(*Pair).node
		current_depth := e.Value.(*Pair).depth
		if node.Left == nil && node.Right == nil {
			min_depth = min(min_depth, current_depth)
		} else {
			if node.Left != nil {
				l.PushBack(&Pair{node.Left, current_depth + 1})
			}
			if node.Right != nil {
				l.PushBack(&Pair{node.Right, current_depth + 1})
			}
		}
	}
	return min_depth
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
