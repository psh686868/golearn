package main

import (
	"fmt"
	"os"
	"bufio"
	"github.com/golearn/basic/functional/fib"
)

// defer 执行顺序 ，先进后出的执行顺序
func tryDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	panic(" error occurred ")
}

//
func writerFile(s string) {

	file, err := os.OpenFile(s, os.O_EXCL, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(pathError)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fibonacci := fib.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, fibonacci())
	}

}

func main() {
	//tryDefer()
	writerFile("aasb.txt")
}
