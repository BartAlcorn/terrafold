package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/turnercode/cp-awfm-common/pkg/log"
	"github.com/turnercode/cp-awfm-common/pkg/version"
)

// SQSMessage is used for passing data.
type Request struct {
	AiringID string `json:"airingID"`
	TitleID  int    `json:"titleID"`
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) (string, error) {
	for _, message := range sqsEvent.Records {
		log.WithPayload("message.Body", message.Body).Debug("SQS Msg received")
		var req Request

		err := json.Unmarshal([]byte(message.Body), &req)
		if err != nil {
			return "", err
		}

		// Fancy Hellow World =====================================================
		then, err := time.Parse("2006-01-02T15:04:05", "2021-12-31T00:00:00")
		if err != nil {
			fmt.Println("ERROR", err)
		}

		remaining := time.Until(then)
		resp := fmt.Sprintf("%v days in 2021!!", int(remaining.Hours()/24))
		fmt.Println("Result:", resp)
		// End  Fancy Hellow World

	} // for range sqsEvent.Records =============================================

	return "", nil
}

// main
func main() {
	version.PrintVersion()
	lambda.Start(handler)
}
