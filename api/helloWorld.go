package api

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello world!")
}

func helloRoute() error {
	http.HandleFunc("/api/hello", hello)
	return nil
}
