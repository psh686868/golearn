package main

import (
	"net/http"
	"github.com/golearn/basic/errorhandler/filelistingserver/filelisting"
	"os"
	"log"
)

type apphandler func(writer http.ResponseWriter, request *http.Request) error


type userError interface {
	error
	Message() string
}

func errWrapper(handler apphandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		e := handler(writer, request)
		if e != nil {
			code := http.StatusOK
			log.Printf("Error occurred "+
				"handling request: %s",
				e.Error())

			if e, ok := e.(userError); ok {
				 http.Error(writer, e.Message(), http.StatusBadRequest)
				 return
			}


			switch {
			case os.IsNotExist(e):
				code = http.StatusNotFound
			case os.IsPermission(e):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {

	http.HandleFunc("/", errWrapper(filelisting.HandlerFileList))
	err := http.ListenAndServe(":7888", nil)
	if err != nil {
		panic(err)
	}
}
