package main

import (
	"github.com/golearn/basic/functional/fib"
	
	"fmt"
	"io"
	"bufio"
	"strings"
)

type intGenf func() int

func (g intGenf)  Read (p []byte) (n int, err error) {
	v := g()
	if (v > 10000) {
		return 0 , io.EOF
	}
	s := fmt.Sprintf("%d\n", v)

	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

	// 循环在这
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fibonacci := fib.Fibonacci()

	for i := 0; i< 10 ; i++ {
		fmt.Println(fibonacci())
	}

	fmt.Printf("#######################")
	var f intGenf  = fib.Fibonacci()
	printFileContents(f)

}
