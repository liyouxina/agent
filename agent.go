package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Handler struct {

}

func (*Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	if body != nil {
		fmt.Println(string(body))
	}
}

func main() {
	http.ListenAndServe("0.0.0.0:8899", &Handler{})
}