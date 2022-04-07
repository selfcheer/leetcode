package tree

import "fmt"

/**
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	nodes := []*TreeNode{root}


	for len(nodes) != 0 {
		tempVals := make([]int, 0)
		count := len(nodes)
		for i := 0; i < count; i++ {
			node := nodes[0]
			nodes = nodes[1:]
			tempVals = append(tempVals, node.Val)
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		ans = append(ans, tempVals)
	}
	return ans
}

func TestLevelOrder() {
	nums := []interface{} {3,9,20,nil,nil,15,7}
	root := SliceToTree(nums)
	result := levelOrder(root)
	fmt.Println(result)
}