package scaffold

import (
	"fmt"
	"os"

	"github.com/ttacon/chalk"
)

// MakeIAC creates the Terraform Infrastucture As Code(IAC) files from templates
func MakeIAC(fldr Folder) {
	WriteTerraformECR(fldr)
	WriteTerraformStageMain(fldr)
	Make(fldr.LPath+"/iac/common", "variables.tf", "variables.tf", fldr.Lambda, fldr.Lambda.Overwrite)
	Make(fldr.LPath+"/iac/common", "main.tf", "header.tf", fldr.Lambda, fldr.Lambda.Overwrite)
	for _, stg := range fldr.Lambda.Stages {
		fmt.Printf("TerraformIAC for %v...\n", stg)
		ll := fldr.Lambda
		ll.Stage = stg
		Make(fldr.LPath+"/iac/"+stg, "terraformer.tf", "terraformer.tf", ll, ll.Overwrite)
		Make(fldr.LPath+"/iac/"+stg, "terraform.tfvars", "terraform.tfvars", ll, ll.Overwrite)

		err := os.Symlink("../common/main.tf", fldr.LPath+"/iac/"+stg+"/main.tf")
		if err != nil {
			fmt.Println("ERROR Symlinking ", err)
		}
		fmt.Println(chalk.Green.Color("symlinking  " + fldr.LPath + "/iac/common/main.tf"))

		err = os.Symlink("../common/variables.tf", fldr.LPath+"/iac/"+stg+"/variables.tf")
		if err != nil {
			fmt.Println("ERROR Symlinking ", err)
		}
		fmt.Println(chalk.Green.Color("symlinking  " + fldr.LPath + "/iac/common/variables.tf"))

		for _, t := range fldr.Lambda.Triggers {
			err = os.Symlink("../common/lambda-"+t+".tf", fldr.LPath+"/iac/"+stg+"/lambda-"+t+".tf")
			if err != nil {
				fmt.Println("ERROR Symlinking ", err)
			}
			fmt.Println(chalk.Green.Color("symlinking  " + fldr.LPath + "/iac/common/lambda-" + t + ".tf"))
		}
	}
}
