package array

import "fmt"

// ############### 数组   数组是值类型 即   注意 拷贝   一般不直接使用数组 ，使用的是切片

func printArray (arr [5]int)  {
	fmt.Println(" print array ## ")
	for i,v := range arr{
		fmt.Printf("arr1 i is %d and v is %d \n" , i,v)
	}
	arr[0] = 100
}
func printArrayAdd (arr *[5]int)  {
	fmt.Println(" print array 传递引用 ## ")
	for i,v := range arr{
		fmt.Printf("arr1 i is %d and v is %d \n" , i,v)
	}
	arr[0] = 100
}

func practicefor() {
	var arr1 [5]int;
	arr2 := [3]int{6, 7, 9}
	arr3  := [...]int {  2, 3, 4, 5, 6}

	for i,v := range arr1{
		fmt.Printf("arr1 i is %d and v is %d \n" , i,v)
	}
	for _,v := range arr2 {
		fmt.Printf("arr2  v is %d \n" ,v)
	}

	for v := range arr3 {
		fmt.Printf("arr3  v is %d \n" ,v)
	}

	printArray(arr1)
	fmt.Println("######## arr3 #########")
	printArray(arr3)
	// printArray(arr2) 这样写会报错

	fmt.Println("######## 数组值传递 所以打印仍然是原来的数 #########")
	for i,v := range arr1{
		fmt.Printf("arr1 i is %d and v is %d \n" , i,v)
	}

	for i := range arr3 {
		fmt.Printf("arr3  i is %d \n" ,i)
	}

	fmt.Println("######## 可以传递数组的引用 #########")
	printArrayAdd(&arr3)
	for i,v := range arr3 {
		fmt.Printf("arr3 i is %d v is %d \n" ,i,v)
	}
}

// slice 数组的切片
func slice()  {
	arr	:= [...]int {0,1,2,3,4,5,6,7}
	//slicetest1(arr)
	//slicetest2(arr)
	sliceAppend(arr)
	fmt.Println("end and arr is ",arr)
}
func slicetest1(arr [8]int)  {
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[:]", arr[:])

	fmt.Println("a1 arr[2:6]", arr[2:6])
	fmt.Println("a2 arr[2:]", arr[2:])

	updateSlice(arr[2:6])
	fmt.Println(" after slices a1 ")
	fmt.Println("a1 arr[2:6]", arr[2:6])

	updateSlice(arr[3:6])
	fmt.Println(" after slices a3 ")
	fmt.Println("a3 arr[3:6]", arr[3:6])

	fmt.Println(" and slicetest1 is ",arr)


}

// why ? 注意 pre ， length ，cap     2， 3， 4 ， 5， 6， 7
func slicetest2(arr [8]int)  {
	fmt.Println("s start is ",arr)
	s1 := arr[2:6]
	//fmt.Println("s[4] is ",s1[4])
	s2 := s1[3:5]
	fmt.Println("s1 [2:6] is ",s1)
	fmt.Println("s2 = s1[3:5] is ",s2)
	//因为s1 是数组 arr的一个视图 ，当 s1 := arr[2:6] 虽然s1打印的是  [2 3 4 5] 但是 s1 真正的数据是 [2 3 4 5 6 7] 但是我们不能直接访问
	// 6,7 这两个数据。 但是再次slice 即 s2 = s1[3:5] 取到 [5 6]
}


func updateSlice(arr []int)  {
	arr[0] =100
}

// 添加元素时如果超越 cap 系统会分配更大的底层数组
func sliceAppend(arr [8]int)  {

	fmt.Println("arr is ",arr)
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println("s1 [2:6] is ",s1)
	fmt.Println("s2 = s1[3:5] is ",s2)

	s3 := append(s2, 20)
	fmt.Println("s3 = s2 append 20 is ",s3)

	fmt.Println("s1 end is ",s1)
	s4 := s1[3:6]
	fmt.Println("s4 end is ",s4)
	fmt.Println("sliceAppend arr end is ",arr)

	s7 := append(s3, 30)
	s8 := append(s7, 40)
	s9 := append(s8, 40)

	fmt.Println("s9 end is ",s9)
	fmt.Println("sliceApp appreng end arr end is ",arr)
}

func main() {
	//practicefor()
	slice()

}
