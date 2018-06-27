package node

import "fmt"

type TreeNode struct {
	Value int
	Left , Right *TreeNode
}
//
// 1.方法调用是值传递， 要改变其值使用指针  2. 结构过大也使用指针接收者， 3. 一致性 ：如果有指针接收者最好都是指针接收者
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value:value}
}
// 这样我们就可以使用 treeNode.method 了这种 (treenode treeNode) printNode()  method 和 printNode(treenode treeNode) 没有什么区别，而且嗾使值传递
func (treenode TreeNode) PrintNode()  {
	fmt.Print(treenode.Value, "  ")
}
// 这样调用的时候 会把调用者的地址传过来, 调用的时候就是 root.setValue(2)  会把 root 的地址传来
func (treenode *TreeNode) SetValue(value int)  {
	treenode.Value = value
}


// 使用工厂方法创建一个 treeNode 注意返回的是局部变量的地址  ， c++这样会挂， java new 堆分配 ，
// go 语言的这个 node 是在栈上分配还是在堆上分配的呢， 我们不需要知道😔，因为go语言会根据编译时确定 到底在哪分配
//  如果 return treeNode{value:value} 可能是在栈上分配，return &treeNode{value:value} 在堆上分配，
// go 通过逃逸分析来指定值到底分配到哪里 我们可以使用 go run -gcflags '-m -l' 查看分配到哪里了 http://legendtkl.com/2017/04/02/golang-alloc/
// 和http://goog-perftools.sourceforge.net/doc/tcmalloc.html
func createNode(value int) *TreeNode{

	return &TreeNode{Value:value}
}


