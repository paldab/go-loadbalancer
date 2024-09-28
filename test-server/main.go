package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type LoadbalancerHandler struct{
    http.Handler
}

func handleServerPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = ":3000"
	}

	if !strings.Contains(port, ":") {
		port = ":" + port
	}

	return port
}

var serverPort string = handleServerPort()
func (lb LoadbalancerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request received on port %s %s %s\n", serverPort, r.Method, r.URL.Path)
	w.Write([]byte("Request received on test server"))
}

func main() {
	mux := http.NewServeMux()

	lb := LoadbalancerHandler{}

	mux.Handle("/", lb)
	fmt.Println("Listening on port: ", serverPort)
	http.ListenAndServe(serverPort, mux)
}
