package handler

import (
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", handleGETBaths)
}
