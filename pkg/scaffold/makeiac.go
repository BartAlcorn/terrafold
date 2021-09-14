package scaffold

import "fmt"

// MakeIAC creates the Terraform Infrastucture As Code(IAC) files from templates
func MakeIAC(fldr Folder) {
	WriteTerraformECR(fldr)
	for _, stg := range fldr.Lambda.Stages {
		fmt.Printf("TerraformIAC for %v...\n", stg)
		ll := fldr.Lambda
		ll.Stage = stg
		Make(fldr.LPath+"/iac/"+stg, "terraformer.tf", "iac-terraformer.tmpl", ll, ll.Overwrite)
		Make(fldr.LPath+"/iac/"+stg, "variables.tf", "iac-variables.tmpl", ll, ll.Overwrite)
		fldr.Lambda.Stage = stg
		Make(fldr.LPath+"/iac/"+stg, "main.tf", "iac-main-header.tmpl", ll, ll.Overwrite)
		WriteTerraformStageMain(fldr)
	}
}
