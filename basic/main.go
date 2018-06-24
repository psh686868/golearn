package main

import (
	"fmt"
)

// 他们是 作用域是包内变量
// 而且只能使用 var 这个定义 不能使用 ：=

var aaa = 1
var sss = "fffff"

var (
	aaaa = 2
	ssss = "fdasfa"
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

func main() {
	variableZeroVale()
	variableInitialValue()
	fmt.Printf("aaa is %d \n", aaa)
	fmt.Printf("aaaa is %d \n", aaaa)
	fmt.Printf("ssss is %d \n", ssss)
}
