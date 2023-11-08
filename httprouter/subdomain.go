package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HostSwitch map[string]http.Handler

// Implement the ServeHTTP method on our new type
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if a http.Handler is registered for the given host.
	// If yes, use it to handle the request.
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
		fmt.Printf("we find the host %s\n", r.Host)
	} else {
		// Handle host names for which no handler is registered
		http.Error(w, "Forbidden", 403) // Or Redirect?
	}
}

func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "hello localhost\n")
}

func main() {
	// Initialize a router as usual
	router := httprouter.New()
	router.GET("/hello", Hello)

	// Make a new HostSwitch and insert the router (our http handler)
	// for example.com and port 12345
	hs := make(HostSwitch)
	hs["localhost:8080"] = router

	// Use the HostSwitch to listen and serve on port 12345
	log.Fatal(http.ListenAndServe(":8080", hs))
}
