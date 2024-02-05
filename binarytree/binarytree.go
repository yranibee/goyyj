package binarytree

// 创建二叉事节点结构体
type BinaryTreeNode struct {
	no int
  left *BinaryTreeNode // 左子节点
  right *BinaryTreeNode // 右子节点
}

// 前序遍历
func preOrder(node *BinaryTreeNode) {
  if node == nil {
    return
  }
  // 当前节点
  fmt.Println(node)
  // 遍历左子树
  preOrder(node.left)
  // 遍历右子树
  preOrder(node.right)
}

// 中序遍历
func infixOrder(node *BinaryTreeNode) {
  if node == nil{
    return
  }
  // 遍历左子树
  infixOrder(node.left)
  // 当前节点
  fmt.Println(node)
  // 遍历右子树
  infixOrder(node.right)
}

// 后续遍历
func postOrder(node *BinaryTreeNode) {
  if node == nil {
    return
  }
  // 遍历左子树
  postOrder(node.left)
  // 遍历右子树
  postOrder(node.right)
  // 当前节点
  fmt.Println(node)
}

// 手动初始化节点函数
func initNode() *BinaryTreeNode {
  root := &BinaryTreeNode{no: 1}
  node2 := &BinaryTreeNode{no: 2}
  node3 := &BinaryTreeNode{no: 3}
  node4 := &BinaryTreeNode{no: 4}
  node5 := &BinaryTreeNode{no: 5}

  // 手动建立树的关系
  root.left = node2
  root.right = node5
  node2.left = node3
  node2.right = node4

  return root
}