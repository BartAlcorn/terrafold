package scaffold

import (
	"fmt"
	"os"
	"strings"
)

// CreateLambdaDir locates and records the location of the Lambda directory,
// related to the application
func CreateLambdaDir(fldr Folder) Folder {

	ln := strings.ToLower(fldr.Lambda.Name)
	lpath := "lambdas/" + ln

	fldr.LPath = lpath

	if len(fldr.Lambda.Stages) == 0 {
		fldr.Lambda.Stages = append(fldr.Lambda.Stages, "dev")
	}

	if _, err := os.Stat(lpath); os.IsNotExist(err) {
		err := os.MkdirAll(lpath, 0755)
		if err != nil {
			fmt.Println("ERROR ", err)
		}
	}
	return fldr

}
