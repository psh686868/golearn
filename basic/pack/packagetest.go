package pack

import (
	"github.com/golearn/basic/entry/node"
	"github.com/golearn/basic/entry"
)

// 1. 首字母大写代表 public 首字母小写代表 private 这适用于所有的数据 ，如 方法， 常量，定义，struct ，，，， 。 它 2。 针对谁 ！！！ 针对于 包
// 也就是针对 package ， 每个目录只能有一个包 ， 目录名不一定和包名相同， 为结构定义的方法必须放在同一个包内，但是可以是不同的文件，
func main()  {
	//var root node.TreeNode
	//root.Value = 2
	//root.Left = &node.TreeNode{}
	//fmt.Println("1 root is  ", root)
	//root.Right = &node.TreeNode{5,nil,nil}
	//fmt.Println("2root is ", root)
	////root.right.left = new (treeNode)
	//root.Right.Left = node.CreateNode(3)
	//fmt.Println("3root is ", root)
	////root.right = new(treeNode {12,nil,nil})
	//
	//nodes := []node.TreeNode {
	//	{3,nil,nil},
	//	{},
	//	{6,nil,nil},
	//}
	//fmt.Println("nodes is " , nodes)
	//
	//root.PrintNode()


	// we will cread a tree ,it root is
	//   			3
	//			0      5
	//		2			0

	rootTree := node.TreeNode{Value:3}

	rootTree.Left = &node.TreeNode{Value:0}
	rootTree.Left.Left = node.CreateNode(2)
	rootTree.Right = node.CreateNode(5)
	rootTree.Right.Right = node.CreateNode(0)
	//fmt.Println("traverse")
	//rootTree.Traverse()

	treeyNode := entry.MyTreeyNode{&rootTree}
	treeyNode.PostOrder()
}