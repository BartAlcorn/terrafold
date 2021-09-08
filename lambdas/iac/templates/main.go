package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// handler
func handler() (string, error) {
	resp := "I am only a scaffold for APP! Some day, I want to grow up to be a REAL lambda!s"
	fmt.Println(resp)

	return resp, nil
}

// main
func main() {
	lambda.Start(handler)
}
