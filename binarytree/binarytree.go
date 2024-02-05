package binarytree

// 创建二叉事节点结构体
type BinaryTreeNode struct {
	no int
  left *BinaryTreeNode // 左子节点
  right *BinaryTreeNode // 右子节点
}

// 前序遍历
