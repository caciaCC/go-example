package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	ans := &strings.Builder{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans.WriteString(strconv.Itoa(root.Val))
		if root.Left != nil {
			ans.WriteByte('(')
			dfs(root.Left)
			ans.WriteByte(')')
			if root.Right != nil {
				ans.WriteByte('(')
				dfs(root.Right)
				ans.WriteByte(')')
			}
		} else if root.Right != nil {
			ans.WriteByte('(')
			ans.WriteByte(')')
			ans.WriteByte('(')
			dfs(root.Right)
			ans.WriteByte(')')
		}
	}
	dfs(root)
	return ans.String()
}

func convertBinaryTreeToString(root *TreeNode) string {
	ans := &strings.Builder{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans.WriteString(strconv.Itoa(root.Val))
		if root.Left != nil {
			ans.WriteByte('(')
			dfs(root.Left)
			ans.WriteByte(')')
			if root.Right != nil {
				ans.WriteByte('(')
				dfs(root.Right)
				ans.WriteByte(')')
			}
		} else if root.Right != nil {
			ans.WriteByte('(')
			ans.WriteByte(')')
			ans.WriteByte('(')
			dfs(root.Right)
			ans.WriteByte(')')
		}
	}
	dfs(root)
	return ans.String()
}

func main() {
	node := &TreeNode{}
	fmt.Println(tree2str(node))
	//	函数命名追求表意清晰
	fmt.Println(convertBinaryTreeToString(node))
}
