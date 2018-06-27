package node

func (node *TreeNode) Traverse()  {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.PrintNode()
	node.Right.Traverse()
}