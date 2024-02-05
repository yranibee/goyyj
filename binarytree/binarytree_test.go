package binarytree

func TestOrder(t *testing.T) {
	root := initNode()

	fmt.Println("==========前序遍历==========")
	preOrder(root)

}