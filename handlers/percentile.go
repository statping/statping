package handlers

import (
	"github.com/statping/statping/types/hits"
	"encoding/json"
	"net/http"
	"io/ioutil"
	// "fmt"
)

func apiPercentileHandler(r *http.Request) interface{} {
	response := apiResponse{
		Status: "success",
		Output: hits.CurrentPercentile.Rank,
	}
	return response
}

func apiPercentileUpdateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	err = json.Unmarshal(body, &hits.CurrentPercentile)
    if err != nil {
		sendErrorJson(err, w, r)
		return
	} else {
		response := apiResponse{
			Status: "success",
			Output: hits.CurrentPercentile.Rank,
		}
		returnJson(response, w, r)
	}
}
