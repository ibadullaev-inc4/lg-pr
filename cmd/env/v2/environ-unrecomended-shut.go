package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}

func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}
