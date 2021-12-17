package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func hello(w http.ResponseWriter, req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(requestDump))
}

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":80", nil)
}
