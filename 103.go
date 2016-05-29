/*
* 103. Binary Tree Zigzag Level Order Traversal
 */
package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func loop(result [][]int, tree []*TreeNode, dir int) [][]int {
	var childs []*TreeNode
	var treeLength = len(tree)
	var i int
	if treeLength == 0 {
		return result
	}

	for i = 0; i < treeLength; i++ {
		var root = tree[i]
		if root.Left != nil {
			childs = append(childs, root.Left)
		}
		if root.Right != nil {
			childs = append(childs, root.Right)
		}
	}
	var newdir int
	var subResult []int
	if dir == 1 { //正序
		newdir = 0
		for i = 0; i < treeLength; i++ {
			var root = tree[i]
			subResult = append(subResult, root.Val)
		}
	} else { //反序
		newdir = 1
		for i = treeLength - 1; i >= 0; i-- {
			var root = tree[i]
			subResult = append(subResult, root.Val)
		}
	}
	result = append(result, subResult)
	return loop(result, childs, newdir)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	var tree []*TreeNode
	if root == nil {
		return result
	}
	tree = append(tree, root)
	return loop(result, tree, 1)
}

func maketreenode(num int, index *int) *TreeNode {
	if num > 5 {
		return nil
	}
	var node TreeNode
	*index = *index + 1
	node.Val = *index
	node.Left = maketreenode(num+1, index)
	node.Right = maketreenode(num+1, index)
	return &node

}
func makecase() *TreeNode {
	var start = 1
	return maketreenode(1, &start)
}

/*
* test
 */
func main() {
	var root *TreeNode = makecase()
	var result [][]int = zigzagLevelOrder(root)
	fmt.Println(result)
}
