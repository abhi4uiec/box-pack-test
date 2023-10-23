package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

type RequestBody struct {
	ItemsOrdered int   `json:"items_ordered"`
	BoxSizes     []int `json:"box_sizes"`
}

func JsonToMap(fileName string) map[string]interface{} {

	jsonFile, err := os.Open(getFilePath(fileName))

	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func getFilePath(fileName string) string {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), fileName)
	return filepath
}
