package tree

import "fmt"

func inorderTraversal(root *TreeNode) []int {
	var res []int
	for root != nil {
		if root.Left != nil {
			pre := root.Left
			for pre.Right != nil && pre.Right != root {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = root
				root = root.Left
			}else {
				root = pre.Right
				res = append(res, root.Val)
				pre.Right = nil
				root = root.Right
			}
		}else {
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}


func TestInorderTraversal() {
	sli := []interface{}{1,nil,2,3}
	fmt.Println(sli)
	root := SliceToTree(sli)
	nums := inorderTraversal(root)
	fmt.Println(nums)
}

