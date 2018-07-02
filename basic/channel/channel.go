package main

import "fmt"

//   				->
//   goroutine             goroutine
//					<-
//     调度器
//
// channnel 是goroutine 和 goroutine 的交互


func chandeme()  {
	c := make(chan int)
	c <- 1
	c <- 2
	n := <- c
	fmt.Println(n)

}

func main() {
	chandeme()
}