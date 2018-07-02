package main

import (
	"fmt"

	"github.com/golearn/basic/interf/mock"
	"github.com/golearn/basic/interf/real"
	"time"
)

const url = "http://www.taofangdd.com"

type Retriever interface {
	// get
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func post(poster Poster) {
	s := poster.Post(url,
		map[string]string{
			"name": "myname",
			"book": "go inaction",
		})
	fmt.Println(s)
}

// download是使用者
func downLoad(r Retriever) string {
	return r.Get(url)
}

// 接口类型包括实现者的值和实现者的 指针
func main() {
	var r Retriever
	r = &mock.Retriever{Content: "this is download"}
	load := downLoad(r)
	fmt.Println(load)
	fmt.Printf(" r is %T, %v is \n", r, r)

	r = &real.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	fmt.Printf(" r is %T, %v is \n", r, r)

	post(&mock.Retriever{})
	//response := r.Get(url)
	//fmt.Println(response)
}

// go语言 判断接口类型
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch  r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:")
	case *real.Retriever:
		fmt.Println("UserAgent:")
	}
}
