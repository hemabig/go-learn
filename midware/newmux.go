package main

import "net/http"

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
	//http.ListenAndServe()
}
