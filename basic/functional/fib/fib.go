package fib

// 斐波那列数列  使用闭包实现  1，  2， 3，  5
//  						a,  b
//								a, a+b
// 

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
