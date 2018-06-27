package queue

type Queue []int

func (qu *Queue) Push (v int){
	//fmt.Println("push qu *Quence is ", *qu)
	//.Println("push qu Quence is ",  qu)
	*qu = append(*qu, v)
}

func (qu *Queue)Pop() int {
	//fmt.Println("pop (*qu)  is ", *qu)
	//fmt.Println("pop (qu)  is ", qu)
	head := (*qu)[0]
	*qu = (*qu)[1:]
	return head
}

func (qu *Queue) IsEmpty() bool  {
	len := len(*qu)
	return len == 0
}
