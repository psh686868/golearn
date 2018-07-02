package entry

import (
	"github.com/golearn/basic/entry/node"
)

type MyTreeyNode struct {
	Node *node.TreeNode
}

func (myNode *MyTreeyNode) PostOrder()  {
	if myNode == nil || myNode.Node == nil {
		return
	}
	treeyNode := MyTreeyNode{myNode.Node.Right}
	//fmt.Println("treeNode is ", treeyNode)
	treeyNode.PostOrder()
	myNode.Node.PrintNode()
	myTreeyNode := MyTreeyNode{myNode.Node.Left}
	myTreeyNode.PostOrder()

}
// 对treeNode的进一步封装