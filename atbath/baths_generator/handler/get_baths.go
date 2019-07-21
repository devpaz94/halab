package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleGETBaths(w http.ResponseWriter, r *http.Request) {

	URL := "https://api.unsplash.com/photos/?client_id="

	AccessKey := "af9387e4fa060d133ddc10e21cd1b42c0003bbffc1256ef15173b4074aabe54f"

	resp, err := http.Get(URL + "/photos/random&client_id=" + AccessKey +)
	if err != nil {
		// handle error
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Fprintf(w, string(body))
}
