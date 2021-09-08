package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/turnercode/cp-awfm-common/pkg/version"
)

func handler() (string, error) {
	then, err := time.Parse("2006-01-02T15:04:05", "2021-12-31T00:00:00")
	if err != nil {
		fmt.Println("ERROR", err)
	}

	remaining := time.Until(then)

	resp := fmt.Sprintf("%v days in 2021!!", int(remaining.Hours()/24))

	fmt.Println("Result:", resp)
	return resp, err
}

// main
func main() {
	version.PrintVersion()
	lambda.Start(handler)
}
