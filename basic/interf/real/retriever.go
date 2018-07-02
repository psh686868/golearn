package real

import (
	"time"
	"net/http"
	"net/http/httputil"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
} 

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	result, e := httputil.DumpResponse(resp, true)

	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()
	return string(result)

}
