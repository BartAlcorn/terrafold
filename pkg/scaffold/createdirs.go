package scaffold

import (
	"fmt"
	"os"
)

// CreateDirs creates all of the needed directories.
func CreateDirs(fldr Folder) {

	for _, t := range fldr.Lambda.Triggers {
		err := os.MkdirAll(fldr.LPath+"/handler/"+t, 0755)
		if err != nil {
			fmt.Println("ERROR ", err)
		}
	}

	err := os.MkdirAll(fldr.LPath+"/iac/base", 0755)
	if err != nil {
		fmt.Println("ERROR ", err)
	}

	for _, s := range fldr.Lambda.Stages {
		err = os.MkdirAll(fldr.LPath+"/iac/"+s, 0755)
		if err != nil {
			fmt.Println("ERROR ", err)
		}
	}

}
