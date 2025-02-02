package api

import (
	"context"
	"encoding/json"
	"net/http"

	lambdaEvents "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/seannyphoenix/bogie/internal/models"
)

func getEvents() lambdaEvents.LambdaFunctionURLResponse {
	res, err := dynamoDBClient.Scan(context.Background(), &dynamodb.ScanInput{
		TableName: aws.String(bogieTable),
	})
	if err != nil {
		return lambdaEvents.LambdaFunctionURLResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error scanning events",
		}
	}

	var events []models.Event
	err = attributevalue.UnmarshalListOfMaps(res.Items, &events)
	if err != nil {
		return lambdaEvents.LambdaFunctionURLResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error unmarshaling events",
		}
	}

	body, err := json.Marshal(events)
	if err != nil {
		return lambdaEvents.LambdaFunctionURLResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error marshaling events",
		}
	}

	return lambdaEvents.LambdaFunctionURLResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}
