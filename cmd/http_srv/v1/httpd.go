package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if name == "" {
		name = "Indigo Montoya"
	}
	fmt.Println(req.Method)
	fmt.Fprint(res, "Hello my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Indigo Montoya"
	}
	fmt.Fprint(res, "Goodbye ", name)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage")
}
