package binaryTree

import "slices"

//给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
//
//示例 1:
//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
//
//示例 2:
//输入: preorder = [-1], inorder = [-1]
//输出: [-1]

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	leftSize := slices.Index(inorder, preorder[0])
	left := buildTree(preorder[1:1+leftSize], inorder[:leftSize])
	right := buildTree(preorder[1+leftSize:], inorder[1+leftSize:])
	return &TreeNode{preorder[0], left, right}
}
