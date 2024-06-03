package main

import (
	greet "hello/dependency_injection"
	"log"
	"net/http"
)

func GreetingHandler(writer http.ResponseWriter, request *http.Request) {
	greet.Greet(writer, request.Host)
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreetingHandler)))
}
