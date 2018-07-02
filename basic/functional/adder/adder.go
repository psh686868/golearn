package main

import "fmt"

// go 的闭包 ，一个函数有A 一些，变量 sum ， 然后返回值是一个 函数并且返回的这个函数 使用了 A里面的 sum。
//  当我们调用 refunc1 = A（） refunc (） refunc1() A里面的变量一直被 refunc （）使用， 2️而 refunc2 = A（）则会开辟一个
// java的闭包如何表示呢

//public static void main(String[] args) {
//	Function<Integer, Integer> adder = adder();
//	Integer sum1 = adder.apply(1);
//	Integer sum2 = adder.apply(2);
//	Integer sum3 = adder.apply(3);
//	Function<Integer, Integer> adder2 = adder();
//	Integer adder2Sum1 = adder2.apply(1);
//	System.out.println("sum1 = " + sum1 + "; sum2 = " + sum2 + "; sum3 = " + sum3 + "; adder2Sum1 = " + adder2Sum1);
//
//	sum1 = 1; sum2 = 3; sum3 = 6; adder2Sum1 = 1
//}
//
//public static Function<Integer, Integer> adder() {
//	final Holder<Integer> sum = new Holder<>(0);
//	return (Integer value) -> {
//	sum.value +=  value;
//	return sum.value;
//	};
//}

func Adder() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}




func main() {
	add := Adder()
	for i := 0; i< 5; i++ {
		sum := add(i)
		fmt.Printf(" 0 + 1+ ....,  i is %d sum is %d \n" , i , sum)
	}

	add1 := Adder()
	add1(1)
	sum1 := add1(2)
	fmt.Printf("sum1 is %d \n", sum1)

	add2 := Adder()
	sum2 := add2(1)
	fmt.Printf("sum2 is %d \n", sum2)
}
