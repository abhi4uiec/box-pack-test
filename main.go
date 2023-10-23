package main

import (
	"Test/service"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

type RequestBody struct {
	ItemsOrdered int   `json:"itemsOrdered"`
	BoxSizes     []int `json:"boxSizes"`
}

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody RequestBody
	err := json.Unmarshal([]byte(req.Body), &requestBody)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	numOfItems := requestBody.ItemsOrdered
	if numOfItems < 1 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Num of items to be packed must be greater than or equal to 1",
		}, nil
	}

	boxSizes := requestBody.BoxSizes
	if len(boxSizes) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Mandatory pack sizes not provided",
		}, nil
	}

	resultMap := service.CalculatePackSizes(boxSizes, numOfItems)
	response, err := json.Marshal(resultMap)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil
}
