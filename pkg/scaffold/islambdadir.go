package scaffold

import (
	"log"
	"os"
	"strings"
)

// IsLambdaDir returns a bool indicatinf if the current working directory is named "lambdas"
func IsLambdaDir() bool {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	parts := strings.Split(path, ":")
	return parts[len(parts)-1] == "lambdas"
}
