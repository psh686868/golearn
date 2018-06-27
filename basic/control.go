package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
	"reflect"
	"runtime"
)

const filename = "/Users/shitou/wokespace/golang/src/github.com/golearn/basic/abc.txt"

// ##############  if 流程  go 与 java 比 特点 ，1. if 条件表达式可以不写 () 2. 可以前面写一些语句 让后 条件表达式
func ifbranch() {

	filecontent, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", filecontent)
	}

	// or we could code in this way
	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", content)
	}
}

// ################## switch 与 java 比的特点   case 后面不用加 break 默认 break ， switch 内容 {} 内容可以 不用写因为 case 里写了
func swichbranch() {
	grade(70)
}
func grade(score int) {
	g := ""
	//case 会默认break
	switch {
	case score < 0 || score > 100:
		panic("score is error")
	case score < 60:
		g = "FFF"
	case score < 80:
		g = "BBB"
	case score <= 100:
		g = "AAA"
	}
	fmt.Printf("score is %q \n", g)
	switch score {
	case 20:
		g = "FFF"

	}
}

// for 1. go语言里没有 while 和 都while   2. go 语言的 死循环 go语言经常使用死循环进行并发编程

// ############# 函数式编程

func apply(op func(int, int) int, a, b int) int {
	value := reflect.ValueOf(op)
	fmt.Println(value)
	p := value.Pointer();
	fmt.Println(p)
	funName := runtime.FuncForPC(p).Name()
	fmt.Printf("funName is %q \n", funName)
	return op(a, b)
}

func printFile() {
	file, err := os.Open(filename);
	if err != nil {
		fmt.Print(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

// 死循环
func forever() {
	for {
		fmt.Println("aaa")
	}
}





func main() {
	ifbranch()
	swichbranch()
	printFile()
	//forever()
	fmt.Println("######## function #########")
	i := apply(func(a int, b int) int {
		return a + b
	}, 1, 2)
	fmt.Println(" apply return i is : %d \n", i)
	fmt.Println("######## for  #########")

}
