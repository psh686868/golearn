package entry

import "github.com/golearn/basic/entry/node"

type MyTreeyNode struct {
	Node *node.TreeNode
}

func (myTreeNode *MyTreeyNode) postOrder()  {
	if myTreeNode == nil || myTreeNode.Node == nil {
		return
	}
	myTreeNode.Node.Right.Traverse()
	myTreeNode.Node.PrintNode()
	myTreeNode.Node.Left.Traverse()

}
// 对treeNode的进一步封装