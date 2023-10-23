package controller

import (
	"Test/service"
	"encoding/json"
	"fmt"
	"net/http"
)

// ComputeBoxSize Handles rest endpoint "/packs"
func ComputeBoxSize(w http.ResponseWriter, r *http.Request) {

	var request RequestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httpStatus(w, http.StatusBadRequest, "There was an error decoding the request body into the struct")
		return
	}

	numOfItems := request.ItemsOrdered
	if numOfItems < 1 {
		httpStatus(w, http.StatusBadRequest, "Num of items to be packed must be greater than or equal to 1")
		return
	}

	boxSizes := request.BoxSizes
	if len(boxSizes) == 0 {
		httpStatus(w, http.StatusBadRequest, "Mandatory pack sizes not provided")
		return
	}

	err = json.NewEncoder(w).Encode(service.CalculatePackSizes(boxSizes, numOfItems))
	if err != nil {
		httpStatus(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func httpStatus(w http.ResponseWriter, status int, errMsg string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, errMsg)
}
