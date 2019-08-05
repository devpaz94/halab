package handler

import (
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", handleGETSplashBath)
	http.HandleFunc("/", handlePOSTSaveBath)
}
