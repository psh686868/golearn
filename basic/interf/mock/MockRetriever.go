package mock

type Retriever struct {
	Content string 
}

// 只要 和main的Retriever 接口的方法相同 就实现了retriever
func (r *Retriever) Get(url string) string{

	return r.Content
}

func (r *Retriever) Post (url string, form map[string]string) string {
	r.Content = form["content"]
	return "ok"
}