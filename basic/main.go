package main

import (
	"fmt"
	"github.com/golearn/basic/queue"
)

// 他们是 作用域是包内变量
// 而且只能使用 var 这个定义 不能使用 ：=

var aaa = 1
var sss = "fffff"

var (
	aaaa = 2
	ssss = "fdasfa"
)

const(
	cfilename3 = "/data/shitou"
	cage3 = 7
)

func variableZeroVale() {
	var aa int
	var bb string
	var cc string
	fmt.Printf("aa is %d , bb is %q , cc is %q ", aa, bb, cc)
	fmt.Print("\n")
}

func variableInitialValue() {
	a, b := 1, 2
	var s string = "hellp word"
	fmt.Printf("a is %d, b is %d s is %q \n", a, b, s)

}

//常量
func consts()  {
	const cfilename = "/data/shitou"
	const cage = 3    //  const cage    常量必须定义 类比 java   static
	const (
		cfilename2 = "fdas"
		cage2 = 5
	)
	fmt.Printf("consts cfilename is %q and cage is %d \n" , cfilename, cage)
	fmt.Printf("consts cfilename2 is %q and cage2 is %d \n" , cfilename2, cage2)
	fmt.Printf("consts cfilename3 is %q and cage3 is %d \n" , cfilename3, cage3)
	//cfilename2 = "fdas"      error 定义后不能改变值
	//cage3 = 7  	error
}

//枚举类型 就是定义一组常量
func enums()  {
	const (
		java = 1
		cpp  = 2
		python = 3
	)
	const (
		kafka = iota  //const declare and start 0  自增
		spring
		mysql
	)


	fmt.Printf("enums language is %d, %d, %d \n" , java , cpp, python)
	fmt.Printf("enums is is %d, %d, %d \n " , kafka , spring, mysql)
}

func main() {
	//variableZeroVale()
	//variableInitialValue()
	//fmt.Printf("aaa is %d \n", aaa)
	//fmt.Printf("aaaa is %d \n", aaaa)
	//fmt.Printf("ssss is %d \n", ssss)
	//consts()
	//enums()
	empty := false
	queues := queue.Queue{}
	queues.Push(2)
	queues.Push(4)
	queues.Push(6)
	queues.Pop()
	queues.Pop()
	empty = queues.IsEmpty()
	fmt.Println(" empty is " , empty)
	fmt.Println()
	queues.Pop()
	empty = queues.IsEmpty()
	fmt.Println(" empty is " , empty)
}
