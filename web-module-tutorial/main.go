package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kd/learn/go/web-service/product"
)

type fooHandler struct {
	Message string
}

type foo struct {
	Message string `json:"message,omitempty"`
	Age     int    `json:"age,omitempty"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar called"))
}

func middlewareHandler(handlerFunc http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware start")
		start := time.Now()
		handlerFunc.ServeHTTP(w, r)
		fmt.Println("Middleware stop, running time:", time.Since(start))
	})
}

func main() {
	// Handling Http Requests
	// http Handle and Handle Func Demo

	http.Handle("/foo", &fooHandler{Message: "Hello World"})
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/products", product.BulkHandler)

	// using middleware handler
	http.Handle("/products/", middlewareHandler(product.Handler))
	http.ListenAndServe(":5000", nil)
}
