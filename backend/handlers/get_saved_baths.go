package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/halab/backend/models"
)

func handleGETSavedBaths(w http.ResponseWriter, req *http.Request) {

	baths, err := models.ReadBaths()
	if err != nil {
		fmt.Println("error")
	}

	resp, err := json.Marshal(&baths)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
