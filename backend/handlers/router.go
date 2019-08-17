package handler

import (
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/spash_baths", handleGETSplashBath)
	http.HandleFunc("/save_bath", handlePOSTSaveBath)
}
