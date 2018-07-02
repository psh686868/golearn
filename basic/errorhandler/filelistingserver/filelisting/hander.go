package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
	"fmt"
)

const prefix = "/file/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (ur userError) Message() string {
	return string(ur)
}

func HandlerFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path

	if index := strings.Index(path, prefix); index == -1 {
		return userError(fmt.Sprintf("path %s must start with %s", path, prefix ))
	}

	filepath := path[len("/file/"):]

	file, err := os.Open(filepath)
	if err != nil {

		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(bytes)
	return nil
}
