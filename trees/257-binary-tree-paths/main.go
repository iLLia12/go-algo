package main

import "strconv"

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

//Solution 1
func dfs(node *TreeNode, path string, paths *[]string) {
	if node != nil {
		if path != "" {
			path += "->"
		}
		path += strconv.Itoa(node.val)
		if node.left == nil && node.right == nil {
			*paths = append(*paths, path)
		} else {
			dfs(node.left, path, paths)
			dfs(node.right, path, paths)
		}
	}
}

func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	dfs(root, "", &paths)
	return paths
}

//Solution 2
func binaryTreePaths2(root *TreeNode) []string {
	var res []string
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node == nil { return }
		var val string
		val = strconv.Itoa(node.val)
		path = path + val
		if node.left == nil && node.right == nil {
			res = append(res, path)
		}
		dfs(node.left, path + "->")
		dfs(node.right, path + "->")
	}
	dfs(root, "")
	return res
}

func main() {
	node := &TreeNode{1, nil, nil}
	node.left = &TreeNode{2, nil, nil} 
	node.left.right = &TreeNode{5, nil, nil}
	node.right = &TreeNode{3, nil, nil}
	binaryTreePaths(node)
}