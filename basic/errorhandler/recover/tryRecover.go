package main

import (
	"fmt"
)

func Recover() {
	// recover 跟在 defer 后面
	defer func() {
		r := recover()
		if r, ok := r.(error); ok {
			fmt.Printf(" error is occurred %q \n", r.Error())
		} else {
			panic(r)
		}
	}()

	//panic(errors.New(" this is an error"))

	b := 0
	a := 32 / b
	fmt.Println(a)

}
func main() {
	Recover()
}
