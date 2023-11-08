package main

import (
	"fmt"
	"net/http"
	"time"
)

type midware func(http.Handler) http.Handler

type Router struct {
	midwarechain []midware
}

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello\n"))
}

func (r *Router) Use(m midware) {
	r.midwarechain = append(r.midwarechain, m)
}

func (r *Router) Add(meth string, handle http.Handler) {
	var merg = handle
	for i := len(r.midwarechain) - 1; i >= 0; i-- {
		merg = r.midwarechain[i](merg)
	}

	http.DefaultServeMux.Handle(meth, merg)
}
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		fmt.Println(timeElapsed)
	})
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		fmt.Println("we are in log middware")
		next.ServeHTTP(wr, r)
	})
}

func main() {
	var r Router
	r.Use(timeMiddleware)
	r.Use(logMiddleware)
	r.Add("/", http.HandlerFunc(hello))
	http.ListenAndServe(":8080", nil)
}
