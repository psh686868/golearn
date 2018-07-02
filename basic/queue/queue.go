package queue

type Queue []interface{}

func (qu *Queue) Push (v interface{}){
	//fmt.Println("push qu *Quence is ", *qu)
	//.Println("push qu Quence is ",  qu)
	*qu = append(*qu, v)
}

func (qu *Queue)Pop() interface{} {
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
