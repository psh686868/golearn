package main

import (
	"time"
	"fmt"
)

//  go 语言协程 coroutine 是编译器 编译时遇到 go 时
//  go 语言有自己的调度器
// 协程 让出控制权的 条件 ： 遇到io操作时， 如果一个协程不遇到io 不会交出控制权 下面程序代码将展示
//     我们可以自己让 goroutine 交出控制权 runtime.Gosched
// go run -race 可以查看多线程问题 如下面 go func （i int） 将 i int去掉
// 当一个 goroutine 写一个共享变量时 另一个goroutine去读它则出错


// goroutine 可能切换的点
// 1.  I/O ,select
// 2. channel
// 3. 等待锁
// 4. 函数调用 (有时)
// 5. runtime.GoSched()
//


func main() {
	var array [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for  {
				fmt.Printf("this is goroutine %d \n", i) // 因为是print 函数发生了 io 所以此时goroutine 让出cpu
				//array[i] = array[i] + 1 // 注释上面，写这一句此时goroutine 就永远不会让出cpu ，并且cpu 用的很多 runtime.Gosched()
				// runtime.Gosched()
				// 吐过
			}
		}(i)
	}

	time.Sleep(time.Minute)
	fmt.Println("this is array ", array)
}
