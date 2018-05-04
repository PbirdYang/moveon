package main

import (
	"net/http"
	"fmt"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("yangmozi,move on!!!!"))
}

func main() {
	http.Handle("/", &helloHandler{})
	//fmt.Print([1]string{"start listening on port"})
	var port = "12345"
	fmt.Printf("listen on port %s",port)
	http.ListenAndServe(":12345", nil)
}