package main

import (
	"fmt"
	"github.com/braintree/manners"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	handler := newHandler()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	manners.ListenAndServe(":8080", handler)
}

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "Indigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}
