package main

import (
	"log"
	"net/http"

	handler "github.com/halab/atbath/baths_generator/handler"
)

func main() {
	handler.HandleRequests()

	log.Fatal(http.ListenAndServe(":8881", nil))
}
