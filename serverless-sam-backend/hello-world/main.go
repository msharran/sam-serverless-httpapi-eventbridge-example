package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// resp, err := http.Get(DefaultHTTPGetAddress)

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	e := eventbridge.New(sess)
	d, err := json.Marshal(map[string]string{
		"Hello": "World",
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	_, err = e.PutEvents(&eventbridge.PutEventsInput{
		Entries: []*eventbridge.PutEventsRequestEntry{
			{
				EventBusName: aws.String(os.Getenv("EVENT_BUS_NAME")),
				Source:       aws.String("fireflies.employee"),
				Detail:       aws.String(string(d)),
				DetailType:   aws.String("EmployeeCreated"),
			},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "Event triggered successfully",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
