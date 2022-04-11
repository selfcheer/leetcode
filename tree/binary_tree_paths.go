package tree

import (
	"fmt"
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	dfs(root, &paths, "")
	return paths
}
func dfs(root *TreeNode, paths *[]string, pre string) {
	if root == nil {
		return
	}
	pre += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, pre)
	}
	dfs(root.Left, paths, pre + "->" )
	dfs(root.Right, paths, pre + "->")
}

func TestBinaryTreePaths() {
	input := []interface{}{1,2,3,nil,5}
	root := SliceToTree(input)
	result := binaryTreePaths(root)
	fmt.Println(result)
}
