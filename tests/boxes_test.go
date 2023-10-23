package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBoxSizes_Success(t *testing.T) {
	res := createRequestContext("POST", "/packs", "validRequest.json")

	if status := res.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

	// convert the response body to map to read the value
	var arrayMap map[int]int
	err := json.Unmarshal([]byte(res.Body.String()), &arrayMap)
	if err != nil {
		fmt.Println("Unable to parse the error response")
	}
	if arrayMap[2000] != 1 &&
		arrayMap[250] != 1 &&
		arrayMap[2000] != 2 {
		t.Errorf("Box sizes calculated incorrectly")
	}
}

func TestBoxSizes_Success_Scenario2(t *testing.T) {
	res := createRequestContext("POST", "/packs", "validRequest-2.json")

	if status := res.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

	// convert the response body to map to read the value
	var arrayMap map[int]int
	err := json.Unmarshal([]byte(res.Body.String()), &arrayMap)
	if err != nil {
		fmt.Println("Unable to parse the error response")
	}
	if arrayMap[500] != 1 {
		t.Errorf("Box sizes calculated incorrectly")
	}
}

func TestBoxSizes_InvalidRequest_ItemsOrdered(t *testing.T) {
	res := createRequestContext("POST", "/packs", "invalidRequest-zero-items-ordered.json")

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

	if res.Body.String() != "Num of items to be packed must be greater than or equal to 1" {
		t.Errorf("Exception must be thrown in case of bad request")
	}
}

func TestBoxSizes_InvalidRequest_EmptyBoxSizes(t *testing.T) {
	res := createRequestContext("POST", "/packs", "invalidRequest-empty-box-sizes.json")

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

	if res.Body.String() != "Mandatory pack sizes not provided" {
		t.Errorf("Exception must be thrown in case of bad request")
	}
}

func createRequestContext(method string, url string, fileName string) *httptest.ResponseRecorder {
	data := JsonToMap(fileName)
	jsonString, _ := json.Marshal(data)
	reader := strings.NewReader(string(jsonString))

	req, _ := http.NewRequest(method, url, reader)
	res := httptest.NewRecorder()

	ComputeBoxSize(res, req)
	return res
}
