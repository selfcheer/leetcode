package tree

type TreeNode struct {
	Val  int
	Left *TreeNode
	Right *TreeNode
}



func SliceToTree(nums []interface{}) (root *TreeNode) {
	if len(nums) == 0 {
		return nil
	}
	root = &TreeNode{Val: nums[0].(int)}
	nums = nums[1:]
	oneLevelNodes := []*TreeNode{root}
	for len(nums) != 0 || len(oneLevelNodes) != 0 {
		count := len(oneLevelNodes)
		for i:= 0; i < count; i ++ {
			node := oneLevelNodes[0]
			oneLevelNodes = oneLevelNodes[1:]
			if len(nums) == 0 {
				break
			}
			if nums[0] != nil {
				node.Left = &TreeNode{Val: nums[0].(int)}
				oneLevelNodes = append(oneLevelNodes, node.Left)
			}
			nums = nums[1:]
			if nums[0] != nil {
				node.Right = &TreeNode{Val: nums[0].(int)}
				oneLevelNodes = append(oneLevelNodes, node.Right)
			}
			nums = nums[1:]
		}
	}
	return root
}

func TreeToSlice(root *TreeNode) (nums []int) {
	nums = make([]int,0)
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		count := len(nodes)
		for i := 0; i < count; i++ {
			node := nodes[0]
			nodes = nodes[1:]
			nums = append(nums, node.Val)
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
	}
	return nums
}