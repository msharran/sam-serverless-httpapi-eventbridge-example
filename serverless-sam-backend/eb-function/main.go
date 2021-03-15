package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type OrderCompleteEvent struct {
	Hello string
}

func Handler(e events.CloudWatchEvent) {
	log.Println(e)
	evnt := &OrderCompleteEvent{}
	if err := json.Unmarshal(e.Detail, evnt); err != nil {
		log.Println(err.Error())
	}
	log.Println(evnt)
}

func main() {
	lambda.Start(Handler)
}
