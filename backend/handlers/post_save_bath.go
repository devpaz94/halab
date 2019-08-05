package handler

import (
	"encoding/json"
	"net/http"

	"github.com/halab/backend/dto"
	"github.com/halab/backend/models"
)

func handlePOSTSaveBath(w http.ResponseWriter, req *http.Request) {

	bath := dto.Bath{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&bath)
	if err != nil {
		panic(err)
	}

	err = models.CreateBath(&bath)
	if err != nil {
		panic(err)
	}

	resp, err := json.Marshal(&bath)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
